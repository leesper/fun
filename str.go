package fun

import "strings"

// Capitalize returns a string with each word capitalized.
func Capitalize(sentence string, separator string) string {
	words := strings.Split(sentence, separator)
	for i := range words {
		words[i] = strings.Title(words[i])
	}
	return strings.Join(words, separator)
}
