package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
)

type Policy struct {
	Min_letters       int
	Min_numbers       int
	Min_characters    int
	Letter_entropy    int
	Number_entropy    int
	Character_entropy int
}

func NewPolicy(ml, mn, mc, le, ne, ce int) Policy {
	return Policy{ml, mn, mc, le, ne, ce}
}

func GeneratePassword(p Policy) string {
	// Policy - legal characters for password
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	numbers := []rune("0123456789")
	chars := []rune("!@#$%^&*_")

	entropy := []int{p.Number_entropy, p.Letter_entropy, p.Character_entropy}
	nums := []int{p.Min_letters, p.Min_numbers, p.Min_characters}
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
	letter_entropy := flag.Int("le", 3, "Letter Entropy. Default is 3")
	number_entropy := flag.Int("ne", 6, "Number entropy. Default is 6")
	char_entropy := flag.Int("ce", 3, "Character entropy. Default is 3")
	numPasswords := flag.Int("np", 1, "Number of passwords to generate. Default is 1")
	flag.Parse()

	// Seed the generator for permutations
	p := NewPolicy(*ml, *mn, *mc, *letter_entropy, *number_entropy, *char_entropy)
	// Generate 16 passwords
	if *numPasswords == 1 {
		fmt.Println(GeneratePassword(p))
	} else {
		for i := 0; i < *numPasswords; i++ {
			fmt.Println("Password ", i+1, " : ", GeneratePassword(p))
		}
	}
}
