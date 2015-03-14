package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
)

type Policy struct {
	MinLetters, MinNumbers, MinCharacters, LetterEntropy, NumberEntropy, CharacterEntropy int
	NoPolicy bool
}

func NewPolicy(ml, mn, mc, le, ne, ce int) Policy {
	return Policy{ml, mn, mc, le, ne, ce, false}
}

func GeneratePassword(p Policy) string {
	// Policy - legal characters for password
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	numbers := []rune("0123456789")
	chars := []rune("!@#$%^&*_")

	entropy := []int{p.NumberEntropy, p.LetterEntropy, p.CharacterEntropy}
	nums := []int{p.MinLetters, p.MinNumbers, p.MinCharacters}
	for i := 0; i < 3; i++ {
		if entropy[i] > 0 {
			r, _ := rand.Int(rand.Reader, big.NewInt(int64(entropy[i])))
			nums[i] += int(r.Int64())
		}
	}

	total_runes := nums[0] + nums[1] + nums[2]

	b := make([]rune, total_runes)
	for i := 0; i < nums[0]; i++ {
		r, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		b[i] = letters[r.Int64()]
	}
	for i := nums[0]; i < (nums[0] + nums[1]); i++ {
		r, _ := rand.Int(rand.Reader, big.NewInt(int64(len(numbers))))
		b[i] = numbers[r.Int64()]
	}
	for i := nums[0] + nums[1]; i < total_runes; i++ {
		r, _ := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		b[i] = chars[r.Int64()]
	}

	// Permute the rune array to generate random
	// permutation
	perm := psPerm(len(b))
	dest := make([]rune, len(b))
	for i, v := range perm {
		dest[v] = b[i]
	}
	return string(dest)
}

func psPerm(n int) []int {
	m := make([]int, n)
	for i := 0; i < n; i++ {
		j, _ := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		k := int(j.Int64())
		m[i] = m[int(k)]
		m[k] = i
	}
	return m
}

func main() {
	ml := flag.Int("ml", 6, "Minimum number of letters. Default is 6.")
	mn := flag.Int("mn", 4, "Minimum number of numbers. Default is 4.")
	mc := flag.Int("mc", 2, "Minimum number of characters. Default is 2.")
	letterEntropy := flag.Int("le", 3, "Letter Entropy. Default is 3")
	numberEntropy := flag.Int("ne", 6, "Number entropy. Default is 6")
	charEntropy := flag.Int("ce", 3, "Character entropy. Default is 3")
	numPasswords := flag.Int("np", 1, "Number of passwords to generate. Default is 1")
	flag.Parse()

	// Seed the generator for permutations
	p := NewPolicy(*ml, *mn, *mc, *letterEntropy, *numberEntropy, *charEntropy)
	// Generate 16 passwords
	if *numPasswords == 1 {
		fmt.Println(GeneratePassword(p))
	} else {
		for i := 0; i < *numPasswords; i++ {
			fmt.Println("Password ", i+1, " : ", GeneratePassword(p))
		}
	}
}
