#!/usr/bin/env bash

# shell utilities

function build() {
    VERSION_DATA=`cat ${GOPATH}/src/source/github.com/TeaWeb/cluster-code/consts/consts.go`
    VERSION_DATA=${VERSION_DATA#*"Version = \""}
    VERSION=${VERSION_DATA%%[!0-9.]*}
    TARGET=${GOPATH}/dist/teaweb-cluster-v${VERSION}
    EXT=""
    if [ ${GOOS} = "windows" ]
    then
        EXT=".exe"
    fi

    echo "[================ building ${GOOS}/${GOARCH}/v${VERSION}] ================]"

    echo "[goversion]using" `go version`
    echo "[create target directory]"

    if [ ! -d ${GOPATH}/dist ]
    then
		mkdir ${GOPATH}/dist
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

    echo "[build static file]"

    # build main & plugin
    go build -ldflags="-s -w" -o ${TARGET}/bin/teaweb-cluster${EXT} ${GOPATH}/src/source/github.com/TeaWeb/cluster-code/main/main.go

    echo "[copy files]"
    cp -R ${GOPATH}/src/main/configs/admin.sample.yml ${TARGET}/configs/admin.yml
    cp -R ${GOPATH}/src/main/configs/server.sample.conf ${TARGET}/configs/server.conf
    cp -R ${GOPATH}/src/main/configs/config.sample.yml ${TARGET}/configs/config.yml
    cp -R ${GOPATH}/src/main/scripts ${TARGET}

    cp -R ${GOPATH}/src/main/web/public ${TARGET}/web/
    cp -R ${GOPATH}/src/main/web/views ${TARGET}/web/

    if [ ${GOOS} = "windows" ]
    then
        cp ${GOPATH}/src/main/start.bat ${TARGET}
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