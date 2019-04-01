# Generate Protobuf models from YANG files

Run the following commands to generate protobuf models from RFC 8345 YANG files:

```console
$ cd ygot-scripts/src
$ go generate
```

The protobuf will be generated under *models/proto/topology* directory. 

**Note**: The protobuf models from from RFC 8345 YANG files have been generated already. 

# Generate Go code from .proto files

1. First install lates version of bazel using the instrcutions that have been posted here: [Installing Bazel] (https://docs.bazel.build/versions/master/install.html)

2. Run the following command to generate Go code from .proto files:

```console
$ cd ygot-scripts/src
$ bazel build //...
```
