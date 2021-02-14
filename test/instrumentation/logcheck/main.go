package main

import (
	"github.com/logcheck/check"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(check.Analyzer)
}
