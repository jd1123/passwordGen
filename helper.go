package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jd1123/passwordGen/randgen"
)

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func writeFile(filename string) {
	const s = `{
        "MinLetters": 10, 
        "MinNumbers": 8, 
        "MinCharacters": 0, 
        "LetterEntropy": 12, 
        "NumberEntropy": 9, 
        "CharacterEntryopy": 0
}`

	err := os.WriteFile(filename, []byte(s), 0644)
	if err != nil {
		panic(err)
	}
}

type ConfigStruct struct {
	MinLetters       int
	MinNumbers       int
	MinCharacters    int
	LetterEntropy    int
	NumberEntropy    int
	CharacterEntropy int
}

func parseConfig(filename string) randgen.Policy {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	decoder := json.NewDecoder(f)
	c := ConfigStruct{}
	err = decoder.Decode(&c)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	policy := randgen.NewPolicy(c.MinLetters, c.MinNumbers, c.MinCharacters, c.LetterEntropy, c.NumberEntropy, c.CharacterEntropy)
	return policy
}

func printPolicy(p randgen.Policy) {
	fmt.Printf("ml: %d mn: %d mc: %d le: %d ne: %d ce: %d\n", p.MinLetters, p.MinNumbers, p.MinCharacters, p.LetterEntropy, p.NumberEntropy, p.CharacterEntropy)
}
