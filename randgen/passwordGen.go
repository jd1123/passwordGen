package randgen

import (
	"crypto/rand"
	"math/big"
)

func GeneratePassword(p Policy) string {
	// Policy - legal characters for password
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	numbers := []rune("0123456789")
	chars := []rune("!@#$%^&*_")

	entropy := []int{p.NumberEntropy, p.LetterEntropy, p.CharacterEntropy}
	nums := []int{p.MinLetters, p.MinNumbers, p.MinCharacters}
	for i := 0; i < 3; i++ {
		if entropy[i] > 0 {
			r, err := rand.Int(rand.Reader, big.NewInt(int64(entropy[i])))
			if err != nil {
				panic(err)
			}
			nums[i] += int(r.Int64())
		}
	}

	total_runes := nums[0] + nums[1] + nums[2]

	b := make([]rune, total_runes)
	for i := 0; i < nums[0]; i++ {
		r, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			panic(err)
		}
		b[i] = letters[r.Int64()]
	}
	for i := nums[0]; i < (nums[0] + nums[1]); i++ {
		r, err := rand.Int(rand.Reader, big.NewInt(int64(len(numbers))))
		if err != nil {
			panic(err)
		}
		b[i] = numbers[r.Int64()]
	}
	for i := nums[0] + nums[1]; i < total_runes; i++ {
		r, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			panic(err)
		}
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
		j, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		if err != nil {
			panic(err)
		}
		k := int(j.Int64())
		m[i] = m[int(k)]
		m[k] = i
	}
	return m
}
