package generator

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	"git.sansera.com/mtkach/golang-test-assignment/utils/rand"
	"github.com/pkg/errors"
)

const (
	buffer = 1000
	freq   = 100
)

var (
	ordersProbabilities = map[interface{}]uint64{
		"1": 30,
		"2": 20,
		"3": 5,
		"4": 5,
		"5": 10,
		"6": 20,
		"7": 10,
	}
	usersProbabilities      = map[interface{}]uint64{}
	currenciesProbabilities = map[interface{}]uint64{
		"UAH": 50,
		"USD": 30,
		"EUR": 20,
	}
	yearsProbabilities = map[interface{}]uint64{
		2018: 10,
		2019: 30,
		2020: 60,
	}
)

func init() {
	for i := 0; i <= 1000; i++ {
		usersProbabilities[fmt.Sprintf("exampleUser%d", i)] = 1
	}
}

// ExampleCSVRow describe example csv rows
type ExampleCSVRow struct {
	OrderID  string
	Date     time.Time
	UserID   string
	Amount   float64
	Currency string
}

// NewExampleCSVRow return new example csv row
func NewExampleCSVRow() ExampleCSVRow {
	year := rand.WeightIndex(yearsProbabilities).(int)
	month := time.Month(rand.FastRand(1, 12))
	day := int(rand.FastRand(1, uint32(GetMaxDays(year, month))))
	hour := int(rand.FastRand(0, 24))

	return ExampleCSVRow{
		OrderID:  rand.WeightIndex(ordersProbabilities).(string),
		UserID:   rand.WeightIndex(usersProbabilities).(string),
		Amount:   rand.FastRandFloat64(100, 1, 100000),
		Currency: rand.WeightIndex(currenciesProbabilities).(string),
		Date:     time.Date(year, month, day, hour, 0, 0, 0, time.UTC),
	}
}

// GetMaxDays return days in moth
func GetMaxDays(year int, month time.Month) int {
	if month == 12 {
		month = 0
	}
	t := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC)
	return t.Day()
}

// CSV return csv row for given data
func (e ExampleCSVRow) CSV() []string {
	return []string{e.OrderID, e.Date.String(), e.UserID, strconv.FormatFloat(e.Amount, 'f', -1, 64), e.Currency}
}

// GenerateExampleCSV generate example file with random data
func GenerateExampleCSV() (err error) {
	var c *Config
	c, err = GetExampleCSVGeneratorConfigFromEnv()
	if err != nil {
		return
	}

	records := make(chan ExampleCSVRow, buffer)
	go func(records chan ExampleCSVRow) {
		for i := 0; i < c.Count; i++ {
			records <- NewExampleCSVRow()
		}
		close(records)
	}(records)

	var file *os.File
	file, err = os.Create(c.Output)
	if err != nil {
		return
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			if err != nil {
				err = errors.Wrap(err, closeErr.Error())
			} else {
				err = closeErr
			}
		}
	}()

	w := csv.NewWriter(file)
	var i int
	for record := range records {
		if err = w.Write(record.CSV()); err != nil {
			return
		}

		i++
		if i%freq == 0 {
			w.Flush()
			if err = w.Error(); err != nil {
				return
			}
			i = 0
		}
	}
	if i > 0 {
		w.Flush()
		if err = w.Error(); err != nil {
			return
		}
	}

	return
}
