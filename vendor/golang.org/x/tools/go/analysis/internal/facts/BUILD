load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "facts.go",
        "imports.go",
    ],
    importmap = "k8s.io/kubernetes/vendor/golang.org/x/tools/go/analysis/internal/facts",
    importpath = "golang.org/x/tools/go/analysis/internal/facts",
    visibility = ["//vendor/golang.org/x/tools/go/analysis:__subpackages__"],
    deps = [
        "//vendor/golang.org/x/tools/go/analysis:go_default_library",
        "//vendor/golang.org/x/tools/go/types/objectpath:go_default_library",
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
