package string_similarity

import (
	"strings"
	"unicode"
)

func CompareString(first string, second string) float64 {
	if first == second {
		return 1
	}

	first = AbsValues(first)
	second = AbsValues(second)

	if len(first) == 0 && len(second) == 0 {
		return 1
	}
	if len(first) == 0 || len(second) == 0 {
		return 0
	}
	if len(first) == 1 && len(second) == 1 {
		return 0
	}
	if len(first) < 2 || len(second) < 2 {
		return 0
	}

	bigrams := make(map[string]int, 0)

	for i := 0; i < len(first)-1; i++ {
		bigram := Substring(first, i, 2)

		var count int
		if val, ok := bigrams[bigram]; ok {
			count = val + 1
		} else {
			count = 1
		}

		bigrams[bigram] = count
	}

	intersectionSize := 0

	for i := 0; i < len(second)-1; i++ {
		bigram := Substring(second, i, 2)

		var count int
		if val, ok := bigrams[bigram]; ok {
			count = val
		} else {
			count = 0
		}

		if count > 0 {
			bigrams[bigram] = count - 1
			intersectionSize++
		}
	}

	return float64(2.0*intersectionSize) / float64(len(first)+len(second)-2)
}

func Substring(input string, start int, end int) string {
	if start < 0 {
		start = 0
	}

	if end < 0 {
		end = 0
	}

	if start > end {
		temp := start
		start = end
		end = temp
	}

	if end > len(input) {
		end = len(input)
	}

	return input[start:end]
}

func AbsValues(input string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, input)
}
