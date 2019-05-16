#!/usr/bin/env bash

. utils.sh

# darwin
export GOPATH=`pwd`/../../
export GOOS=darwin
export GOARCH=amd64

build


# linux 64
export GOPATH=`pwd`/../../
export GOOS=linux
export GOARCH=amd64

build


# linux 32
export GOPATH=`pwd`/../../
export GOOS=linux
export GOARCH=386

build

# linux arm64
export GOPATH=`pwd`/../../
export GOOS=linux
export GOARCH=arm64

build

# linux arm32
export GOPATH=`pwd`/../../
export GOOS=linux
export GOARCH=arm

build

# windows 64
export GOPATH=`pwd`/../../
export GOOS=windows
export GOARCH=amd64

build

# windows 32
export GOPATH=`pwd`/../../
export GOOS=windows
export GOARCH=386

build