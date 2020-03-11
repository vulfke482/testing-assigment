package src

import (
	"strconv"
	"time"
)

type InputRecord struct {
	OrderID  string
	Date     time.Time
	UserID   string
	Amount   float64
	Currency string
}

// CSV return csv row for given data
func (e InputRecord) CSV() []string {
	return []string{e.OrderID, e.Date.String(), e.UserID, strconv.FormatFloat(e.Amount, 'f', -1, 64), e.Currency}
}


func FromCSV(data []string) (InputRecord, error) {
	OrderID := data[0]

	Date, err:= time.Parse("2006-01-02 15:04:05 -0700 MST", data[1])
	if err != nil {
		return InputRecord{}, err
	}

	UserID := data[2]

	Amount, err := strconv.ParseFloat(data[3], 64)
	if err != nil {
		return InputRecord{}, err
	}

	Currency := data[4]

	record := InputRecord{
		OrderID:OrderID,
		Date:Date,
		UserID:UserID,
		Amount: Amount,
		Currency: Currency,
	}

	return record, nil
}
