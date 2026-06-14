package gigasecond

// Write a function AddGigasecond that works with time.Time.

import (
	"testing"
	"time"
)

type testCase struct {
	description string
	in          string
	want        string
}

var addCases = []testCase{
	{
		description: "date only specification of time",
		in:          "2011-04-25",
		want:        "2043-01-01T01:46:40",
	},
	{
		description: "second test for date only specification of time",
		in:          "1977-06-13",
		want:        "2009-02-19T01:46:40",
	},
	{
		description: "third test for date only specification of time",
		in:          "1959-07-19",
		want:        "1991-03-27T01:46:40",
	},
	{
		description: "full time specified",
		in:          "2015-01-24T22:00:00",
		want:        "2046-10-02T23:46:40",
	},
	{
		description: "full time with day roll-over",
		in:          "2015-01-24T23:59:59",
		want:        "2046-10-03T01:46:39",
	},
}

// date formats used in test data
const (
	fmtD  = "2006-01-02"
	fmtDT = "2006-01-02T15:04:05"
)

func TestAddGigasecond(t *testing.T) {
	for _, tc := range addCases {
		t.Run(tc.description, func(t *testing.T) {
			in := parse(tc.in, t)
			want := parse(tc.want, t)
			got := AddGigasecond(in)
			if !got.Equal(want) {
				t.Fatalf("AddGigasecond(%v) = %v, want: %v", in, got, want)
			}
		})
	}
}

func parse(s string, t *testing.T) time.Time {
	tc, err := time.Parse(fmtDT, s) // try full date time format first
	if err != nil {
		tc, err = time.Parse(fmtD, s) // also allow just date
	}
	if err != nil {
		t.Fatalf("error in test setup: TestAddGigasecond requires datetime in one of the following formats: \nformat 1:%q\nformat 2:%q\ngot:%q", fmtD, fmtDT, s)
	}
	return tc
}

func BenchmarkAddGigasecond(b *testing.B) {
	for b.Loop() {
		AddGigasecond(time.Time{})
	}
}
