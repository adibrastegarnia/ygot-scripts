# Generate Protobuf models from YANG files

Run the following commands to generate protobuf models from RFC 8345 YANG files:

```console
$ cd ygot-scripts/src
$ go generate
```

# Generate Go code from .proto files

1. First install lates version of bazel using the instrcutions that have been posted here: [https://docs.bazel.build/versions/master/install.html] (Installing Bazel)

2. Run the following command to generate Go code from .proto files:

```console
$ cd ygot-scripts/src
$ bazel build //...
```
