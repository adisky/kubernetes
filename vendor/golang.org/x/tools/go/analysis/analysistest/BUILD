load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["analysistest.go"],
    importmap = "k8s.io/kubernetes/vendor/golang.org/x/tools/go/analysis/analysistest",
    importpath = "golang.org/x/tools/go/analysis/analysistest",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/golang.org/x/tools/go/analysis:go_default_library",
        "//vendor/golang.org/x/tools/go/analysis/internal/checker:go_default_library",
        "//vendor/golang.org/x/tools/go/packages:go_default_library",
        "//vendor/golang.org/x/tools/internal/lsp/diff:go_default_library",
        "//vendor/golang.org/x/tools/internal/lsp/diff/myers:go_default_library",
        "//vendor/golang.org/x/tools/internal/span:go_default_library",
        "//vendor/golang.org/x/tools/internal/testenv:go_default_library",
        "//vendor/golang.org/x/tools/txtar:go_default_library",
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
