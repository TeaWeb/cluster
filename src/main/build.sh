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