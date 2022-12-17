#!/bin/sh

#
# Tests are run when on a wip-ci-* branch, see .woodpecker/releases-helper.yml
# It should be changed to run it every time this file is changed when 1.18 is used because 1.17 does not have
# webhooks with the information for that filtering.
#

set -ex

: ${DOCKER_HOST:=unix:///var/run/docker.sock}
: ${ARCHS:=amd64 arm64}
: ${INTEGRATION_USER:=forgejo-integration}
: ${INTEGRATION_IMAGE:=codeberg.org/$INTEGRATION_USER/forgejo}
: ${CI_REPO_OWNER:=dachary}
: ${CI_COMMIT_TAG:=v17.1.42-2}
: ${TAG:=${CI_COMMIT_TAG##v}}
: ${SHORT_TAG=${TAG%.*-*}}
: ${CI_REPO_LINK:=https://codeberg.org/dachary/forgejo}
: ${DOMAIN:=codeberg.org}

: ${VERIFY:=true}
VERIFY_COMMAND='gitea --version'
VERIFY_STRING='built with'

publish() {
    for suffix in '' '-rootless' ; do
	images=""
	for arch in $ARCHS ; do
	    #
	    # Get the image from the integration user
	    #
	    image=$(image_name $INTEGRATION_USER $suffix)
	    docker pull --platform linux/$arch $image
	    #
	    # Verify it is usable
	    #
	    if $VERIFY ; then
		docker run --platform linux/$arch --rm $image $VERIFY_COMMAND | grep "$VERIFY_STRING"
	    fi
	    #
	    # Push the image with a tag reflecting the architecture to the repo owner
	    #
	    arch_image=$(arch_image_name $CI_REPO_OWNER $arch $suffix)
	    docker tag $image $arch_image
	    docker push $arch_image
	    images="$images $arch_image"
	done

	#
	# Push a manifest with all the architectures to the repo owner
	#
	manifest=$(image_name $CI_REPO_OWNER $suffix)
	docker manifest rm $manifest || true
	docker manifest create $manifest $images
	image_put $CI_REPO_OWNER $(image_tag $suffix) $manifest
	image_put $CI_REPO_OWNER $(short_image_tag $suffix) $manifest
	#
	# Sanity check to ensure the manifest that are published can actualy
	# be used.
	#
	for arch in $ARCHS ; do
	    docker pull --platform linux/$arch $(image_name $CI_REPO_OWNER $suffix)
	    docker pull --platform linux/$arch $(short_image_name $CI_REPO_OWNER $suffix)
	done
    done
}

boot() {
    if docker version ; then
	return
    fi
    apk --update --no-cache add coredns jq curl
    ( echo ".:53 {" ; echo "  forward . /etc/resolv.conf"; echo "}" ) > /etc/coredns/Corefile
    coredns -conf /etc/coredns/Corefile &
    /usr/local/bin/dockerd --data-root /var/lib/docker --host=$DOCKER_HOST --dns 172.17.0.3 &
    for i in $(seq 60) ; do
	docker version && break
	sleep 1
    done
    docker version || exit 1
}

authenticate() {
    echo "$RELEASETEAMTOKEN" | docker login --password-stdin --username "$RELEASETEAMUSER" $DOMAIN
    token=$(curl -u$RELEASETEAMUSER:$RELEASETEAMTOKEN -sS https://$DOMAIN/v2/token | jq --raw-output .token)
}

image_delete() {
    curl -sS  -H "Authorization: token $token" -X DELETE https://$DOMAIN/v2/$1/forgejo/manifests/$2
}

image_put() {
    docker manifest inspect $3 > /tmp/manifest.json
    curl -sS  -H "Authorization: token $token" -X PUT --data-binary @/tmp/manifest.json https://$DOMAIN/v2/$1/forgejo/manifests/$2
}

main() {
    boot
    authenticate
    publish
}

image_name() {
    echo $DOMAIN/$1/forgejo:$(image_tag $2)
}

image_tag() {
    echo $TAG$1
}

short_image_name() {
    echo $DOMAIN/$1/forgejo:$(short_image_tag $2)
}

short_image_tag() {
    echo $SHORT_TAG$1
}

arch_image_name() {
    echo $DOMAIN/$1/forgejo:$(arch_image_tag $2 $3)
}

arch_image_tag() {
    echo $TAG-$1$2
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

${@:-main}
