load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["singlechecker.go"],
    importmap = "k8s.io/kubernetes/vendor/golang.org/x/tools/go/analysis/singlechecker",
    importpath = "golang.org/x/tools/go/analysis/singlechecker",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/golang.org/x/tools/go/analysis:go_default_library",
        "//vendor/golang.org/x/tools/go/analysis/internal/analysisflags:go_default_library",
        "//vendor/golang.org/x/tools/go/analysis/internal/checker:go_default_library",
        "//vendor/golang.org/x/tools/go/analysis/unitchecker:go_default_library",
    ],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
