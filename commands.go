package main

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/jd1123/passwordGen/randgen"
)

func GenCommand() cli.Command {
	return cli.Command{
		Name:        "new",
		Usage:       "Generate a new password",
		Description: "Generate a new psuedorandom password",
		Action: func(c *cli.Context) {
			p := randgen.NewPolicy(c.Int("ml"), c.Int("mn"), c.Int("mc"), c.Int("le"), c.Int("ne"), c.Int("ce"))
			fmt.Println(randgen.GeneratePassword(p))
		},
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "ml",
				Value: 7,
				Usage: "Minimum number of letters",
			},
			cli.IntFlag{
				Name:  "mn",
				Value: 4,
				Usage: "Minimum number of numbers",
			},
			cli.IntFlag{
				Name:  "mc",
				Value: 3,
				Usage: "Minimum number of characters",
			},
			cli.IntFlag{
				Name:  "le",
				Value: 0,
				Usage: "Letter entropy (for variable length passwords)",
			},
			cli.IntFlag{
				Name:  "ne",
				Value: 0,
				Usage: "Number entropy (for variable length passwords)",
			},
			cli.IntFlag{
				Name:  "ce",
				Value: 0,
				Usage: "Character entropy (for variable length passwords)",
			},
		},
	}
}
