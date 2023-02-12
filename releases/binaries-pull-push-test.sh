#!/bin/sh

set -ex

test_teardown() {
    setup_api
    api DELETE repos/$PUSH_USER/forgejo/releases/tags/$TAG || true
    api DELETE repos/$PUSH_USER/forgejo/tags/$TAG || true
    rm -fr dist/release
    setup_tea
    $BIN_DIR/tea login delete $RELEASETEAMUSER || true
}

test_setup() {
    mkdir -p $RELEASE_DIR
    touch $RELEASE_DIR/file-one.txt
    touch $RELEASE_DIR/file-two.txt
}

test_ensure_tag() {
    api DELETE repos/$PUSH_USER/forgejo/tags/$TAG || true
    #
    # idempotent
    #
    ensure_tag
    api GET repos/$PUSH_USER/forgejo/tags/$TAG > /tmp/tag1.json
    ensure_tag
    api GET repos/$PUSH_USER/forgejo/tags/$TAG > /tmp/tag2.json
    diff -u /tmp/tag[12].json
    #
    # sanity check on the SHA of an existing tag
    #
    (
	CI_COMMIT_SHA=12345
	! ensure_tag
    )
    api DELETE repos/$PUSH_USER/forgejo/tags/$TAG
}

#
# Running the test locally instead of within Woodpecker
#
# 1. Setup: obtain a token at https://codeberg.org/user/settings/applications
# 2. Run: RELEASETEAMUSER=<username> RELEASETEAMTOKEn=<apptoken> binaries-pull-push-test.sh test_run
# 3. Verify: (optional) manual verification at https://codeberg.org/<username>/forgejo/releases
# 4. Cleanup: RELEASETEAMUSER=<username> RELEASETEAMTOKEn=<apptoken> binaries-pull-push-test.sh test_teardown
#
test_run() {
    test_teardown
    to_push=/tmp/binaries-releases-to-push
    pulled=/tmp/binaries-releases-pulled
    RELEASE_DIR=$to_push
    test_setup
    test_ensure_tag
    echo "================================ TEST BEGIN"
    push
    RELEASE_DIR=$pulled
    pull
    diff -r $to_push $pulled
    echo "================================ TEST END"
}

: ${CI_REPO_OWNER:=dachary}
: ${PULL_USER=$CI_REPO_OWNER}
: ${PUSH_USER=$CI_REPO_OWNER}
: ${CI_COMMIT_TAG:=W17.8.20-1}
: ${CI_COMMIT_SHA:=$(git rev-parse HEAD)}

. $(dirname $0)/binaries-pull-push.sh
