This directory contains tool for checking use of unstructured logs in a package. It is created to prevent regression after packages have been migrated to use structured logs.

To run the tool use `go run main.go logcheck.go <package_name>`   
`e.g go run main.go logcheck.go $KUBE_ROOT/pkg/kubelet/lifecycle/`
