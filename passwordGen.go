package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func GeneratePassword(min_letters, min_numbers, min_chars int) string {
	// Policy - legal characters for password
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	numbers := []rune("0123456789")
	chars := []rune("!@#$%^&*")
	// Entropy of number of chars
	letter_ent := 3
	number_ent := 6
	char_ent := 3
	num_letters := rand.Intn(letter_ent) + min_letters
	num_numbers := rand.Intn(number_ent) + min_numbers
	num_chars := rand.Intn(char_ent) + min_chars
	total_runes := num_letters + num_numbers + num_chars
	b := make([]rune, total_runes)
	for i := 0; i < num_letters; i++ {
		b[i] = letters[rand.Intn(len(letters))]
	}
	for i := num_letters; i < (num_letters + num_numbers); i++ {
		b[i] = numbers[rand.Intn(len(numbers))]
	}
	for i := num_letters + num_numbers; i < total_runes; i++ {
		b[i] = chars[rand.Intn(len(chars))]
	}

	// Permute the rune array to generate random
	// permutation
	perm := rand.Perm(len(b))
	dest := make([]rune, len(b))
	for i, v := range perm {
		dest[v] = b[i]
	}
	return string(dest)
}

func main() {
	ml := flag.Int("ml", 6, "Minimum number of letters. Default is 6.")
	mn := flag.Int("mn", 4, "Minimum number of numbers. Default is 4.")
	mc := flag.Int("mc", 2, "Minimum number of characters. Default is 2.")
	numPasswords := flag.Int("np", 16, "Number of passwords to generate. Default is 16")
	flag.Parse()
	// Seed the generator
	rand.Seed(time.Now().Unix())

	// Generate 16 passwords
	for i := 0; i < *numPasswords; i++ {
		fmt.Println("Password ", i, " : ", GeneratePassword(*ml, *mn, *mc))
	}
}
