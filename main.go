// package main

// import (
// 	"os"

// 	"github.com/fatih/color"
// 	"github.com/p0bailey/tofuinit/cmd"
// )

// func main() {
// 	if len(os.Args) < 2 {
// 		color.Red(`Please provide the directory name as an argument for the module name.

// tofuinit <module_name>`)
// 		os.Exit(1)
// 	}

// 	dirName := os.Args[1]
// 	cmd.CreateDirectoryAndFiles(dirName)
// }

package main

import (
	"flag"
	"os"

	"github.com/fatih/color"
	"github.com/p0bailey/tofuinit/cmd"
)

func main() {
	help := flag.Bool("help", false, "Display help menu")
	flag.Parse()

	if *help || len(flag.Args()) < 1 {
		displayHelp()
		os.Exit(1)
	}

	dirName := flag.Arg(0)
	cmd.CreateDirectoryAndFiles(dirName)
}

func displayHelp() {
	color.Red(`Usage: tofuinit <module_name>

tofuinit is a tool for initializing a new IaC module with the specified name.

Options:
  --help     Display this help menu and exit

Arguments:
  module_name    The name of the module to be created`)
}
