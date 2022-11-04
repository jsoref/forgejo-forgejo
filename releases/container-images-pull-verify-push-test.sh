#!/bin/sh

set -ex

image_delete() {
    curl -sS -H @$TOKEN_HEADER -X DELETE https://$DOMAIN/v2/$1/forgejo/manifests/$2
}

#
# Create the same set of images that buildx would
#
test_setup() {
    dir=$(dirname $0)

    for suffix in '' '-rootless' ; do
	(
	    cd $dir
	    manifests=""
	    for arch in $ARCHS ; do
		image=$(arch_image_name $PULL_USER $arch $suffix)
		docker build -f Dockerfile$suffix --platform linux/$arch -t $image .
		docker push $image
		images="$images $image"
	    done
	    manifest=$(image_name $PULL_USER $suffix)
	    docker manifest rm $manifest || true
	    docker manifest create $manifest $images
	    image_put $PULL_USER $(image_tag $suffix) $manifest
	)
    done
}

test_teardown() {
    authenticate
    for suffix in '' '-rootless' ; do
	image_delete $PULL_USER $(image_tag $suffix)
	image_delete $PUSH_USER $(image_tag $suffix)
	image_delete $PUSH_USER $(short_image_tag $suffix)
	for arch in $ARCHS ; do
	    image_delete $PULL_USER $(arch_image_tag $arch $suffix)
	    image_delete $PUSH_USER $(arch_image_tag $arch $suffix)
	done
    done
}

#
# Running the test locally instead of within Woodpecker
#
# 1. Setup: obtain a token at https://codeberg.org/user/settings/applications
# 2. Run: RELEASETEAMUSER=<username> RELEASETEAMTOKEn=<apptoken> container-images-pull-verify-push-test.sh test_run
# 3. Verify: (optional) manual verification at https://codeberg.org/<username>/-/packages/container/forgejo/versions
# 4. Cleanup: RELEASETEAMUSER=<username> RELEASETEAMTOKEn=<apptoken> container-images-pull-verify-push-test.sh test_teardown
#
test_run() {
    boot
    test_teardown
    test_setup
    VERIFY_STRING=something
    VERIFY_COMMAND="echo $VERIFY_STRING"
    echo "================================ TEST BEGIN"
    main
    echo "================================ TEST END"
}

: ${CI_REPO_OWNER:=dachary}
: ${PULL_USER:=$CI_REPO_OWNER}
: ${PUSH_USER:=$CI_REPO_OWNER}
: ${CI_COMMIT_TAG:=v17.1.42-2}

. $(dirname $0)/container-images-pull-verify-push.sh
