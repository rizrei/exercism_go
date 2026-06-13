package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	return regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\].*`).MatchString(text)
}

func SplitLogLine(text string) []string {
	return regexp.MustCompile(`<[~*=-]*>`).Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	r := regexp.MustCompile(`(?i)".*password.*"`)
	for _, l := range lines {
		if r.MatchString(l) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	return regexp.MustCompile(`end-of-line\d*`).ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	var result []string
	r := regexp.MustCompile(`User\s+(\w+)`)

	for _, line := range lines {
		if matches := r.FindStringSubmatch(line); matches != nil {
			line = fmt.Sprintf("[USR] %s %s", matches[1], line)
		}
		result = append(result, line)
	}

	return result
}
