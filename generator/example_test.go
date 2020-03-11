package generator

import (
	"encoding/csv"
	"os"
	"testing"

	"git.sansera.com/mtkach/golang-test-assignment/utils/testutil"
)

func TestNewExampleCSVRow(t *testing.T) {
	r := NewExampleCSVRow()
	testutil.TestingAssertEqual(t, r.Amount > 0, true)
	testutil.TestingAssertEqual(t, len(r.UserID) >= 11, true)
	testutil.TestingAssertEqual(t, r.OrderID != "", true)
	testutil.TestingAssertEqual(t, r.Currency != "", true)
	testutil.TestingAssertEqual(t, r.Date.Unix() > 0, true)
	testutil.TestingAssertEqual(t, r.Date.Year() > 2017, true)
}

func TestGenerateExampleCSV(t *testing.T) {
	err := GenerateExampleCSV()
	testutil.TestingAssertEqual(t, err, nil)
	c, err := GetExampleCSVGeneratorConfigFromEnv()
	testutil.TestingAssertEqual(t, err, nil)
	f, err := os.Open(c.Output)
	testutil.TestingAssertEqual(t, err, nil)
	reader := csv.NewReader(f)
	rows, err := reader.ReadAll()
	testutil.TestingAssertEqual(t, err, nil)
	testutil.TestingAssertEqual(t, len(rows), c.Count)
	err = os.Remove(c.Output)
	testutil.TestingAssertEqual(t, err, nil)
}
