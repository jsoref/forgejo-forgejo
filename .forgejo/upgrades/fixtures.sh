#!/bin/bash
# SPDX-License-Identifier: MIT

#ONEPIXEL="iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mP8z8BQDwAEhQGAhKmMIQAAAABJRU5ErkJggg=="
#
# one pixel scaled to 290x290 because that's what versions lower or equal to v1.19.4-0 want
# by default and any other size will be transformed which make it difficult to compare.
#
ONEPIXEL="iVBORw0KGgoAAAANSUhEUgAAASIAAAEiCAYAAABdvt+2AAADrElEQVR4nOzUMRHAMADEsL9eeQd6AsOLhMCT/7udAYS+OgDAiICcEQE5IwJyRgTkjAjIGRGQMyIgZ0RAzoiAnBEBOSMCckYE5IwIyBkRkDMiIGdEQM6IgJwRATkjAnJGBOSMCMgZEZAzIiBnREDOiICcEQE5IwJyRgTkjAjIGRGQMyIgZ0RAzoiAnBEBOSMCckYE5IwIyBkRkDMiIGdEQM6IgJwRATkjAnJGBOSMCMgZEZAzIiBnREDOiICcEQE5IwJyRgTkjAjIGRGQMyIgZ0RAzoiAnBEBOSMCckYE5IwIyBkRkDMiIGdEQM6IgJwRATkjAnJGBOSMCMgZEZAzIiBnREDOiICcEQE5IwJyRgTkjAjIGRGQMyIgZ0RAzoiAnBEBOSMCckYE5IwIyBkRkDMiIGdEQM6IgJwRATkjAnJGBOSMCMgZEZAzIiBnREDOiICcEQE5IwJyRgTkjAjIGRGQMyIgZ0RAzoiAnBEBOSMCckYE5IwIyBkRkDMiIGdEQM6IgJwRATkjAnJGBOSMCMgZEZAzIiBnREDOiICcEQE5IwJyRgTkjAjIGRGQMyIgZ0RAzoiAnBEBOSMCckYE5IwIyBkRkDMiIGdEQM6IgJwRATkjAnJGBOSMCMgZEZAzIiBnREDOiICcEQE5IwJyRgTkjAjIGRGQMyIgZ0RAzoiAnBEBOSMCckYE5IwIyBkRkDMiIGdEQM6IgJwRATkjAnJGBOSMCMgZEZAzIiBnREDOiICcEQE5IwJyRgTkjAjIGRGQMyIgZ0RAzoiAnBEBOSMCckYE5IwIyBkRkDMiIGdEQM6IgJwRATkjAnJGBOSMCMgZEZAzIiBnREDOiICcEQE5IwJyRgTkjAjIGRGQMyIgZ0RAzoiAnBEBOSMCckYE5IwIyBkRkDMiIGdEQM6IgJwRATkjAnJGBOSMCMgZEZAzIiBnREDOiICcEQE5IwJyRgTkjAjIGRGQMyIgZ0RAzoiAnBEBOSMCckYE5IwIyBkRkDMiIGdEQM6IgJwRATkjAnJGBOSMCMgZEZAzIiBnREDOiICcEQE5IwJyRgTkjAjIGRGQMyIgZ0RAzoiAnBEBOSMCckYE5IwIyBkRkDMiIGdEQM6IgJwRATkjAnJGBOSMCMgZEZAzIiBnREDOiICcEQE5IwJyRgTkjAjIGRGQMyIgZ0RAzoiAnBEBOSMCckYE5IwIyBkRkDMiIGdEQM6IgJwRATkjAnJGBOSMCMgZEZAzIiBnREDuBQAA//+4jAPFe1H1tgAAAABJRU5ErkJggg=="

function fixture_get_paths_s3() {
    local path=$1

    (
        echo -n $path/
        mc ls --quiet --recursive testS3/$path | sed -e 's/.* //'
    ) > $DIR/path
}

function fixture_get_paths_local() {
    local path=$1
    local work_path=$DIR/forgejo-work-path

    ( cd $work_path ; find $path -type f) > $DIR/path
}

function fixture_get_one_path() {
    local storage=$1
    local path=$2

    fixture_get_paths_$storage $path

    if test $(wc -l < $DIR/path) != 1 ; then
        echo expected one path but got
        cat $DIR/path
        return 1
    fi
    cat $DIR/path
}

function fixture_repo_archive_create() {
    retry curl -f -sS http://${HOST_PORT}/root/fixture/archive/main.zip -o /dev/null
}

function fixture_repo_archive_assert_s3() {
    mc ls --recursive testS3/forgejo/repo-archive | grep --quiet '.zip$'
}

function fixture_repo_archive_assert_local() {
    local path=$1
    local work_path=$DIR/forgejo-work-path

    find $work_path/$path | grep --quiet '.zip$'
}

function fixture_lfs_create() {
    (
        cd $DIR/fixture
        git lfs track "*.txt"
        echo CONTENT > file.txt
        git add .
        git commit -m 'lfs files'
        git push
    )
}

function fixture_lfs_assert_s3() {
    local content=$(mc cat testS3/forgejo/lfs/d6/1e/5fa787e50330288923bd0c9866b44643925965144262288447cf52f9f9b7)
    test "$content" = CONTENT
}

function fixture_lfs_assert_local() {
    local path=$1
    local work_path=$DIR/forgejo-work-path

    local content=$(mc cat $work_path/$path/d6/1e/5fa787e50330288923bd0c9866b44643925965144262288447cf52f9f9b7)
    test "$content" = CONTENT
}

function fixture_packages_create() {
    echo PACKAGE_CONTENT > $DIR/fixture/package
    $work_path/forgejo-api -X DELETE http://${HOST_PORT}/api/packages/${FORGEJO_USER}/generic/test_package/1.0.0/file.txt || true
    $work_path/forgejo-api --upload-file $DIR/fixture/package http://${HOST_PORT}/api/packages/${FORGEJO_USER}/generic/test_package/1.0.0/file.txt
}

function fixture_packages_assert_s3() {
    local content=$(mc cat testS3/forgejo/packages/aa/cf/aacf02e660d813e95d2854e27926ba1ad5c87299dc5f7661d5f08f076c6bbc17)
    test "$content" = PACKAGE_CONTENT
}

function fixture_packages_assert_local() {
    local path=$1

    local content=$(cat $work_path/$path/aa/cf/aacf02e660d813e95d2854e27926ba1ad5c87299dc5f7661d5f08f076c6bbc17)
    test "$content" = PACKAGE_CONTENT
}

function fixture_avatars_create() {
    echo -n $ONEPIXEL | base64 --decode > $DIR/avatar.png
    $work_path/forgejo-client --form avatar=@$DIR/avatar.png http://${HOST_PORT}/user/settings/avatar
}

function fixture_avatars_assert_s3() {
    local filename=$(fixture_get_one_path s3 forgejo/avatars)
    local content=$(mc cat testS3/$filename | base64 -w0)
    test "$content" = "$ONEPIXEL"
}

function fixture_avatars_assert_local() {
    local path=$1

    local filename=$(fixture_get_one_path local $path)
    local content=$(cat $work_path/$filename | base64 -w0)
    test "$content" = "$ONEPIXEL"
}

function fixture_repo_avatars_create() {
    echo -n $ONEPIXEL | base64 --decode > $DIR/repo-avatar.png
    $work_path/forgejo-client --form avatar=@$DIR/repo-avatar.png http://${HOST_PORT}/${FORGEJO_USER}/${FORGEJO_REPO}/settings/avatar
    # v1.21 only
    #$work_path/forgejo-api -X POST --data-raw '{"body":"'$avatar'"}' http://${HOST_PORT}/api/v1/repos/${FORGEJO_USER}/${FORGEJO_REPO}/avatar
}

function fixture_repo_avatars_assert_s3() {
    local filename=$(fixture_get_one_path s3 forgejo/repo-avatars)
    local content=$(mc cat testS3/$filename | base64 -w0)
    test "$content" = "$ONEPIXEL"
}

function fixture_repo_avatars_assert_local() {
    local path=$1

    local filename=$(fixture_get_one_path local $path)
    local content=$(cat $work_path/$filename | base64 -w0)
    test "$content" = "$ONEPIXEL"
}

function fixture_attachments_create_1_18() {
    echo -n $ONEPIXEL | base64 --decode > $DIR/attachment.png
    $work_path/forgejo-client --trace-ascii - --form file=@$DIR/attachment.png http://${HOST_PORT}/${FORGEJO_USER}/${FORGEJO_REPO}/issues/attachments
}

function fixture_attachments_create() {
    if $work_path/forgejo-api http://${HOST_PORT}/api/v1/version | grep --quiet --fixed-strings 1.18. ; then
        fixture_attachments_create_1_18
        return
    fi
    id=$($work_path/forgejo-api --data-raw '{"title":"TITLE"}' http://${HOST_PORT}/api/v1/repos/${FORGEJO_USER}/${FORGEJO_REPO}/issues | jq .id)
    echo -n $ONEPIXEL | base64 --decode > $DIR/attachment.png
    $work_path/forgejo-client -H @$DIR/forgejo-work-path/forgejo-header --form name=attachment.png --form attachment=@$DIR/attachment.png http://${HOST_PORT}/api/v1/repos/${FORGEJO_USER}/${FORGEJO_REPO}/issues/$id/assets
}

function fixture_attachments_assert_s3() {
    local filename=$(fixture_get_one_path s3 forgejo/attachments)
    local content=$(mc cat testS3/$filename | base64 -w0)
    test "$content" = "$ONEPIXEL"
}

function fixture_attachments_assert_local() {
    local path=$1

    local filename=$(fixture_get_one_path local $path)
    local content=$(cat $work_path/$filename | base64 -w0)
    test "$content" = "$ONEPIXEL"
}

function fixture_create() {
    local work_path=$DIR/forgejo-work-path

    rm -fr $DIR/fixture
    mkdir -p $DIR/fixture
    (
        cd $DIR/fixture
        git init
        git checkout -b main
        git remote add origin http://${FORGEJO_USER}:${FORGEJO_PASSWORD}@${HOST_PORT}/${FORGEJO_USER}/${FORGEJO_REPO}
        git config user.email root@example.com
        git config user.name username
        echo SOMETHING > README
        git add README
        git commit -m 'initial commit'
        git push --set-upstream --force origin main
    )
    for fun in ${STORAGE_FUN} ; do
        fixture_${fun}_create
    done
}
