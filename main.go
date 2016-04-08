package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/jd1123/passwordGen/randgen"
)

// main cli function
func main() {
	app := cli.NewApp()
	app.Name = "passwordGen"
	app.Usage = "Generate Psuedorandom Passwords"
	app.Commands = []cli.Command{
		GenCommand(),
	}

	app.Action = func(c *cli.Context) {
		fmt.Println(randgen.GeneratePassword(randgen.StandardPolicy))
	}

	app.Run(os.Args)
}
