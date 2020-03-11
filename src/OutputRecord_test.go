package src

import (
	"reflect"
	"testing"
)


func TestOutputRecord_CSV(t *testing.T) {
	expected := []string{"exampleUser504","47319.95","USD", "0"}

	outRecord := OutputRecord{
		"exampleUser504",
		47319.95,
		"USD",
		0,
	}

	result := outRecord.CSV()

	if !reflect.DeepEqual(expected, result) {
		t.Error()
	} else {
		t.Log()
	}

}