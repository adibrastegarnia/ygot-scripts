#!/bin/bash

go get ./...
go run src/github.com/openconfig/ygot/proto_generator/protogenerator.go  -generate_fakeroot -base_import_path="./models/proto"  -path=./yang -output_dir=./models/proto  -fakeroot_name=network-topology  -package_hierarchy   -package_name=topology -enum_package_name=enums ./yang/ietf-network-topology.yang

