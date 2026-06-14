package ledger

import (
	"errors"
	"fmt"
	"slices"
	"strings"
	"time"
)

// Entry represents a single ledger entry.
type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

type Config struct {
	headers     [3]string
	thousandSep string
	decimalSep  string
	positiveFmt string
	negativeFmt string
	dateLayout  string
}

var (
	supportedLocales = map[string]Config{
		"nl-NL": {
			headers:     [3]string{"Datum", "Omschrijving", "Verandering"},
			thousandSep: ".",
			decimalSep:  ",",
			positiveFmt: "%s %s%s%02d ",
			negativeFmt: "%s -%s%s%02d ",
			dateLayout:  "02-01-2006",
		},
		"en-US": {
			headers:     [3]string{"Date", "Description", "Change"},
			thousandSep: ",",
			decimalSep:  ".",
			positiveFmt: "%s%s%s%02d ",
			negativeFmt: "(%s%s%s%02d)",
			dateLayout:  "01/02/2006",
		},
	}

	currencySymbols = map[string]string{
		"EUR": "€",
		"USD": "$",
	}
)

func FormatLedger(currency, locale string, entries []Entry) (string, error) {
	symbol, err := validateCurrency(currency)
	if err != nil {
		return "", err
	}

	config, err := validateLocale(locale)
	if err != nil {
		return "", err
	}

	entriesCopy := slices.Clone(entries)
	sortEntries(entriesCopy)

	var builder strings.Builder
	writeHeader(&builder, config)

	for _, entry := range entriesCopy {
		line, err := formatEntryLine(entry, symbol, config)
		if err != nil {
			return "", err
		}
		builder.WriteString(line)
	}

	return builder.String(), nil
}

func validateCurrency(currency string) (string, error) {
	symbol, ok := currencySymbols[currency]
	if !ok {
		return "", errors.New("invalid currency")
	}
	return symbol, nil
}

func validateLocale(locale string) (Config, error) {
	config, ok := supportedLocales[locale]
	if !ok {
		return Config{}, errors.New("invalid locale")
	}
	return config, nil
}

func sortEntries(entries []Entry) {
	slices.SortFunc(entries, func(a, b Entry) int {
		if a.Date != b.Date {
			return strings.Compare(a.Date, b.Date)
		}
		if a.Description != b.Description {
			return strings.Compare(a.Description, b.Description)
		}
		return a.Change - b.Change
	})
}

func writeHeader(builder *strings.Builder, config Config) {
	fmt.Fprintf(
		builder,
		"%-10s | %-25s | %-13s\n",
		config.headers[0],
		config.headers[1],
		config.headers[2],
	)
}

func formatEntryLine(entry Entry, symbol string, config Config) (string, error) {
	date, err := formatDate(entry.Date, config.dateLayout)
	if err != nil {
		return "", err
	}

	description := truncateDescription(entry.Description)
	amount := formatAmount(entry.Change, symbol, config)

	return fmt.Sprintf("%-10s | %-25s | %13s\n", date, description, amount), nil
}

func formatDate(dateStr, layout string) (string, error) {
	parsed, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return "", err
	}

	return parsed.Format(layout), nil
}

func truncateDescription(description string) string {
	if len(description) <= 25 {
		return description
	}

	return description[:22] + "..."
}

func formatAmount(change int, symbol string, config Config) string {
	format := config.positiveFmt
	if change < 0 {
		change = -change
		format = config.negativeFmt
	}

	major := change / 100
	minor := change % 100
	majorStr := formatWithThousands(major, config.thousandSep)

	return fmt.Sprintf(format, symbol, majorStr, config.decimalSep, minor)
}

func formatWithThousands(value int, separator string) string {
	if value < 1000 {
		return fmt.Sprintf("%d", value)
	}

	var parts []string
	for value > 0 {
		part := value % 1000
		value /= 1000

		if value > 0 {
			parts = append(parts, fmt.Sprintf("%03d", part))
		} else {
			parts = append(parts, fmt.Sprintf("%d", part))
		}
	}

	slices.Reverse(parts)
	return strings.Join(parts, separator)
}
