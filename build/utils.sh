#!/usr/bin/env bash

# shell utilities

function build() {
	ROOT=`pwd`/../
    VERSION_DATA=`cat ${ROOT}/internal/consts/consts.go`
    VERSION_DATA=${VERSION_DATA#*"Version = \""}
    VERSION=${VERSION_DATA%%[!0-9.]*}
    TARGET=${ROOT}/dist/teaweb-cluster-v${VERSION}
    EXT=""
    if [ ${GOOS} = "windows" ]
    then
        EXT=".exe"
    fi

    echo "[================ building ${GOOS}/${GOARCH}/v${VERSION}] ================]"

    echo "[goversion]using" `go version`
    echo "[create target directory]"

    if [ ! -d ${ROOT}/dist ]
    then
		mkdir ${ROOT}/dist
    fi

    if [ -d ${TARGET} ]
    then
        rm -rf ${TARGET}
    fi

    mkdir ${TARGET}
    mkdir ${TARGET}/bin
    mkdir ${TARGET}/web
    mkdir ${TARGET}/web/tmp
    mkdir ${TARGET}/configs
    mkdir ${TARGET}/data
    mkdir ${TARGET}/logs

    echo "[build static file]"

    # build main & plugin
    go build -ldflags="-s -w" -o ${TARGET}/bin/teaweb-cluster${EXT} ${ROOT}/cmd/cluster/main.go

    echo "[copy files]"
    cp -R ${ROOT}/build/configs/admin.sample.yml ${TARGET}/configs/admin.yml
    cp -R ${ROOT}/build/configs/server.sample.conf ${TARGET}/configs/server.conf
    cp -R ${ROOT}/build/configs/config.sample.yml ${TARGET}/configs/config.yml
    cp -R ${ROOT}/build/scripts ${TARGET}

    cp -R ${ROOT}/web/public ${TARGET}/web/
    cp -R ${ROOT}/web/views ${TARGET}/web/

    if [ ${GOOS} = "windows" ]
    then
        cp ${ROOT}/build/start.bat ${TARGET}
    fi

    echo "[zip files]"
    cd ${TARGET}/../
    if [ -f teaweb-cluster-${GOOS}-${GOARCH}-v${VERSION}.zip ]
    then
        rm -f teaweb-cluster-${GOOS}-${GOARCH}-v${VERSION}.zip
    fi
    zip -r -X -q teaweb-cluster-${GOOS}-${GOARCH}-v${VERSION}.zip  teaweb-cluster-v${VERSION}/
    cd -

    echo "[clean files]"
    rm -rf ${TARGET}

    echo "[done]"
}