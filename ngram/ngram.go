package main

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

// SplitOnNonLetters splits a string on non-letter runes
func SplitOnNonLetters(s string) []string {
	notALetter := func(char rune) bool { return !unicode.IsLetter(char) }
	return strings.FieldsFunc(s, notALetter)
}

var str = "This is a 'sentence' about this thing I wrote. I wrote it yesterday."

func main() {

	str = strings.ToLower(str)
	parts := SplitOnNonLetters(str)
	fmt.Printf("%+v\n", parts)

	fmt.Println(ngrams(parts, 2))
	fmt.Println(ngrams(parts, 3))
}

func ngrams(words []string, size int) (count map[string]uint32) {

	count = make(map[string]uint32, 0)
	offset := int(math.Floor(float64(size / 2)))

	max := len(words)
	for i, _ := range words {
		if i < offset || i+size-offset > max {
			continue
		}
		gram := strings.Join(words[i-offset:i+size-offset], " ")
		count[gram]++
	}

	return count
}
