// The package linters provides an analyzer OsExitCheckAnalyzer
// that checks the call os.exit from the function main of the package main.
package linters

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

// OsExitCheckAnalyzer describes the type required to add an analyzer to a multichecker.
var OsExitCheckAnalyzer = &analysis.Analyzer{
	Name: "osexitcheck",
	Doc:  "looks for the use of os.Exit in a function main of the package main",
	Run:  run,
}

// run analyzes function main in the package main and looks for the os.exit
func run(pass *analysis.Pass) (interface{}, error) {
	selExpr := func(x *ast.SelectorExpr) {
		if id, ok := x.X.(*ast.Ident); ok {
			if (x.Sel.Name == "Exit") && (id.Name == "os") {
				pass.Reportf(x.Pos(), "prohibited expression os.Exit")
			}
		}
	}

	fDecl := func(x *ast.FuncDecl) bool {
		return x.Name.Name == "main"
	}

	for _, file := range pass.Files {
		if (file.Name.Name == "main") && !ast.IsGenerated(file) {
			ast.Inspect(file, func(n ast.Node) bool {
				switch x := n.(type) {
				case *ast.FuncDecl:
					return fDecl(x)
				case *ast.SelectorExpr:
					selExpr(x)
				}
				return true
			})
		}
	}
	return nil, nil
}
