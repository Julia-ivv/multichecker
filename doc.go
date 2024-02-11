// The package provides a multichecker.
//
// # Multichecker contains analyzers
//
// - all analyzers of the package https://golang.org/x/tools/go/analysis/passes;
//
// - analyzer for checking print-function's names https://github.com/jirfag/go-printf-func-name/pkg/analyzer;
//
// - analyzer to find unused code https://honnef.co/go/tools/unused;
//
// - analyzer to check the use of os.exit in the main function of the main package;
//
// - all analyzers of the package https://pkg.go.dev/honnef.co/go/tools/staticcheck;
//
// - analyzer ST1005 (Incorrectly formatted error string) of the package https://pkg.go.dev/honnef.co/go/tools/stylecheck;
//
// - analyzer ST1019 (Importing the same package multiple times) of the package https://pkg.go.dev/honnef.co/go/tools/stylecheck.
//
// The description of the analyzers can be read follow the relevant links.
//
// # Multichecker configuration file
//
// Staticcheck and stylecheck analyzers are read from a file config.json.
// The configuration file must be in the same directory as the executable file.
// If config.json is not found, only all staticcheck analyzers are used.
//
// Example config.json.
//
//	{
//	    "staticcheck": [
//	        "SA1000",
//	        "SA1001",
//	        "SA9008"
//	    ],
//	    "stylecheck": [
//	        "ST1005",
//	        "ST1019"
//	    ]
//	}
//
// # Run multichecker
//
// Download the analyzer binary file here.....
//
// Run for checking by all analyzers:
//
//	newMultiChecker <project_path>
//
// To select specific analyzers, use the -NAME flag.
//
// Examples:
//
//		newMultiChecker ./...			//checking all files in the current directory and subdirectories
//		newMultiChecker file_name.go		//checking file_name.go
//	 	newMultiChecker ./dir1 ./dir2/...	//checking all files in a subdirectory dir1 and in all subdirectories dir2
//		newMultiChecker -osexitcheck ./...	//run osexitcheck for all files
//
// For more detail run:
//
//	newMultiChecker help
//
// or
//
//	newMultiChecker help name
package main
