package fastlog

import (
	"bufio"
	"encoding/json"
	"os"
	"testing"

	"git.sansera.com/mtkach/golang-test-assignment/utils/testutil"
)

func TestLog(t *testing.T) {
	l, err := NewWrappedLogger("test", "./log", 0)
	if err != nil {
		t.Fatal(err)
	}
	l.Important("Test") // not logged
	l.Debug("Test")     // not logged
	l.Error("Test")
	l.Error("Test")

	file, err := os.Open("./log")
	if err != nil {
		t.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	var i int
	for scanner.Scan() {
		var d map[string]string
		err = json.Unmarshal(scanner.Bytes(), &d)
		if err != nil {
			t.Fatal(err)
		}

		if _, ok := d["logger"]; !ok {
			t.Fatal("Not found logger key")
		}
		i++
	}

	testutil.TestingAssertEqual(t, i, 2)
	testutil.TestingAssertEqual(t, scanner.Err(), nil)
	testutil.TestingAssertEqual(t, file.Close(), nil)
	testutil.TestingAssertEqual(t, os.Remove("./log"), nil)
}
