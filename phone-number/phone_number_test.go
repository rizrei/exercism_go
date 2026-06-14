package phonenumber

import (
	"testing"
)

type testCase struct {
	description       string
	input             string
	expectErr         bool
	expectedNumber    string
	expectedAreaCode  string
	expectedFormatted string
}

var testCases = []testCase{
	{
		description:       "cleans the number",
		input:             "(223) 456-7890",
		expectErr:         false,
		expectedNumber:    "2234567890",
		expectedAreaCode:  "223",
		expectedFormatted: "(223) 456-7890",
	},
	{
		description:       "cleans numbers with dots",
		input:             "223.456.7890",
		expectErr:         false,
		expectedNumber:    "2234567890",
		expectedAreaCode:  "223",
		expectedFormatted: "(223) 456-7890",
	},
	{
		description:       "cleans numbers with multiple spaces",
		input:             "223 456   7890   ",
		expectErr:         false,
		expectedNumber:    "2234567890",
		expectedAreaCode:  "223",
		expectedFormatted: "(223) 456-7890",
	},
	{
		description:       "invalid when 9 digits",
		input:             "123456789",
		expectErr:         true,
		expectedNumber:    "",
		expectedAreaCode:  "",
		expectedFormatted: "",
	},
	{
		description:       "invalid when 11 digits does not start with a 1",
		input:             "22234567890",
		expectErr:         true,
		expectedNumber:    "",
		expectedAreaCode:  "",
		expectedFormatted: "",
	},
	{
		description:       "valid when 11 digits and starting with 1",
		input:             "12234567890",
		expectErr:         false,
		expectedNumber:    "2234567890",
		expectedAreaCode:  "223",
		expectedFormatted: "(223) 456-7890",
	},
	{
		description:       "valid when 11 digits and starting with 1 even with punctuation",
		input:             "+1 (223) 456-7890",
		expectErr:         false,
		expectedNumber:    "2234567890",
		expectedAreaCode:  "223",
		expectedFormatted: "(223) 456-7890",
	},
	{
		description:       "invalid when more than 11 digits",
		input:             "321234567890",
		expectErr:         true,
		expectedNumber:    "",
		expectedAreaCode:  "",
		expectedFormatted: "",
	},
	{
		description:       "invalid with letters",
		input:             "523-abc-7890",
		expectErr:         true,
		expectedNumber:    "",
		expectedAreaCode:  "",
		expectedFormatted: "",
	},
	{
		description:       "invalid with punctuations",
		input:             "523-@:!-7890",
		expectErr:         true,
		expectedNumber:    "",
		expectedAreaCode:  "",
		expectedFormatted: "",
	},
	{
		description:       "invalid if area code starts with 0",
		input:             "(023) 456-7890",
		expectErr:         true,
		expectedNumber:    "",
		expectedAreaCode:  "",
		expectedFormatted: "",
	},
	{
		description:       "invalid if area code starts with 1",
		input:             "(123) 456-7890",
		expectErr:         true,
		expectedNumber:    "",
		expectedAreaCode:  "",
		expectedFormatted: "",
	},
	{
		description:       "invalid if exchange code starts with 0",
		input:             "(223) 056-7890",
		expectErr:         true,
		expectedNumber:    "",
		expectedAreaCode:  "",
		expectedFormatted: "",
	},
	{
		description:       "invalid if exchange code starts with 1",
		input:             "(223) 156-7890",
		expectErr:         true,
		expectedNumber:    "",
		expectedAreaCode:  "",
		expectedFormatted: "",
	},
	{
		description:       "invalid if area code starts with 0 on valid 11-digit number",
		input:             "1 (023) 456-7890",
		expectErr:         true,
		expectedNumber:    "",
		expectedAreaCode:  "",
		expectedFormatted: "",
	},
	{
		description:       "invalid if area code starts with 1 on valid 11-digit number",
		input:             "1 (123) 456-7890",
		expectErr:         true,
		expectedNumber:    "",
		expectedAreaCode:  "",
		expectedFormatted: "",
	},
	{
		description:       "invalid if exchange code starts with 0 on valid 11-digit number",
		input:             "1 (223) 056-7890",
		expectErr:         true,
		expectedNumber:    "",
		expectedAreaCode:  "",
		expectedFormatted: "",
	},
	{
		description:       "invalid if exchange code starts with 1 on valid 11-digit number",
		input:             "1 (223) 156-7890",
		expectErr:         true,
		expectedNumber:    "",
		expectedAreaCode:  "",
		expectedFormatted: "",
	},
}

func TestPhoneNumber(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			checkFunction(t, Number, "Number", tc.expectedNumber, tc)
			checkFunction(t, AreaCode, "AreaCode", tc.expectedAreaCode, tc)
			checkFunction(t, Format, "Format", tc.expectedFormatted, tc)
		})
	}
}

func checkFunction(
	t *testing.T,
	f func(string) (string, error),
	funcName string,
	want string,
	tc testCase,
) {
	got, gotErr := f(tc.input)
	switch {
	case tc.expectErr:
		if gotErr == nil {
			t.Errorf("%s(%q) expected error, got: %q", funcName, tc.input, got)
		}
	case gotErr != nil:
		t.Errorf("%s(%q) returned error: %v, want: %q", funcName, tc.input, gotErr, want)
	case got != want:
		t.Errorf("%s(%q) = %q, want: %q", funcName, tc.input, got, want)
	}
}

func BenchmarkNumber(b *testing.B) {
	for b.Loop() {
		for _, test := range testCases {
			Number(test.input)
		}
	}
}

func BenchmarkAreaCode(b *testing.B) {
	for b.Loop() {
		for _, test := range testCases {
			AreaCode(test.input)
		}
	}
}

func BenchmarkFormat(b *testing.B) {
	for b.Loop() {
		for _, test := range testCases {
			Format(test.input)
		}
	}
}
