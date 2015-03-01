package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GeneratePassword() string {
	// Policy - legal characters for password
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	numbers := []rune("0123456789")
	chars := []rune("!@#$%^&*")
	// Minimum number of each type of rune
	min_letters := 6
	min_numbers := 4
	min_chars := 2
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
	// Seed the generator
	rand.Seed(time.Now().Unix())

	// Generate 16 passwords
	for i := 0; i < 16; i++ {
		fmt.Println("Password ", i, " : ", GeneratePassword())
	}
}
