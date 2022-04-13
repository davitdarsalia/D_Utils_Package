package stringUtils

import (
	"fmt"
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

func ReverseString(s string) string {
	runeArr := []rune(s)

	for i, k := 0, len(runeArr)-1; i < k; i, k = i+1, k-1 {
		runeArr[i], runeArr[k] = runeArr[k], runeArr[i]
	}
	fmt.Println(string(runeArr))
	return string(runeArr)
}
