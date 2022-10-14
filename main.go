package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/alexflint/go-arg"
	"github.com/jd1123/passwordGen/randgen"
)

// main cli function
func main() {
	user, err := user.Current()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	homeDirectory := user.HomeDir
	configfile := homeDirectory + "/.pginfo"
	var defaultPolicy randgen.Policy

	if fileExists(configfile) {
		var a args
		arg.MustParse(&a)
		defaultPolicy = parseConfig(configfile)
		if a.Verbose {
			printPolicy(defaultPolicy)
		}
	} else {
		var a args_noconfig
		arg.MustParse(&a)
		fmt.Println("No config file found. Creating default policy in ~/.pginfo...")
		writeFile(configfile)
		defaultPolicy = parseConfig(configfile)
		if a.Verbose {
			printPolicy(defaultPolicy)
		}
	}
	fmt.Println(randgen.GeneratePassword(defaultPolicy))
}
