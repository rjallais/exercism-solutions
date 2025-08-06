package expenses

import "errors"

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
func Filter(in []Record, predicate func(Record) bool) []Record {
	var result = make([]Record, 0)
	for _, rec := range in {
		if predicate(rec) {
			result = append(result, rec)
		}
	}
	return result
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(record Record) bool {
		if p.From <= record.Day && record.Day <= p.To {
			return true
		} else {
			return false
		}
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(record Record) bool {
		if record.Category == c {
			return true
		} else {
			return false
		}
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	var total = 0.0
	for _, record := range in {
		if p.From <= record.Day && record.Day <= p.To {
			total += record.Amount
		}
	}
	return total
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	var (
		total      = 0.0
		categories = make(map[string]bool)
	)
	for _, record := range in {
		categories[record.Category] = true
		if p.From <= record.Day && record.Day <= p.To && c == record.Category {
			total += record.Amount
		}
	}
	if !categories[c] {
		return 0, errors.New("unknown category entertainment")
	}
	return total, nil
}
