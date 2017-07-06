package builtin

import (
	"fmt"
	"os"
)

// To add builtin commands, modify GshBuiltinCmd,
// GshBuiltinFunc, and the depended function.
var GshBuiltinFunc = []func(){
	gshBuiltinCd,
	gshBuiltinHelp,
	gshBuiltinExit,
}

var GshBuiltinCmd = []string{
	"cd",
	"help",
	"exit",
}

func gshBuiltinCd() {
	fmt.Println("cd")
}

func gshBuiltinHelp() {
	fmt.Println(
		"Mizuki Urushida's GSH\n",
		"Type program and arguments, and hit enter.\n",
		"The following are built in:\n",
	)

	for _, cmd := range GshBuiltinCmd {
		fmt.Printf("\t%s\n", cmd)
	}

	fmt.Println("Use the man com and for information on other programs.")
}

func gshBuiltinExit() {
	os.Exit(0)
}
