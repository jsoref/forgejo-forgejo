#!/bin/sh

#
# Tests are run when on a wip-ci-* branch, see .woodpecker/releases-helper.yml
# It should be changed to run it every time this file is changed when 1.18 is used because 1.17 does not have
# webhooks with the information for that filtering.
#

set -ex

image_delete() {
    curl -sS  -H "Authorization: token $token" -X DELETE https://$DOMAIN/v2/$1/forgejo/manifests/$2
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
		image=$(arch_image_name $INTEGRATION_USER $arch $suffix)
		docker build -f Dockerfile$suffix --platform linux/$arch -t $image .
		docker push $image
		images="$images $image"
	    done
	    manifest=$(image_name $INTEGRATION_USER $suffix)
	    docker manifest rm $manifest || true
	    docker manifest create $manifest $images
	    image_put $INTEGRATION_USER $(image_tag $suffix) $manifest
	)
    done
}

test_teardown() {
    authenticate
    for suffix in '' '-rootless' ; do
	image_delete $INTEGRATION_USER $(image_tag $suffix)
	image_delete $CI_REPO_OWNER $(image_tag $suffix)
	image_delete $CI_REPO_OWNER $(short_image_tag $suffix)
	for arch in $ARCHS ; do
	    image_delete $INTEGRATION_USER $(arch_image_tag $arch $suffix)
	    image_delete $CI_REPO_OWNER $(arch_image_tag $arch $suffix)
	done
    done
}

#
# Running the test locally instead of withing Woodpecker
#
# 1. Setup: obtain a token at https://codeberg.org/user/settings/applications
# 2. Run: RELEASETEAMUSER=<username> RELEASETEAMTOKEn=<apptoken> container-images-pull-verify-push.sh test
# 3. Verify: (optional) manual verification at https://codeberg.org/<username>/-/packages/container/forgejo/versions
# 4. Cleanup: RELEASETEAMUSER=<username> RELEASETEAMTOKEn=<apptoken> container-images-pull-verify-push.sh test_teardown
#
test() {
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
: ${CI_COMMIT_TAG:=v17.1.42-2}

. $(dirname $0)/container-images-pull-verify-push.sh
