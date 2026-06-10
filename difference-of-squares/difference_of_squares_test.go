package differenceofsquares

import "testing"

type testCase struct {
	description string
	fn          func(int) int
	fnName      string
	input       int
	expected    int
}

var testCases = []testCase{
	{
		description: "sum of squares 1",
		fn:          SumOfSquares,
		fnName:      "SumOfSquares",
		input:       1,
		expected:    1,
	},
	{
		description: "sum of squares 5",
		fn:          SumOfSquares,
		fnName:      "SumOfSquares",
		input:       5,
		expected:    55,
	},
	{
		description: "sum of squares 100",
		fn:          SumOfSquares,
		fnName:      "SumOfSquares",
		input:       100,
		expected:    338350,
	},
	{
		description: "square of sum 1",
		fn:          SquareOfSum,
		fnName:      "SquareOfSum",
		input:       1,
		expected:    1,
	},
	{
		description: "square of sum 5",
		fn:          SquareOfSum,
		fnName:      "SquareOfSum",
		input:       5,
		expected:    225,
	},
	{
		description: "square of sum 100",
		fn:          SquareOfSum,
		fnName:      "SquareOfSum",
		input:       100,
		expected:    25502500,
	},
	{
		description: "difference of squares 1",
		fn:          Difference,
		fnName:      "Difference",
		input:       1,
		expected:    0,
	},
	{
		description: "difference of squares 5",
		fn:          Difference,
		fnName:      "Difference",
		input:       5,
		expected:    170,
	},
	{
		description: "difference of squares 100",
		fn:          Difference,
		fnName:      "Difference",
		input:       100,
		expected:    25164150,
	},
}

func TestDifferenceOfSquares(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if got := tc.fn(tc.input); got != tc.expected {
				t.Fatalf("%s(%d) = %d, want: %d", tc.fnName, tc.input, got, tc.expected)
			}
		})
	}
}
func BenchmarkSquareOfSum(b *testing.B) {
	for b.Loop() {
		SquareOfSum(100)
	}
}
func BenchmarkSumOfSquares(b *testing.B) {
	for b.Loop() {
		SumOfSquares(100)
	}
}
func BenchmarkDifference(b *testing.B) {
	for b.Loop() {
		Difference(100)
	}
}
