package test

import "strings"

func addSpace(word string) string {
	// Split the word
	char := strings.Split(word, "")
	// now join the characters
	spacedWord := strings.Join(char, " ")
	return spacedWord
}
