package mystrings

import "strings"

func Reverse(s string) string {
	result := ""
	for _, v := range s {
		result = string(v) + result
	}
	return result
}

func addSpace(word string) string {
	// Split the word
	char := strings.Split(word, "")
	// now join the characters
	spacedWord := strings.Join(char, " ")
	return spacedWord
}
