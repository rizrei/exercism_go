package logs

import (
	"strings"
	"unicode/utf8"
)

var applicationChars = map[rune]string{
	'❗': "recommendation",
	'🔍': "search",
	'☀': "weather",
}

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, char := range log {
		switch char {
		case '❗', '🔍', '☀':
			return applicationChars[char]
		default:
			continue
		}
	}

	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var b strings.Builder

	for _, char := range log {
		switch char {
		case oldRune:
			b.WriteRune(newRune)
		default:
			b.WriteRune(char)
		}
	}

	return b.String()
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
