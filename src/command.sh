#!/bin/bash

go get ./...

go run github.com/openconfig/ygot/proto_generator/protogenerator.go  -generate_fakeroot -base_import_path="./models/proto"  -path=./yang -output_dir=./models/proto  -fakeroot_name=network-topology  -package_hierarchy   -package_name=topology -enum_package_name=enums ./yang/ietf-network-topology.yang



proto_imports=".:${GOPATH}/src/github.com/google/protobuf/src:${GOPATH}/src"
find models/proto -name "*.proto" | while read l; do
  protoc -I=$proto_imports --go_out=. $l
done