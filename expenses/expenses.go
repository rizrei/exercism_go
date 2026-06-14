package expenses

import "fmt"

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(records []Record, predicate func(Record) bool) []Record {
	var result []Record

	for _, record := range records {
		if predicate(record) {
			result = append(result, record)
		}
	}

	return result
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(period DaysPeriod) func(Record) bool {
	return func(record Record) bool {
		return record.Day >= period.From && record.Day <= period.To
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(category string) func(Record) bool {
	return func(record Record) bool {
		return record.Category == category
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(records []Record, period DaysPeriod) float64 {
	return totalByFilter(records, ByDaysPeriod(period))
}

func totalByFilter(records []Record, predicate func(Record) bool) float64 {
	var total float64

	for _, record := range records {
		if predicate(record) {
			total += record.Amount
		}
	}

	return total
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(records []Record, period DaysPeriod, category string) (float64, error) {
	filteredRecords := Filter(records, ByCategory(category))

	if len(filteredRecords) == 0 && len(records) != 0 {
		return 0, fmt.Errorf("unknown category %s", category)
	}

	return TotalByPeriod(filteredRecords, period), nil
}
