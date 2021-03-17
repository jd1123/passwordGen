package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/codegangsta/cli"
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
	defaultPolicy := randgen.StandardPolicy
	if fileExists(configfile) {
		defaultPolicy = parseConfig(configfile)
	} else {
		fmt.Println("No config file found. Creating default policy in ~/.pginfo...")
		writeFile(configfile)
		defaultPolicy = parseConfig(configfile)
	}

	app := cli.NewApp()
	app.Name = "passwordGen"
	app.Usage = "Generate Psuedorandom Passwords"
	app.Commands = []cli.Command{
		GenCommand(),
	}

	app.Action = func(c *cli.Context) {
		fmt.Println(randgen.GeneratePassword(defaultPolicy))
	}

	app.Run(os.Args)
}
