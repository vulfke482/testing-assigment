package src

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)


func generateData() []InputRecord {
	return []InputRecord{
		{
			"4",
			time.Date(2019, 7, 5, 5, 0, 0, 0, time.UTC),
			"exampleUser504",
			5,
			"USD",
		},
		{
			"4",
			time.Date(2019, 7, 5, 5, 0, 0, 0, time.UTC),
			"exampleUser504",
			5,
			"USD",
		},
		{
			"4",
			time.Date(2018, 7, 5, 5, 0, 0, 0, time.UTC),
			"exampleUser504",
			5,
			"USD",
		},
		{
			"4",
			time.Date(2019, 7, 5, 5, 0, 0, 0, time.UTC),
			"exampleUser504",
			5,
			"UAH",
		},
		{
			"4",
			time.Date(2019, 7, 5, 5, 0, 0, 0, time.UTC),
			"exampleUser504",
			6,
			"UAH",
		},
	}
}

func dateFilterResult() []InputRecord {
	return []InputRecord{
		{
			"4",
			time.Date(2019, 7, 5, 5, 0, 0, 0, time.UTC),
			"exampleUser504",
			5,
			"USD",
		},
		{
			"4",
			time.Date(2019, 7, 5, 5, 0, 0, 0, time.UTC),
			"exampleUser504",
			5,
			"USD",
		},
		{
			"4",
			time.Date(2019, 7, 5, 5, 0, 0, 0, time.UTC),
			"exampleUser504",
			5,
			"UAH",
		},
		{
			"4",
			time.Date(2019, 7, 5, 5, 0, 0, 0, time.UTC),
			"exampleUser504",
			6,
			"UAH",
		},
	}
}

func TestDateFilter(t *testing.T) {
	data := generateData()

	filteredData, err := DateFilter(data, time.Date(2019, 6, 5, 5, 0, 0, 0, time.UTC))

	if err != nil {
		fmt.Println("All it bed")
		t.Error()
	}

	neededResult := dateFilterResult()

	if !reflect.DeepEqual(filteredData, neededResult) {
		t.Error()
	} else {
		t.Logf("TestDataFilter passed.")
	}

}

func TestAggregateAmount(t *testing.T) {
	t.Log()
}

func TestProcessData(t *testing.T) {
	t.Log()
}
