/*
Copyright 2021 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

// Doc explaining the tool.
const Doc = "Tool to check use of unstructured logging patterns."

// Analyzer runs static analysis.
var Analyzer = &analysis.Analyzer{
	Name: "logcheck",
	Doc:  Doc,
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {

	for _, file := range pass.Files {

		ast.Inspect(file, func(n ast.Node) bool {

			if fexpr, ok := n.(*ast.CallExpr); ok {
				checkForFunctionExpr(fexpr.Fun, pass)
			}

			return true
		})
	}
	return nil, nil
}

// checkForFunctionExpr checks for unstructured logging function, prints error if found any.
func checkForFunctionExpr(fun ast.Expr, pass *analysis.Pass) {

	if selExpr, ok := fun.(*ast.SelectorExpr); ok {

		fName := selExpr.Sel.Name
		pName, ok := selExpr.X.(*ast.Ident)

		if ok && pName.Name == "klog" && isUnstructured((fName)) {

			msg := fmt.Sprintf("unstructured logging function %q should not be used", fName)
			pass.Report(analysis.Diagnostic{
				Pos:     fun.Pos(),
				Message: msg,
			})
		}
	}
}

func isUnstructured(fName string) bool {

	unstrucured := []string{"Infof", "Info", "Infoln", "InfoDepth", "Info", "Warning", "Warningf", "Warningln", "WarningDepth", "Error", "Errorf", "Errorln", "ErrorDepth", "Fatal", "Fatalf", "Fatalln", "FatalDepth"}

	for _, name := range unstrucured {
		if fName == name {
			return true
		}
	}
	return false
}
