package hamming

import "testing"

type testCase struct {
	description string
	s1          string
	s2          string
	want        int
	expectError bool
}

var testCases = []testCase{
	{
		description: "empty strands",
		s1:          "",
		s2:          "",
		want:        0,
		expectError: false,
	},
	{
		description: "single letter identical strands",
		s1:          "A",
		s2:          "A",
		want:        0,
		expectError: false,
	},
	{
		description: "single letter different strands",
		s1:          "G",
		s2:          "T",
		want:        1,
		expectError: false,
	},
	{
		description: "long identical strands",
		s1:          "GGACTGAAATCTG",
		s2:          "GGACTGAAATCTG",
		want:        0,
		expectError: false,
	},
	{
		description: "long different strands",
		s1:          "GGACGGATTCTG",
		s2:          "AGGACGGATTCT",
		want:        9,
		expectError: false,
	},
	{
		description: "disallow first strand longer",
		s1:          "AATG",
		s2:          "AAA",
		want:        0,
		expectError: true,
	},
	{
		description: "disallow second strand longer",
		s1:          "ATA",
		s2:          "AGTG",
		want:        0,
		expectError: true,
	},
	{
		description: "disallow empty first strand",
		s1:          "",
		s2:          "G",
		want:        0,
		expectError: true,
	},
	{
		description: "disallow empty second strand",
		s1:          "G",
		s2:          "",
		want:        0,
		expectError: true,
	},
}

func TestHamming(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			got, err := Distance(tc.s1, tc.s2)
			switch {
			case tc.expectError:
				if err == nil {
					t.Fatalf("Distance(%q, %q) expected error, got: %d", tc.s1, tc.s2, got)
				}
			case err != nil:
				t.Fatalf("Distance(%q, %q) returned error: %q, want: %d", tc.s1, tc.s2, err, tc.want)
			case got != tc.want:
				t.Fatalf("Distance(%q, %q) = %d, want %d", tc.s1, tc.s2, got, tc.want)
			}
		})
	}
}
func BenchmarkHamming(b *testing.B) {
	for b.Loop() {
		for _, tc := range testCases {
			Distance(tc.s1, tc.s2)
		}
	}
}
