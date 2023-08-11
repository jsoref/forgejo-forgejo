#!/bin/bash
# SPDX-License-Identifier: MIT

set -ex

HOST_PORT=0.0.0.0:3000
STORAGE_PATHS="attachments avatars lfs packages repo-archive repo-avatars"
DIR=/tmp/forgejo-upgrades
SELF_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
PS4='${BASH_SOURCE[0]}:$LINENO: ${FUNCNAME[0]}:  '

function maybe_sudo() {
    if test $(id -u) != 0 ; then
	SUDO=sudo
    fi
}

function dependencies() {
    if ! which curl daemon > /dev/null ; then
	maybe_sudo
	$SUDO apt-get install -y -qq curl daemon
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
    build 1.20.3-0 5.0.2+0-gitea-1.20.3
    build 1.21.0-0 6.0.0+0-gitea-1.21.0
}

function wait_for() {
    rm -f $DIR/setup-forgejo.out
    success=false
    for delay in 1 1 5 5 15 ; do
	if "$@" >> $DIR/setup-forgejo.out 2>&1 ; then
	    success=true
	    break
	fi
	cat $DIR/setup-forgejo.out
	echo waiting $delay
	sleep $delay
    done
    if test $success = false ; then
	cat $DIR/setup-forgejo.out
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

function start() {
    local version=$1

    download $version
    local work_path=$DIR/forgejo-work-path
    daemon --chdir=$DIR --unsafe --env="TERM=$TERM" --env="HOME=$HOME" --env="PATH=$PATH" --pidfile=$DIR/forgejo-pid --errlog=$DIR/forgejo-err.log --output=$DIR/forgejo-out.log -- $DIR/forgejo-$version --config $work_path/app.ini --work-path $work_path
    if ! wait_for grep 'Starting server on' $work_path/log/forgejo.log ; then
	cat $DIR/*.log
	cat $work_path/log/*.log
	return 1
    fi
    create_user $version
    $work_path/forgejo-api http://${HOST_PORT}/api/v1/version
}

function create_user() {
    local version=$1

    local work_path=$DIR/forgejo-work-path

    if test -f $work_path/forgejo-token; then
	return
    fi

    local user=root
    local password=admin1234
    local cli="$DIR/forgejo-$version --config $work_path/app.ini --work-path $work_path"
    $cli admin user create --admin --username "$user" --password "$password" --email "$user@example.com"
    local scopes="--scopes all"
    if echo $version | grep --quiet 1.18. ; then
	scopes=""
    fi
    $cli admin user generate-access-token -u $user --raw $scopes > $work_path/forgejo-token
    ( echo -n 'Authorization: token ' ; cat $work_path/forgejo-token ) > $work_path/forgejo-header
    ( echo "#!/bin/sh" ; echo 'curl -sS -H "Content-Type: application/json" -H @'$work_path/forgejo-header' "$@"' ) > $work_path/forgejo-api && chmod +x $work_path/forgejo-api
}

function stop() {
    if test -f $DIR/forgejo-pid ; then
	local pid=$(cat $DIR/forgejo-pid)
	kill -TERM $pid
	pidwait $pid || true
	for delay in 1 1 2 2 5 5 ; do
	    if ! test -f $DIR/forgejo-pid ; then
		break
	    fi
	    sleep $delay
	done
	! test -f $DIR/forgejo-pid
    fi
    cleanup_logs
}

function reset() {
    local config=$1
    local work_path=$DIR/forgejo-work-path
    rm -fr $work_path
    mkdir -p $work_path
    WORK_PATH=$work_path envsubst < $SELF_DIR/$config-app.ini > $work_path/app.ini
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

    echo "================ See also https://codeberg.org/forgejo/forgejo/pulls/1225"


    echo "================ downgrading from 1.20.3-0 to 1.20.2-0 fails"
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

    echo "================ See also https://codeberg.org/forgejo/forgejo/pulls/1225"

    echo "================ using < 1.20.3-0 and [storage].PATH merge all storage"
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

    echo "================ upgrading from 1.20.2-0 with [storage].PATH fails"
    download 1.20.3-0
    timeout 60 $DIR/forgejo-1.20.3-0 --config $work_path/app.ini --work-path $work_path || true
    if ! grep --fixed-strings --quiet '[storage].PATH is set and may create storage issues' $work_path/log/forgejo.log ; then
	cat $work_path/log/forgejo.log
	return 1
    fi
}

function test_bug_storage_misplace() {
    local work_path=$DIR/forgejo-work-path

    echo "================ See also https://codeberg.org/forgejo/forgejo/pulls/1225"

    echo "================ using < 1.20 and conflicting sections misplace storage"
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

    echo "================ using < 1.20.[12]-0 and conflicting sections ignores [storage.*]"
    for version in 1.20.2-0 ; do
	stop
	reset misplace
	start $version
	for path in ${STORAGE_PATHS} ; do
	    test -d $work_path/elsewhere/$path
	done
    done

    stop

    echo "================ upgrading from 1.20.2-0 with conflicting sections fails"
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
	echo "================ using $config app.ini"
	reset $config

	for version in 1.18.5-0 1.19.4-0 1.20.2-0 1.20.3-0 1.21.0-0 ; do
	    echo "================ run $version"
	    cleanup_storage
	    start $version
	    verify_storage
	    stop
	done
    done
}

function test_upgrades() {
    stop
    dependencies
    build_all
    test_successful_upgrades
    test_bug_storage_misplace
    test_bug_storage_merged
    test_downgrade_1.20.2_fails    
}

"$@"
