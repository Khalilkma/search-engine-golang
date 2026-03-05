package indexer

import (
	"strings"
	"unicode"
)

/*
Steps:
1. Tokenizer -> breaks text into words
2. FilterStopwords -> removes common words like "the", "and", etc
3. Analyze -> runs the whole pipeline
*/

func Tokenizer(text string) []string {

	var tokens []string
	var current strings.Builder

	for _, r := range text {

		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			current.WriteRune(unicode.ToLower(r))
			continue
		}

		if current.Len() > 0 {
			tokens = append(tokens, current.String())
			current.Reset()
		}
	}

	if current.Len() > 0 {
		tokens = append(tokens, current.String())
	}

	return tokens
}

func FilterStopwords(tokens []string) []string {

	result := make([]string, 0, len(tokens))

	for _, t := range tokens {

		if _, exists := stopwords[t]; !exists {
			result = append(result, t)
		}
	}

	return result
}

func Analyze(text string) []string {

	tokens := Tokenizer(text)

	tokens = FilterStopwords(tokens)

	return tokens
}

var stopwords = map[string]struct{}{
	"the":  {},
	"is":   {},
	"and":  {},
	"a":    {},
	"to":   {},
	"of":   {},
	"in":   {},
	"for":  {},
	"on":   {},
	"with": {},
}
