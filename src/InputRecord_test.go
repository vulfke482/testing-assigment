package src

import (
	"reflect"
	"testing"
	"time"
	"fmt"
)

func TestFromCSV(t *testing.T) {
	csv_example := []string{"4","2019-07-05 05:00:00 +0000 UTC","exampleUser504","47319.95","USD"}
	result, err := FromCSV(csv_example)

	if err != nil {
		fmt.Println("All it bed")
		t.Error()
	}

	expected := InputRecord{
		"4",
		time.Date(2019, 7, 5, 5, 0, 0, 0, time.UTC),
		"exampleUser504",
		47319.95,
		"USD",
	}

	if !reflect.DeepEqual(expected, result) {
		t.Error()
	} else {
		t.Logf("TestFromCsv passed")
	}
}

func TestInputRecord_CSV(t *testing.T) {
	expected := []string{"4","2019-07-05 05:00:00 +0000 UTC","exampleUser504","47319.95","USD"}

	tmp := InputRecord{
		"4",
		time.Date(2019, 7, 5, 5, 0, 0, 0, time.UTC),
		"exampleUser504",
		47319.95,
		"USD",
	}

	result := tmp.CSV()

	fmt.Println(expected, result)

	if !reflect.DeepEqual(expected, result) {
		t.Error()
	} else {
		t.Log("TestInputRecord passed")
	}

}