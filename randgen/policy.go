package randgen

type Policy struct {
	MinLetters, MinNumbers, MinCharacters, LetterEntropy, NumberEntropy, CharacterEntropy int
	NoPolicy                                                                              bool
}

func NewPolicy(ml, mn, mc, le, ne, ce int) Policy {
	return Policy{ml, mn, mc, le, ne, ce, false}
}

var StandardPolicy Policy = NewPolicy(15, 12, 0, 5, 5, 0)
