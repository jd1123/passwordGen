package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jd1123/passwordGen/randgen"
)

type args_noconfig struct {
	Ml      string `help:"Min Letters" default:"15"`
	Mn      string `help:"Min Numbers" default:"12"`
	Mc      string `help:"Min Characters" default:"0"`
	Le      string `help:"Letter Entropy" default:"5"`
	Ne      string `help:"Number Entropy" default:"5"`
	Ce      string `help:"Character Entropy" default:"0"`
	Verbose bool   `arg:"-v, --verbose" help:"verbose"`
}

type args struct {
	Ml      string `help:"Min Letters"`
	Mn      string `help:"Min Numbers"`
	Mc      string `help:"Min Characters"`
	Le      string `help:"Letter Entropy"`
	Ne      string `help:"Number Entropy"`
	Ce      string `help:"Character Entropy"`
	Verbose bool   `arg:"-v, --verbose" help:"verbose"`
}

var Version = "development version"

func (args) Version() string {
	return "v: " + Version
}

func (args) Description() string {
	return "Strong psueodrandom passwords"
}

func createPolicy(a args) randgen.Policy {
	//ml, err := strconv.Atoi(a.Ml)
	ml, err := strconv.Atoi(a.Ml)
	fmt.Println(a.Mn)
	if err != nil {
		fmt.Println("Error: commandline arguments must be integers")
		os.Exit(1)
	}
	mn, err := strconv.Atoi(a.Mn)
	if err != nil {
		fmt.Println("Error: commandline arguments must be integers")
		os.Exit(1)
	}
	mc, err := strconv.Atoi(a.Mc)
	if err != nil {
		fmt.Println("Error: commandline arguments must be integers")
		os.Exit(1)
	}
	le, err := strconv.Atoi(a.Le)
	if err != nil {
		fmt.Println("Error: commandline arguments must be integers")
		os.Exit(1)
	}
	ne, err := strconv.Atoi(a.Ne)
	if err != nil {
		fmt.Println("Error: commandline arguments must be integers")
		os.Exit(1)
	}
	ce, err := strconv.Atoi(a.Ce)
	if err != nil {
		fmt.Println("Error: commandline arguments must be integers")
		os.Exit(1)
	}
	return randgen.NewPolicy(ml, mn, mc, le, ne, ce)
}
