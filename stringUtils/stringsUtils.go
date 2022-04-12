package stringUtils

import (
	"unicode"
)

func MakeUpperCase(s string) string {
	var runeArr []rune

	for _, letter := range s {
		runeArr = append(runeArr, unicode.ToUpper(letter))
	}
	// Conversion From Slice Of Runes To String
	res := string(runeArr)
	return res
}

func MakeLowerCase(s string) string {
	var runeArr []rune

	for _, letter := range s {
		runeArr = append(runeArr, unicode.ToLower(letter))
	}
	// Conversion From Slice Of Runes To String
	res := string(runeArr)
	return res
}
