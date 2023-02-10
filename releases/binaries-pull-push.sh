#!/bin/sh

set -ex

: ${PULL_USER:=forgejo-integration}
if test "$CI_REPO" = "forgejo/release" ; then
    : ${PUSH_USER:=forgejo}
else
    : ${PUSH_USER:=forgejo-experimental}
fi
: ${TAG:=${CI_COMMIT_TAG}}
: ${DOMAIN:=codeberg.org}
: ${RELEASE_DIR:=dist/release}
: ${BIN_DIR:=/tmp}
: ${TEA_VERSION:=0.9.0}


setup_tea() {
    if ! test -f $BIN_DIR/tea ; then
	curl -sL https://dl.gitea.io/tea/$TEA_VERSION/tea-$TEA_VERSION-linux-amd64 > $BIN_DIR/tea
	chmod +x $BIN_DIR/tea
    fi
}

upload() {
    ASSETS=$(ls $RELEASE_DIR/* | sed -e 's/^/-a /')
    echo "${CI_COMMIT_TAG##v}" | grep -qi '\-rc' && export RELEASETYPE="--prerelease" && echo "Uploading as Pre-Release"
    echo "${CI_COMMIT_TAG##v}" | grep -qi '\-test' && export RELEASETYPE="--draft" && echo "Uploading as Draft"
    test ${RELEASETYPE+false} || echo "Uploading as Stable"
    anchor=$(echo $CI_COMMIT_TAG | sed -e 's/^v//' -e 's/[^a-zA-Z0-9]/-/g')
    api POST repos/$PUSH_USER/forgejo/tags --data-raw '{"tag_name": "'$CI_COMMIT_TAG'", "target": "'$CI_COMMIT_SHA'"}'
    $BIN_DIR/tea release create $ASSETS --repo $PUSH_USER/forgejo --note "See https://codeberg.org/forgejo/forgejo/src/branch/forgejo/RELEASE-NOTES.md#${anchor}" --tag $CI_COMMIT_TAG --title $CI_COMMIT_TAG ${RELEASETYPE}
    api GET repos/$PUSH_USER/forgejo/tags/$TAG > /tmp/tag.json
    if test "$(jq --raw-output .commit.sha < /tmp/tag.json)" != "$CI_COMMIT_SHA" ; then
	cat /tmp/tag.json
	echo "the tag SHA in the $PUSH_USER repository does not match the tag SHA that triggered the build: $CI_COMMIT_SHA"
	false
    fi
}

push() {
    if ! which curl ; then
	apk --update --no-cache add curl
    fi
    setup_tea
    GITEA_SERVER_TOKEN=$RELEASETEAMTOKEN $BIN_DIR/tea login add --name $RELEASETEAMUSER --url $DOMAIN
    upload
}

setup_api() {
    if ! which jq || ! which curl ; then
	apk --update --no-cache add jq curl
    fi
}

api() {
    method=$1
    shift
    path=$1
    shift

    curl --fail -X $method -sS -H "Content-Type: application/json" -H "Authorization: token $RELEASETEAMTOKEN" "$@" https://$DOMAIN/api/v1/$path
}

pull() {
    setup_api
    (
	mkdir -p $RELEASE_DIR
	cd $RELEASE_DIR
	api GET repos/$PULL_USER/forgejo/releases/tags/$TAG > /tmp/assets.json
	jq --raw-output '.assets[] | "\(.name) \(.browser_download_url)"' < /tmp/assets.json | while read name url ; do
	    wget --quiet -O $name $url
	done
    )
}


missing() {
    echo need pull or push argument got nothing
    exit 1
}

${@:-missing}
