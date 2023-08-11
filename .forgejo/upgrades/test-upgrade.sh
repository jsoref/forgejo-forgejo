#!/bin/bash
# SPDX-License-Identifier: MIT

#
# Debug loop from the source tree:
#
# ./.forgejo/upgrades/test-upgrade.sh dependencies
# ./.forgejo/upgrades/test-upgrade.sh build_all
# VERBOSE=true ./.forgejo/upgrades/test-upgrade.sh test_downgrade_1.20.2_fails
#
# Everything happens in /tmp/forgejo-upgrades
#

PREFIX===============
HOST_PORT=0.0.0.0:3000
STORAGE_PATHS="attachments avatars lfs packages repo-archive repo-avatars"
STORAGE_FUN="attachments avatars lfs packages repo_archive repo_avatars"
DIR=/tmp/forgejo-upgrades
if ${VERBOSE:-false} ; then
    set -ex
    PS4='${BASH_SOURCE[0]}:$LINENO: ${FUNCNAME[0]}:  '
else
    set -e
fi
SELF_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
: ${FORGEJO_USER:=root}
: ${FORGEJO_REPO:=fixture}
: ${FORGEJO_PASSWORD:=admin1234}

source $SELF_DIR/fixtures.sh

function maybe_sudo() {
    if test $(id -u) != 0 ; then
        SUDO=sudo
    fi
}

function log_info() {
    echo "$PREFIX $@"
}

function dependencies() {
    maybe_sudo
    if ! which curl daemon jq git-lfs > /dev/null ; then
        $SUDO apt-get install -y -qq curl daemon git-lfs jq
    fi

    if ! test -f /usr/local/bin/mc || ! test -f /usr/local/bin/minio  > /dev/null ; then
        $SUDO curl --fail -sS https://dl.min.io/client/mc/release/linux-amd64/mc -o /usr/local/bin/mc
        $SUDO curl --fail -sS https://dl.min.io/server/minio/release/linux-amd64/archive/minio.RELEASE.2023-08-23T10-07-06Z -o /usr/local/bin/minio
    fi
    if ! test -x /usr/local/bin/mc || ! test -x /usr/local/bin/minio  > /dev/null ; then
        $SUDO chmod +x /usr/local/bin/mc
        $SUDO chmod +x /usr/local/bin/minio
    fi

    if ! test -f /usr/local/bin/garage > /dev/null ; then
        $SUDO curl --fail -sS https://garagehq.deuxfleurs.fr/_releases/v0.8.2/x86_64-unknown-linux-musl/garage -o /usr/local/bin/garage
    fi
    if ! test -x /usr/local/bin/garage  > /dev/null ; then
        $SUDO chmod +x /usr/local/bin/garage
    fi
}

function build() {
    local version=$1
    local semver=$2

    if ! test -f $DIR/forgejo-$version ; then
        mkdir -p $DIR
        make VERSION=v$version GITEA_VERSION=v$version FORGEJO_VERSION=$semver TAGS='bindata sqlite sqlite_unlock_notify' generate gitea
        mv gitea $DIR/forgejo-$version
    fi
}

function build_all() {
    test -f Makefile
    build 1.21.0-0 6.0.0+0-gitea-1.21.0
}

function retry() {
    rm -f $DIR/wait-for.out
    success=false
    for delay in 1 1 5 5 15 ; do
        if "$@" >> $DIR/wait-for.out 2>&1 ; then
            success=true
            break
        fi
        cat $DIR/wait-for.out
        echo waiting $delay
        sleep $delay
    done
    if test $success = false ; then
        cat $DIR/wait-for.out
        return 1
    fi
}

function download() {
    local version=$1

    if ! test -f $DIR/forgejo-$version ; then
        mkdir -p $DIR
        wget -O $DIR/forgejo-$version --quiet https://codeberg.org/forgejo/forgejo/releases/download/v$version/forgejo-$version-linux-amd64
        chmod +x $DIR/forgejo-$version
    fi
}

function cleanup_logs() {
    local work_path=$DIR/forgejo-work-path

    rm -f $DIR/*.log
    rm -f $work_path/log/*.log
}

function clobber() {
    rm -fr /tmp/forgejo-upgrades
}

function start_forgejo() {
    local version=$1

    download $version
    local work_path=$DIR/forgejo-work-path
    daemon --chdir=$DIR --unsafe --env="TERM=$TERM" --env="HOME=$HOME" --env="PATH=$PATH" --pidfile=$DIR/forgejo-pid --errlog=$DIR/forgejo-err.log --output=$DIR/forgejo-out.log -- $DIR/forgejo-$version --config $work_path/app.ini --work-path $work_path
    if ! retry grep 'Starting server on' $work_path/log/forgejo.log ; then
        cat $DIR/*.log
        cat $work_path/log/*.log
        return 1
    fi
    create_user $version
}

function start_minio() {
    mkdir -p $DIR/minio
    daemon --chdir=$DIR --unsafe \
           --env="PATH=$PATH" \
           --env=MINIO_ROOT_USER=123456 \
           --env=MINIO_ROOT_PASSWORD=12345678 \
           --env=MINIO_VOLUMES=$DIR/minio \
           --pidfile=$DIR/minio-pid --errlog=$DIR/minio-err.log --output=$DIR/minio-out.log -- /usr/local/bin/minio server
    retry mc alias set testS3 http://127.0.0.1:9000 123456 12345678
}

function start_garage() {
    mkdir -p $DIR/garage/{data,meta}
    cat > $DIR/garage/garage.toml <<EOF
metadata_dir = "$DIR/garage/meta"
data_dir = "$DIR/garage/data"
db_engine = "lmdb"

replication_mode = "none"

rpc_bind_addr = "127.0.0.1:3901"
rpc_public_addr = "127.0.0.1:3901"
rpc_secret = "$(openssl rand -hex 32)"

[s3_api]
s3_region = "us-east-1"
api_bind_addr = "127.0.0.1:9000"
root_domain = ".s3.garage.localhost"

[s3_web]
bind_addr = "127.0.0.1:3902"
root_domain = ".web.garage.localhost"
index = "index.html"

[k2v_api]
api_bind_addr = "127.0.0.1:3904"

[admin]
api_bind_addr = "127.0.0.1:3903"
admin_token = "$(openssl rand -base64 32)"
EOF

    daemon --chdir=$DIR --unsafe \
           --env="PATH=$PATH" \
           --env=RUST_LOG=garage_api=debug \
           --pidfile=$DIR/garage-pid --errlog=$DIR/garage-err.log --output=$DIR/garage-out.log -- /usr/local/bin/garage -c $DIR/garage/garage.toml server

    retry garage -c $DIR/garage/garage.toml status
    garage -c $DIR/garage/garage.toml layout assign -z dc1 -c 1 $(garage -c $DIR/garage/garage.toml status | tail -1 | grep -o '[0-9a-z]*' | head -1)
    ver=$(garage -c $DIR/garage/garage.toml layout show | grep -oP '(?<=Current cluster layout version: )\d+')
    garage -c $DIR/garage/garage.toml layout apply --version $((ver+1))
    garage -c $DIR/garage/garage.toml key info test || garage -c $DIR/garage/garage.toml key import -n test 123456 12345678
    garage -c $DIR/garage/garage.toml key allow --create-bucket test
    retry mc alias set testS3 http://127.0.0.1:9000 123456 12345678
}

function start_s3() {
    local s3_backend=$1

    start_$s3_backend
}

function start() {
    local version=$1
    local s3_backend=${2:-minio}

    start_s3 $s3_backend
    start_forgejo $version
}

function create_user() {
    local version=$1

    local work_path=$DIR/forgejo-work-path

    if test -f $work_path/forgejo-token; then
        return
    fi

    local cli="$DIR/forgejo-$version --config $work_path/app.ini --work-path $work_path"
    $cli admin user create --admin --username "$FORGEJO_USER" --password "$FORGEJO_PASSWORD" --email "$FORGEJO_USER@example.com"
    local scopes="--scopes all"
    if echo $version | grep --quiet 1.18. ; then
        scopes=""
    fi

    #
    # forgejo-cli is to use with api/v1 enpoints
    #
    # tail -1 is because there are logs creating noise in the output in v1.19.4-0
    #
    $cli admin user generate-access-token -u $FORGEJO_USER --raw $scopes | tail -1 > $work_path/forgejo-token
    ( echo -n 'Authorization: token ' ; cat $work_path/forgejo-token ) > $work_path/forgejo-header
    ( echo "#!/bin/sh" ; echo 'curl -f -sS -H "Content-Type: application/json" -H @'$work_path/forgejo-header' "$@"' ) > $work_path/forgejo-api && chmod +x $work_path/forgejo-api
    $work_path/forgejo-api http://${HOST_PORT}/api/v1/version

    #
    # forgejo-client is to use with web endpoints
    #
    #
    # login and obtain a CSRF, all stored in the cookie file
    #
    ( echo "#!/bin/sh" ; echo 'curl --cookie-jar '$DIR/cookies' --cookie '$DIR/cookies' -f -sS "$@"' ) > $work_path/forgejo-client-update-cookies && chmod +x $work_path/forgejo-client-update-cookies
    $work_path/forgejo-client-update-cookies http://${HOST_PORT}/user/login -o /dev/null
    $work_path/forgejo-client-update-cookies --verbose -X POST --data user_name=${FORGEJO_USER} --data password=${FORGEJO_PASSWORD} http://${HOST_PORT}/user/login >& $DIR/login.html
    $work_path/forgejo-client-update-cookies http://${HOST_PORT}/user/login -o /dev/null
    local csrf=$(sed -n -e '/csrf/s/.*csrf\t//p' $DIR/cookies)
    #
    # use the cookie file but do not modify it
    #
    ( echo "#!/bin/sh" ; echo 'curl --cookie '$DIR/cookies' -H "X-Csrf-Token: '$csrf'" -f -sS "$@"' ) > $work_path/forgejo-client && chmod +x $work_path/forgejo-client
}

function stop_daemon() {
    local daemon=$1

    if test -f $DIR/$daemon-pid ; then
        local pid=$(cat $DIR/$daemon-pid)
        kill -TERM $pid
        pidwait $pid || true
        for delay in 1 1 2 2 5 5 ; do
            if ! test -f $DIR/$daemon-pid ; then
                break
            fi
            sleep $delay
        done
        ! test -f $DIR/$daemon-pid
    fi
}

function stop() {
    stop_daemon forgejo
    stop_daemon minio
    stop_daemon garage

    cleanup_logs
}

function reset_forgejo() {
    local config=$1
    local work_path=$DIR/forgejo-work-path
    rm -fr $work_path
    mkdir -p $work_path
    WORK_PATH=$work_path envsubst < $SELF_DIR/$config-app.ini > $work_path/app.ini
}

function reset_minio() {
    rm -fr $DIR/minio
}

function reset_garage() {
    rm -fr $DIR/garage
}

function reset() {
    local config=$1
    reset_forgejo $config
    reset_minio
    reset_garage
}

function verify_storage() {
    local work_path=$DIR/forgejo-work-path

    for path in ${STORAGE_PATHS} ; do
        test -d $work_path/data/$path
    done
}

function cleanup_storage() {
    local work_path=$DIR/forgejo-work-path

    for path in ${STORAGE_PATHS} ; do
        rm -fr $work_path/data/$path
    done
}

function test_downgrade_1.20.2_fails() {
    local work_path=$DIR/forgejo-work-path

    log_info "See also https://codeberg.org/forgejo/forgejo/pulls/1225"

    log_info "downgrading from 1.20.3-0 to 1.20.2-0 fails"
    stop
    reset default
    start 1.20.3-0
    stop
    download 1.20.2-0
    timeout 60 $DIR/forgejo-1.20.2-0 --config $work_path/app.ini --work-path $work_path || true
    if ! grep --fixed-strings --quiet 'use the newer database' $work_path/log/forgejo.log ; then
        cat $work_path/log/forgejo.log
        return 1
    fi
}

function test_bug_storage_merged() {
    local work_path=$DIR/forgejo-work-path

    log_info "See also https://codeberg.org/forgejo/forgejo/pulls/1225"

    log_info "using < 1.20.3-0 and [storage].PATH merge all storage"
    for version in 1.18.5-0 1.19.4-0 1.20.2-0 ; do
        stop
        reset merged
        start $version
        for path in ${STORAGE_PATHS} ; do
            ! test -d $work_path/data/$path
        done
        for path in ${STORAGE_PATHS} ; do
            ! test -d $work_path/merged/$path
        done
        test -d $work_path/merged
    done
    stop

    log_info "upgrading from 1.20.2-0 with [storage].PATH fails"
    download 1.20.3-0
    timeout 60 $DIR/forgejo-1.20.3-0 --config $work_path/app.ini --work-path $work_path || true
    if ! grep --fixed-strings --quiet '[storage].PATH is set and may create storage issues' $work_path/log/forgejo.log ; then
        cat $work_path/log/forgejo.log
        return 1
    fi
}

function test_bug_storage_relative_path() {
    local work_path=$DIR/forgejo-work-path

    log_info "using < 1.20.3-0 legacy [server].XXXX and [picture].XXXX are relative to WORK_PATH"
    for version in 1.18.5-0 1.19.4-0 1.20.2-0 ; do
        stop
        reset legagy-relative
        start $version
        test -d $work_path/relative-lfs
        test -d $work_path/relative-avatars
        test -d $work_path/relative-repo-avatars
    done

    log_info "using >= 1.20.3-0 legacy [server].XXXX and [picture].XXXX are relative to APP_DATA_PATH"
    for version in 1.20.3-0 1.21.0-0 ; do
        stop
        reset legagy-relative
        start $version
        test -d $work_path/data/relative-lfs
        test -d $work_path/data/relative-avatars
        test -d $work_path/data/relative-repo-avatars
    done

    log_info "using >= 1.20.3-0 relative [storage.XXXX].PATHS are relative to APP_DATA_PATH"
    for version in 1.20.3-0 1.21.0-0 ; do
        stop
        reset storage-relative
        start $version
        for path in ${STORAGE_PATHS} ; do
            test -d $work_path/data/relative-$path
        done
    done

    log_info "using 1.20.[12]-0 relative [storage.XXXX].PATHS are inconsistent"
    for version in 1.20.2-0 ; do
        stop
        reset storage-relative
        start $version
        test -d $work_path/data/packages
        test -d $work_path/relative-repo-archive
        test -d $work_path/relative-attachments
        test -d $work_path/relative-lfs
        test -d $work_path/data/avatars
        test -d $work_path/data/repo-avatars
    done

    log_info "using < 1.20 relative [storage.XXXX].PATHS are inconsistent"
    for version in 1.18.5-0 1.19.4-0 ; do
        stop
        reset storage-relative
        start $version
        test -d $work_path/relative-packages
        test -d $work_path/relative-repo-archive
        test -d $work_path/relative-attachments
        test -d $work_path/data/lfs
        test -d $work_path/data/avatars
        test -d $work_path/data/repo-avatars
    done

    log_info "using < 1.20.3-0 relative [XXXX].PATHS are relative to WORK_PATH"
    for version in 1.18.5-0 1.19.4-0 1.20.2-0 ; do
        stop
        reset relative
        start $version
        for path in ${STORAGE_PATHS} ; do
            test -d $work_path/relative-$path
        done
    done

    log_info "using >= 1.20.3-0 relative [XXXX].PATHS are relative to APP_DATA_PATH"
    for version in 1.20.3-0 1.21.0-0 ; do
        stop
        reset relative
        start $version
        for path in ${STORAGE_PATHS} ; do
            test -d $work_path/data/relative-$path
        done
    done

    stop
}

function test_bug_storage_s3_misplace() {
    local work_path=$DIR/forgejo-work-path
    local s3_backend=${2:-minio}

    log_info "See also https://codeberg.org/forgejo/forgejo/issues/1338"

    for version in 1.20.2-0 1.20.3-0 ; do
        log_info "Forgejo $version & $s3_backend"
        stop
        reset misplace-s3
        start $version $s3_backend
        fixture_create
        for fun in ${STORAGE_FUN} ; do
            fixture_${fun}_assert_s3
        done
    done

    for version in 1.18.5-0 1.19.4-0 ; do
        log_info "Forgejo $version & $s3_backend"
        stop
        reset misplace-s3
        start $version $s3_backend
        fixture_create
        #
        # some storage are in S3
        #
        fixture_attachments_assert_s3
        fixture_lfs_assert_s3
        #
        # others are in local
        #
        fixture_repo_archive_assert_local elsewhere/repo-archive
        fixture_avatars_assert_local elsewhere/avatars
        fixture_packages_assert_local elsewhere/packages
        fixture_repo_avatars_assert_local elsewhere/repo-avatars
    done
}

function test_storage_stable_s3() {
    local work_path=$DIR/forgejo-work-path
    local s3_backend=${1:-minio}

    log_info "See also https://codeberg.org/forgejo/forgejo/issues/1338"

    for version in 1.18.5-0 1.19.4-0 1.20.2-0 1.20.3-0 ; do
        log_info "Forgejo $version & $s3_backend"
        stop
        reset stable-s3
        start $version $s3_backend
        fixture_create
        for fun in ${STORAGE_FUN} ; do
            fixture_${fun}_assert_s3
        done
    done
}

function test_bug_storage_misplace() {
    local work_path=$DIR/forgejo-work-path

    log_info "See also https://codeberg.org/forgejo/forgejo/pulls/1225"

    log_info "using < 1.20 and conflicting sections misplace storage"
    for version in 1.18.5-0 1.19.4-0 ; do
        stop
        reset misplace
        start $version
        #
        # some storage are where they should be
        #
        test -d $work_path/data/packages
        test -d $work_path/data/repo-archive
        test -d $work_path/data/attachments
        #
        # others are under APP_DATA_PATH
        #
        test -d $work_path/elsewhere/lfs
        test -d $work_path/elsewhere/avatars
        test -d $work_path/elsewhere/repo-avatars
    done

    log_info "using < 1.20.[12]-0 and conflicting sections ignores [storage.*]"
    for version in 1.20.2-0 ; do
        stop
        reset misplace
        start $version
        for path in ${STORAGE_PATHS} ; do
            test -d $work_path/elsewhere/$path
        done
    done

    stop

    log_info "upgrading from 1.20.2-0 with conflicting sections fails"
    download 1.20.3-0
    timeout 60 $DIR/forgejo-1.20.3-0 --config $work_path/app.ini --work-path $work_path || true
    for path in ${STORAGE_PATHS} ; do
        if ! grep --fixed-strings --quiet "[storage.$path] may conflict" $work_path/log/forgejo.log ; then
            cat $work_path/log/forgejo.log
            return 1
        fi
    done
}

function test_successful_upgrades() {
    for config in default specific ; do
        log_info "using $config app.ini"
        reset $config

        for version in 1.18.5-0 1.19.4-0 1.20.2-0 1.20.3-0 1.21.0-0 ; do
            log_info "run $version"
            cleanup_storage
            start $version
            verify_storage
            stop
        done
    done
}

function run() {
    local fun=$1
    shift

    echo Start running $fun
    mkdir -p $DIR
    > $DIR/$fun.out
    tail --follow $DIR/$fun.out | sed --unbuffered -n -e "/^$PREFIX/s/^$PREFIX //p" &
    pid=$!
    if ! VERBOSE=true ${BASH_SOURCE[0]} $fun "$@" >& $DIR/$fun.out ; then
        kill $pid
        cat $DIR/$fun.out
        echo Failure running $fun
        return 1
    fi
    kill $pid
    echo Success running $fun
}

function test_upgrades() {
    run stop
    run dependencies
    run build_all
    run test_successful_upgrades
    run test_bug_storage_misplace
    run test_bug_storage_merged
    run test_downgrade_1.20.2_fails
    run test_bug_storage_s3_misplace
    run test_storage_stable_s3 minio
    run test_storage_stable_s3 garage
}

"$@"
