package main

import (
	"fmt"
	"github.com/openconfig/ygot/ygot"
)


//go:generate go run src/github.com/openconfig/ygot/proto_generator/protogenerator.go  -generate_fakeroot -base_import_path="../pkg/proto"  -path=../yang -output_dir=../pkg/proto    -package_hierarchy   -package_name=topology -enum_package_name=enums ../yang/ietf-network-topology.yang 

func main() {
	
}