load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")



# Go Proto libraries

# topology go proto library
go_proto_library(
    name = "topo_go_proto",
    importpath = "topo",
    proto = ":topo_proto",
    visibility = ["//visibility:public"],
    deps = [
        "ietf_network_go_proto",
        "ywrapper_go_proto",
        "yext_go_proto",
    ],
)

# ietf_network go proto library
go_proto_library(
    name = "ietf_network_go_proto",
    importpath = "topo/ietf_network",
    proto = ":ietf_network_proto",
    visibility = ["//visibility:public"],
    deps = [
        "networks_go_proto",
        "ywrapper_go_proto",
        "yext_go_proto",
    ],
)

# networks go proto library 
go_proto_library(
    name = "networks_go_proto",
    importpath = "topo/ietf_network/networks",
    proto = ":networks_proto",
    visibility = ["//visibility:public"],
    deps = [
        "network_go_proto",
        "ywrapper_go_proto",
        "yext_go_proto",
    ],
)

# network go proto library
go_proto_library(
    name = "network_go_proto",
    importpath = "topo/ietf_network/networks/network",
    proto = ":network_proto",
    visibility = ["//visibility:public"],
    deps = [
        "link_go_proto",
        "node_go_proto",
        "ywrapper_go_proto",
        "yext_go_proto",
    ],
)

# link go proto library 
go_proto_library(
    name = "link_go_proto",
    importpath = "topo/ietf_network/networks/network/link",
    proto = ":link_proto",
    visibility = ["//visibility:public"],
    deps = [
        "ywrapper_go_proto",
        "yext_go_proto",
    ],
)
# node go proto library
go_proto_library(
    name = "node_go_proto",
    importpath = "topo/ietf_network/networks/network/node",
    proto = ":node_proto",
    visibility = ["//visibility:public"],
    deps = [
        "termination_point_go_proto",
        "ywrapper_go_proto",
        "yext_go_proto",
    ],
)

# termination point go proto library
go_proto_library(
    name = "termination_point_go_proto",
    importpath = "topo/ietf_network/networks/network/node/termination_point",
    proto = ":termination_point_proto",
    visibility = ["//visibility:public"],
    deps = [
        "ywrapper_go_proto",
        "yext_go_proto",
    ],
)

# ywrapper go proto library

go_proto_library(
    name = "ywrapper_go_proto",
    importpath = "github.com/openconfig/ygot/proto/ywrapper",
    proto = ":ywrapper_proto",
    visibility = ["//visibility:public"],
)

# yext go proto library 
go_proto_library(
    name = "yext_go_proto",
    importpath = "github.com/openconfig/ygot/proto/yext",
    proto = ":yext_proto",
    visibility = ["//visibility:public"],
)

# proto libraries

# topology proto 
proto_library (
    name = "topo_proto",
    srcs = ["models/proto/topo/topo.proto"],
    deps = [
        "@com_google_protobuf//:any_proto",
        ":ietf_network_proto",
        ":ywrapper_proto",
        ":yext_proto",
    ],
)


# ietf_network proto
proto_library (
    name = "ietf_network_proto",
    srcs = ["models/proto/topo/ietf_network/ietf_network.proto"],
    deps = [
        "@com_google_protobuf//:any_proto",
        ":networks_proto",
        ":ywrapper_proto",
        ":yext_proto",
    ],
)


# networks proto 
proto_library (
    name = "networks_proto",
    srcs = ["models/proto/topo/ietf_network/networks/networks.proto"],
    deps = [
        "@com_google_protobuf//:any_proto",
        ":network_proto",
        ":ywrapper_proto",
        ":yext_proto",
    ],
)

# network proto 
proto_library (
    name = "network_proto",
    srcs = ["models/proto/topo/ietf_network/networks/network/network.proto"],
    deps = [
        "@com_google_protobuf//:any_proto",
        ":link_proto",
        ":node_proto",
        ":ywrapper_proto",
        ":yext_proto",
    ],
)

# link proto
proto_library (
    name = "link_proto",
    srcs = ["models/proto/topo/ietf_network/networks/network/link/link.proto"],
    deps = [
        "@com_google_protobuf//:any_proto",
        ":ywrapper_proto",
        ":yext_proto",
    ],
)

# node proto
proto_library (
    name = "node_proto",
    srcs = ["models/proto/topo/ietf_network/networks/network/node/node.proto"],
    deps = [
        "@com_google_protobuf//:any_proto",
        ":termination_point_proto",
        ":ywrapper_proto",
        ":yext_proto",
    ],
)

# termination_point proto
proto_library (
    name = "termination_point_proto",
    srcs = ["models/proto/topo/ietf_network/networks/network/node/termination_point/termination_point.proto"],
    deps = [
        "@com_google_protobuf//:any_proto",
        ":ywrapper_proto",
        ":yext_proto",
    ],
)



# ywrapper proto
proto_library(
    name = "ywrapper_proto",
    srcs = ["github.com/openconfig/ygot/proto/ywrapper/ywrapper.proto"],
    deps = [
        "@com_google_protobuf//:any_proto",
    ],
)

# yext proto 
proto_library(
    name = "yext_proto",
    srcs = ["github.com/openconfig/ygot/proto/yext/yext.proto"],
    deps = [
        "@com_google_protobuf//:any_proto",
        "@com_google_protobuf//:descriptor_proto"
        
    ],
)
