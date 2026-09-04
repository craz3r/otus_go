package hw03frequencyanalysis

import (
	"slices"
	"strings"
	"unicode"
)

func Top10(input string) []string {
	if input == "" {
		return []string{}
	}

	words := strings.Fields(input)
	freqs := make(map[string]int)

	for _, word := range words {
		word = strings.ToLower(word)

		isPunctWord := true

		if letter := strings.ContainsFunc(word, func(r rune) bool {
			return !unicode.IsPunct(r) && !unicode.IsSymbol(r)
		}); letter {
			isPunctWord = false
		}

		if isPunctWord && len(word) > 1 {
			freqs[word]++
			continue
		}

		if word := strings.TrimFunc(word, func(r rune) bool {
			return unicode.IsPunct(r) || unicode.IsSymbol(r)
		}); word != "" {
			freqs[word]++
		}
	}

	uniqueWords := make([]string, 0, len(freqs))

	for word := range freqs {
		uniqueWords = append(uniqueWords, word)
	}

	slices.SortFunc(uniqueWords, func(a, b string) int {
		if r := freqs[b] - freqs[a]; r != 0 {
			return r
		}

		return strings.Compare(a, b)
	})

	return uniqueWords[:min(len(uniqueWords), 10)]
}
