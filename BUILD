load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

go_proto_library(
    name = "link_go_proto",
    importpath = "topology/ietf_network/networks/network/link",
    proto = ":link_proto",
    visibility = ["//visibility:public"],
    deps = [
        "ywrapper_go_proto",
        "yext_go_proto",
    ],
)

go_proto_library(
    name = "ywrapper_go_proto",
    importpath = "src/github.com/openconfig/ygot/proto/ywrapper",
    proto = ":ywrapper_proto",
    visibility = ["//visibility:public"],
)

go_proto_library(
    name = "yext_go_proto",
    importpath = "src/github.com/openconfig/ygot/proto/yext",
    proto = ":yext_proto",
    visibility = ["//visibility:public"],
)



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
