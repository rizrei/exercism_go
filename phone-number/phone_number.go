package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

var (
	phoneNumber10DigitRegexp = regexp.MustCompile(`^[2-9][0-9]{2}[2-9][0-9]{6}$`)
	phoneNumber11DigitRegexp = regexp.MustCompile(`^1[2-9][0-9]{2}[2-9][0-9]{6}$`)
	invalidPhoneNumberError  = errors.New("invalid phone number")
)

func Number(phoneNumber string) (string, error) {
	var builder strings.Builder

	for _, r := range phoneNumber {
		if unicode.IsDigit(r) {
			builder.WriteRune(r)
		}
	}

	phoneNumber = builder.String()

	switch {
	case phoneNumber10DigitRegexp.MatchString(phoneNumber):
		return phoneNumber, nil
	case phoneNumber11DigitRegexp.MatchString(phoneNumber):
		return phoneNumber[1:], nil
	default:
		return "", invalidPhoneNumberError
	}
}

func AreaCode(phoneNumber string) (string, error) {
	phoneNumber, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}

	return phoneNumber[0:3], nil
}

func Format(phoneNumber string) (string, error) {
	phoneNumber, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("(%s) %s-%s", phoneNumber[0:3], phoneNumber[3:6], phoneNumber[6:]), nil
}
