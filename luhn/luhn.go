package luhn

import "strings"

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")

	if len(id) <= 1 {
		return false
	}

	sum := 0
	double := false

	for i := len(id) - 1; i >= 0; i-- {
		c := id[i]

		switch {
		case c < '0' || c > '9':
			return false

		default:
			d := int(c - '0')

			if double {
				d *= 2
				if d > 9 {
					d -= 9
				}
			}

			sum += d
			double = !double
		}
	}

	return sum%10 == 0
}