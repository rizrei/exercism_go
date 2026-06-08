package twofer

import "testing"

type testCase struct {
	description string
	input       string
	expected    string
}

var testCases = []testCase{
	{
		description: "no name given",
		input:       "",
		expected:    "One for you, one for me.",
	},
	{
		description: "a name given",
		input:       "Alice",
		expected:    "One for Alice, one for me.",
	},
	{
		description: "another name given",
		input:       "Bob",
		expected:    "One for Bob, one for me.",
	},
}

func TestShareWith(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			got := ShareWith(tc.input)
			if got != tc.expected {
				t.Fatalf("ShareWith(%q)\n got: %q\nwant: %q", tc.input, got, tc.expected)
			}
		})
	}
}
func BenchmarkShareWith(b *testing.B) {
	for b.Loop() {
		for _, test := range testCases {
			ShareWith(test.input)
		}
	}
}
