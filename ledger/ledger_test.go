package ledger

import (
	"reflect"
	"strings"
	"testing"
)

type testCase = struct {
	description string
	currency    string
	locale      string
	entries     []Entry
	expected    string
}

var testCases = []testCase{
	{
		description: "empty ledger",
		currency:    "USD",
		locale:      "en-US",
		entries:     []Entry{},
		expected: strings.Join([]string{
			"Date       | Description               | Change       ",
		}, "\n") + "\n",
	},
	{
		description: "one entry",
		currency:    "USD",
		locale:      "en-US",
		entries: []Entry{
			{
				Date:        "2015-01-01",
				Description: "Buy present",
				Change:      -1000,
			},
		},
		expected: strings.Join([]string{
			"Date       | Description               | Change       ",
			"01/01/2015 | Buy present               |      ($10.00)",
		}, "\n") + "\n",
	},
	{
		description: "credit and debit",
		currency:    "USD",
		locale:      "en-US",
		entries: []Entry{
			{
				Date:        "2015-01-02",
				Description: "Get present",
				Change:      1000,
			},
			{
				Date:        "2015-01-01",
				Description: "Buy present",
				Change:      -1000,
			},
		},
		expected: strings.Join([]string{
			"Date       | Description               | Change       ",
			"01/01/2015 | Buy present               |      ($10.00)",
			"01/02/2015 | Get present               |       $10.00 ",
		}, "\n") + "\n",
	},
	{
		description: "final order tie breaker is change",
		currency:    "USD",
		locale:      "en-US",
		entries: []Entry{
			{
				Date:        "2015-01-01",
				Description: "Something",
				Change:      0,
			},
			{
				Date:        "2015-01-01",
				Description: "Something",
				Change:      -1,
			},
			{
				Date:        "2015-01-01",
				Description: "Something",
				Change:      1,
			},
		},
		expected: strings.Join([]string{
			"Date       | Description               | Change       ",
			"01/01/2015 | Something                 |       ($0.01)",
			"01/01/2015 | Something                 |        $0.00 ",
			"01/01/2015 | Something                 |        $0.01 ",
		}, "\n") + "\n",
	},
	{
		description: "overlong description is truncated",
		currency:    "USD",
		locale:      "en-US",
		entries: []Entry{
			{
				Date:        "2015-01-01",
				Description: "Freude schoner Gotterfunken",
				Change:      -123456,
			},
		},
		expected: strings.Join([]string{
			"Date       | Description               | Change       ",
			"01/01/2015 | Freude schoner Gotterf... |   ($1,234.56)",
		}, "\n") + "\n",
	},
	{
		description: "euros",
		currency:    "EUR",
		locale:      "en-US",
		entries: []Entry{
			{
				Date:        "2015-01-01",
				Description: "Buy present",
				Change:      -1000,
			},
		},
		expected: strings.Join([]string{
			"Date       | Description               | Change       ",
			"01/01/2015 | Buy present               |      (€10.00)",
		}, "\n") + "\n",
	},
	{
		description: "Dutch locale",
		currency:    "USD",
		locale:      "nl-NL",
		entries: []Entry{
			{
				Date:        "2015-03-12",
				Description: "Buy present",
				Change:      123456,
			},
		},
		expected: strings.Join([]string{
			"Datum      | Omschrijving              | Verandering  ",
			"12-03-2015 | Buy present               |   $ 1.234,56 ",
		}, "\n") + "\n",
	},
	{
		description: "Dutch locale and euros",
		currency:    "EUR",
		locale:      "nl-NL",
		entries: []Entry{
			{
				Date:        "2015-03-12",
				Description: "Buy present",
				Change:      123456,
			},
		},
		expected: strings.Join([]string{
			"Datum      | Omschrijving              | Verandering  ",
			"12-03-2015 | Buy present               |   € 1.234,56 ",
		}, "\n") + "\n",
	},
	{
		description: "Dutch negative number with 3 digits before decimal point",
		currency:    "USD",
		locale:      "nl-NL",
		entries: []Entry{
			{
				Date:        "2015-03-12",
				Description: "Buy present",
				Change:      -12345,
			},
		},
		expected: strings.Join([]string{
			"Datum      | Omschrijving              | Verandering  ",
			"12-03-2015 | Buy present               |    $ -123,45 ",
		}, "\n") + "\n",
	},
	{
		description: "American negative number with 3 digits before decimal point",
		currency:    "USD",
		locale:      "en-US",
		entries: []Entry{
			{
				Date:        "2015-03-12",
				Description: "Buy present",
				Change:      -12345,
			},
		},
		expected: strings.Join([]string{
			"Date       | Description               | Change       ",
			"03/12/2015 | Buy present               |     ($123.45)",
		}, "\n") + "\n",
	},
	{
		description: "multiple entries on same date ordered by description",
		currency:    "USD",
		locale:      "en-US",
		entries: []Entry{
			{
				Date:        "2015-01-01",
				Description: "Get present",
				Change:      1000,
			},
			{
				Date:        "2015-01-01",
				Description: "Buy present",
				Change:      -1000,
			},
		},
		expected: strings.Join([]string{
			"Date       | Description               | Change       ",
			"01/01/2015 | Buy present               |      ($10.00)",
			"01/01/2015 | Get present               |       $10.00 ",
		}, "\n") + "\n",
	},
}

var failureTestCases = []testCase{
	{
		description: "empty currency",
		currency:    "",
		locale:      "en-US",
		entries:     nil,
	},
	{
		description: "invalid currency",
		currency:    "ABC",
		locale:      "en-US",
		entries:     nil,
	},
	{
		description: "empty locale",
		currency:    "USD",
		locale:      "",
		entries:     nil,
	},
	{
		description: "invalid locale",
		currency:    "USD",
		locale:      "nl-US",
		entries:     nil,
	},
	{
		description: "invalid date (way too high month)",
		currency:    "USD",
		locale:      "en-US",
		entries: []Entry{
			{
				Date:        "2015-131-11",
				Description: "Buy present",
				Change:      12345,
			},
		},
	},
	{
		description: "invalid date (wrong separator)",
		currency:    "USD",
		locale:      "en-US",
		entries: []Entry{
			{
				Date:        "2015-12/11",
				Description: "Buy present",
				Change:      12345,
			},
		},
	},
}

func TestFormatLedgerSuccess(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			actual, err := FormatLedger(tc.currency, tc.locale, tc.entries)
			if err != nil {
				t.Fatalf("FormatLedger for test %q returned unexpected error %q", tc.description, err)
			}
			if actual != tc.expected {
				t.Fatalf("FormatLedger for test %q failed\ngot:\n%#v\nwant:\n%#v", tc.description, actual, tc.expected)
			}
		})
	}
}

func TestFormatLedgerFailure(t *testing.T) {
	for _, tc := range failureTestCases {
		t.Run(tc.description, func(t *testing.T) {
			_, err := FormatLedger(tc.currency, tc.locale, tc.entries)
			if err == nil {
				t.Fatalf("FormatLedger for test %q expected error, got nil", tc.description)
			}
		})
	}
}

func TestFormatLedgerNotChangeInput(t *testing.T) {
	entries := []Entry{
		{
			Date:        "2015-01-02",
			Description: "Freude schöner Götterfunken",
			Change:      1000,
		},
		{
			Date:        "2015-01-01",
			Description: "Buy present",
			Change:      -1000,
		},
	}
	entriesCopy := make([]Entry, len(entries))
	copy(entriesCopy, entries)
	FormatLedger("USD", "en-US", entries)
	if !reflect.DeepEqual(entries, entriesCopy) {
		t.Fatalf("FormatLedger modifies the input entries array")
	}
}

func BenchmarkFormatLedger(b *testing.B) {
	for b.Loop() {
		for _, tc := range testCases {
			FormatLedger(tc.currency, tc.locale, tc.entries)
		}
	}
}
