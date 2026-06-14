package anagram

import (
	"slices"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	var result []string
	subject = strings.ToLower(subject)
	sortedSubject := sortString(subject)

	for _, candidate := range candidates {
		if isAnagram(subject, sortedSubject, candidate) {
			result = append(result, candidate)
		}
	}

	return result
}

func isAnagram(subject, sortedSubject, candidate string) bool {
	candidate = strings.ToLower(candidate)
	return candidate != subject && sortedSubject == sortString(candidate)
}

func sortString(s string) string {
	runes := []rune(s)

	slices.Sort(runes)

	return string(runes)
}
