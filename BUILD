load("@build_stack_rules_proto//go:go_proto_library.bzl", "go_proto_library")


proto_library (
    name = "link_proto",
    srcs = ["src/models/proto/topology/ietf_network/networks/network/link/link.proto"],
    deps = [
        "@com_google_protobuf//:any_proto",
        ":ywrapper_proto",
        ":yext_proto",
    ],
)


proto_library(
    name = "ywrapper_proto",
    srcs = ["src/github.com/openconfig/ygot/proto/ywrapper/ywrapper.proto"],
    deps = [
        "@com_google_protobuf//:any_proto",
    ],
)


proto_library(
    name = "yext_proto",
    srcs = ["src/github.com/openconfig/ygot/proto/yext/yext.proto"],
    deps = [
        "@com_google_protobuf//:any_proto",
        "@com_google_protobuf//:descriptor_proto"
        
    ],
)
