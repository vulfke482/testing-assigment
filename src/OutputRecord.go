package src

import "strconv"

// OutputRecord is a struct for output data.
type OutputRecord struct {
	UserID   string
	Amount   float64
	Currency string
	Fee      float64
}

// newOutputRecord creates new OutputRecord
func newOutputRecord(UserID string, Currency string, Amount float64, Fee float64) OutputRecord {
	return OutputRecord{
		UserID,
		Amount,
		Currency,
		Fee,
	}
}

// CSV return csv row for output data
func (e OutputRecord) CSV() []string {
	return []string{e.UserID, strconv.FormatFloat(e.Amount, 'f', -1, 64), e.Currency, strconv.FormatFloat(e.Fee, 'f', -1, 64)}
}
