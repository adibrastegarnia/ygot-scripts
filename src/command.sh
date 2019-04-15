#!/bin/bash

go get ./...

go run third_party/ygot/generator/generator.go -path=yang -output_file=pkg/topo/topo.go -package_name=topo -generate_fakeroot -generate_getters -generate_append -generate_delete -generate_leaf_getters -generate_leaf_setters -include_model_data -fakeroot_name=topology yang/ietf-network.yang yang/ietf-network-topology.yang
go run third_party/ygot/proto_generator/protogenerator.go  -generate_fakeroot  -base_import_path="./models/proto"  -path=./yang  -output_dir=./models/proto  -fakeroot_name=topology  -package_hierarchy    -package_name=topo  -enum_package_name=enums ./yang/ietf-network-topology.yang  
