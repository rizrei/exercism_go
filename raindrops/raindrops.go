package raindrops

import (
	"strconv"
	"strings"
)

func Convert(number int) string {
	var raindrops = []struct {
		factor int
		sound  string
	}{
		{3, "Pling"},
		{5, "Plang"},
		{7, "Plong"},
	}

	var result strings.Builder

	for _, raindrop := range raindrops {
		if number%raindrop.factor == 0 {
			result.WriteString(raindrop.sound)
		}
	}

	if result.String() == "" {
		return strconv.Itoa(number)
	}

	return result.String()
}
