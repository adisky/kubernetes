load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "analysis.go",
        "diagnostic.go",
        "doc.go",
        "validate.go",
    ],
    importmap = "k8s.io/kubernetes/vendor/golang.org/x/tools/go/analysis",
    importpath = "golang.org/x/tools/go/analysis",
    visibility = ["//visibility:public"],
    deps = ["//vendor/golang.org/x/tools/internal/analysisinternal:go_default_library"],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [
        ":package-srcs",
        "//vendor/golang.org/x/tools/go/analysis/analysistest:all-srcs",
        "//vendor/golang.org/x/tools/go/analysis/internal/analysisflags:all-srcs",
        "//vendor/golang.org/x/tools/go/analysis/internal/checker:all-srcs",
        "//vendor/golang.org/x/tools/go/analysis/internal/facts:all-srcs",
        "//vendor/golang.org/x/tools/go/analysis/singlechecker:all-srcs",
        "//vendor/golang.org/x/tools/go/analysis/unitchecker:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
