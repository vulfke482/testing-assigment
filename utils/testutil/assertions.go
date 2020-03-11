package testutil

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

// TestingAssertEqual will check if the two objects are the same
// If not, it will raise an error
func TestingAssertEqual(t *testing.T, actual, expected interface{}) {
	t.Helper()
	switch actual.(type) {
	case error:
		expectedErr, _ := expected.(error)
		TestingAssertSameError(t, actual.(error), expectedErr)
	case float64:
		TestingAssertFloat(t, actual.(float64), expected.(float64))
	default:
		if !TestingIsEqual(actual, expected) {
			t.Fatalf("mismatched values: %v != %v", actual, expected)
		}
	}
}

// TestingIsEqual compares 2 object is deeply equal or not
func TestingIsEqual(actual, expected interface{}) bool {
	return reflect.DeepEqual(actual, expected)
}

// TestingNoError call fatal if err no nil
func TestingNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

// TestingContains this is wrapper for assert.Contains
func TestingContains(t *testing.T, actual, expected interface{}) bool {
	t.Helper()
	return assert.Contains(t, expected, actual)
}

// TestingIsSameError checks if the given errors are either same, or similar
func TestingIsSameError(received, expected error) bool {
	return errors.Cause(received) == errors.Cause(expected)
}

// TestingAssertSameError checks the two errors are same/similar
// If not it will raise an error
func TestingAssertSameError(t *testing.T, received, expected error) {
	t.Helper()
	if !TestingIsSameError(received, expected) {
		t.Fatalf("mismatched errors, received: %+v expected: %+v", received, expected)
	}
}

// TestingAssertFloat asserts the two float values are similar
// It will round the float to 6 decimal places
func TestingAssertFloat(t *testing.T, received, expected float64) {
	t.Helper()
	if !TestingIsSameFloat(received, expected) {
		t.Fatalf("mismatched float, received: %f, expected: %f", received, expected)
	}
}

// TestingIsSameFloat compares the two float values are similar enough
// It will round the float to 6 decimal places
func TestingIsSameFloat(received, expected float64) bool {
	return fmt.Sprintf("%.6f", received) == fmt.Sprintf("%.6f", expected)
}

// TestingAssertIsSameErrorMsg tests if the error message string is same for the two errors
// It will handle nil errors
func TestingAssertIsSameErrorMsg(t *testing.T, received, expected error) {
	t.Helper()
	// received and expected can be nil
	recMsg := fmt.Sprintf("%v", received)
	expMsg := fmt.Sprintf("%v", expected)
	if recMsg != expMsg {
		t.Fatalf("mismatched errors, received: %v, expected: %v", recMsg, expMsg)
	}
}

// TestingParseJSON parses the json for testing
func TestingParseJSON(t *testing.T, input string) interface{} {
	t.Helper()
	i := map[string]interface{}{}
	if err := json.Unmarshal([]byte(input), &i); err != nil {
		t.Fatalf("Unable to parse json: %s", input)
	}
	return i
}

// TestingAssertJSON tests the two nested jsons to be equal
// We expecte the json to be a map[string]interface{} at the top level
func TestingAssertJSON(t *testing.T, received, expected string) {
	t.Helper()
	if expected == "" {
		if received != "" {
			t.Fatalf("expected empty string, got: %s", received)
		}
		return
	}
	rec := TestingParseJSON(t, received)
	exp := TestingParseJSON(t, expected)
	TestingAssertEqual(t, rec, exp)
}

// TestingAssertJSONSlice tests two JSON`s slices and ignore order
func TestingAssertJSONSlice(t *testing.T, received, expected string) {
	t.Helper()
	if expected == "" {
		if received != "" {
			t.Fatalf("expected empty string, got: %s", received)
		}
		return
	}
	receivedSlice := []interface{}{}
	if err := json.Unmarshal([]byte(received), &receivedSlice); err != nil {
		t.Fatalf("Unable to parse json: %s", received)
	}
	expectedSlice := []interface{}{}
	if err := json.Unmarshal([]byte(expected), &expectedSlice); err != nil {
		t.Fatalf("Unable to parse json: %s", expected)
	}

	if len(receivedSlice) != len(expectedSlice) {
		t.Fatalf("mismatched values: %v != %v", received, expected)
	}

	var flag bool
	for _, expected := range expectedSlice {
		flag = false
		for _, received := range receivedSlice {
			if TestingIsEqual(received, expected) {
				flag = true
				break
			}
		}

		if !flag {
			break
		}
	}

	if !flag {
		t.Fatalf("mismatched values: %v != %v", received, expected)
	}
}
