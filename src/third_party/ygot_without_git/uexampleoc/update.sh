#!/bin/bash

# Hack to ensure that if we are running on OS X with a homebrew installed
# GNU sed then we can still run sed.
runsed() {
  if hash gsed 2>/dev/null; then
    gsed "$@"
  else
    sed "$@"
  fi
}

git clone https://github.com/openconfig/public.git
mkdir deps
cp ../demo/getting_started/yang/{ietf,iana}* deps
go run ../generator/generator.go -path=public,deps -output_file=oc.go \
  -package_name=uexampleoc -generate_fakeroot -fakeroot_name=device \
  -exclude_modules=ietf-interfaces \
  -generate_rename \
  -generate_append \
  -generate_getters \
  -generate_leaf_getters \
  -annotations \
  public/release/models/system/openconfig-system.yang \
  public/release/models/platform/openconfig-platform.yang \
  public/release/models/optical-transport/openconfig-terminal-device.yang
runsed -i 's/This package was generated by.*/NOTE WELL: This is an example code file that is distributed with ygot.\nIt should not be used within your application, as it WILL change,\nwithout warning. Rather, you should generate structs directly from\nOpenConfig models using the ygot package.\n\nThis package was generated by github.com\/openconfig\/ygot/g' oc.go
gofmt -w -s oc.go
rm -rf deps public
