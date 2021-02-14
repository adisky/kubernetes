package check

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

// Doc explaining the tool.
const Doc = "Tool to check regression for migrated packages to structured logs."

// Analyzer runs static analysis.
var Analyzer = &analysis.Analyzer{
	Name: "logcheck",
	Doc:  Doc,
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {

	for _, file := range pass.Files {

		ast.Inspect(file, func(n ast.Node) bool {

			fexpr, ok := n.(*ast.CallExpr)

			if ok {
				// it contains function expression
				fun := fexpr.Fun
				selExpr, ok := fun.(*ast.SelectorExpr)

				if ok {
					// sel.X is package name
					expr := selExpr.X
					// se.Sel is function name
					fName := selExpr.Sel.Name
					pName, ok := expr.(*ast.Ident)
					if ok {
						if pName.Name == "klog" && (fName == "Errorf" || fName == "Infof") {
							msg := fmt.Sprintf("unstructured logging function %q should not be used", fName)
							pass.Report(analysis.Diagnostic{
								Pos:     fun.Pos(),
								Message: msg,
							})
						}
					}

				}

			}

			return true
		})

	}

	return nil, nil
}
