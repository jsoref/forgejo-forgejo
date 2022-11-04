#!/bin/sh

set -ex

: ${DOCKER_HOST:=unix:///var/run/docker.sock}
: ${ARCHS:=amd64 arm64}
: ${PULL_USER:=forgejo-integration}
if test "$CI_REPO" = "forgejo/release" ; then
    : ${PUSH_USER:=forgejo}
else
    : ${PUSH_USER:=forgejo-experimental}
fi
: ${INTEGRATION_IMAGE:=codeberg.org/$PULL_USER/forgejo}
: ${TAG:=${CI_COMMIT_TAG##v}}
: ${SHORT_TAG=${TAG%.*-*}}
: ${DOMAIN:=codeberg.org}
: ${TOKEN_HEADER:=/tmp/token$$}
trap "rm -f ${TOKEN_HEADER}" EXIT

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
	    image=$(image_name $PULL_USER $suffix)
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
	    arch_image=$(arch_image_name $PUSH_USER $arch $suffix)
	    docker tag $image $arch_image
	    docker push $arch_image
	    images="$images $arch_image"
	done

	#
	# Push a manifest with all the architectures to the repo owner
	#
	manifest=$(image_name $PUSH_USER $suffix)
	docker manifest rm $manifest || true
	docker manifest create $manifest $images
	image_put $PUSH_USER $(image_tag $suffix) $manifest
	image_put $PUSH_USER $(short_image_tag $suffix) $manifest
	#
	# Sanity check to ensure the manifest that are published can actualy
	# be used.
	#
	for arch in $ARCHS ; do
	    docker pull --platform linux/$arch $(image_name $PUSH_USER $suffix)
	    docker pull --platform linux/$arch $(short_image_name $PUSH_USER $suffix)
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
    curl -u$RELEASETEAMUSER:$RELEASETEAMTOKEN -sS https://$DOMAIN/v2/token | jq --raw-output '"Authorization: token \(.token)"' > $TOKEN_HEADER
}

image_put() {
    docker manifest inspect $3 > /tmp/manifest.json
    curl -sS -H @$TOKEN_HEADER -X PUT --data-binary @/tmp/manifest.json https://$DOMAIN/v2/$1/forgejo/manifests/$2
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

${@:-main}
