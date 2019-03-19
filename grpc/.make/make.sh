#!/bin/sh

set -e

build() {
    time go build -o main
    ls -lah main
}

deploy() {
	true
}

info() {
	true
}

package() {
    true
}


gen() {
    metactl gen \
        --gen-config .make/gen.yaml \
        --version v2019_02

    cd schema

#    protoc -I . schema/*/* --go_out=plugins=grpc:.
#    protoc -I . schema/indexpage/indexpage.proto --go_out=plugins=grpc:.

#    protoc --proto_path=schema --go_out=plugins=grpc:. schema/brand/brand_.proto
#
#    exit

    for file in ./*/*.proto
    do
        name=${file#*/}
        protoc -I . ${name} --go_out=plugins=grpc:. &
    done
}

provision() {
    true
}