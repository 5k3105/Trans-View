package ass_test

import (
	"fmt"
	"testing"

	"github.com/KuroiKitsu/go-ass"
)

var expectedTimeValues = map[string]ass.Time{
	"0.1":         100 * ass.Millisecond,
	"0.12":        120 * ass.Millisecond,
	"0.123":       123 * ass.Millisecond,
	"1":           1 * ass.Second,
	"1.5":         1500 * ass.Millisecond,
	"1.234":       1234 * ass.Millisecond,
	"1:23":        1*ass.Minute + 23*ass.Second,
	"1:02.3":      1*ass.Minute + 2*ass.Second + 300*ass.Millisecond,
	"1:02.34":     1*ass.Minute + 2*ass.Second + 340*ass.Millisecond,
	"1:02.345":    1*ass.Minute + 2*ass.Second + 345*ass.Millisecond,
	"1:23.456":    1*ass.Minute + 23*ass.Second + 456*ass.Millisecond,
	"1:02:34":     1*ass.Hour + 2*ass.Minute + 34*ass.Second,
	"1:23:45":     1*ass.Hour + 23*ass.Minute + 45*ass.Second,
	"1:23:45.67":  1*ass.Hour + 23*ass.Minute + 45*ass.Second + 670*ass.Millisecond,
	"1:23:45.678": 1*ass.Hour + 23*ass.Minute + 45*ass.Second + 678*ass.Millisecond,
}

var expectedTimeStrings = map[ass.Time]string{
	1 * ass.Millisecond:    "0:00:00.00",
	9 * ass.Millisecond:    "0:00:00.00",
	10 * ass.Millisecond:   "0:00:00.01",
	99 * ass.Millisecond:   "0:00:00.09",
	100 * ass.Millisecond:  "0:00:00.10",
	110 * ass.Millisecond:  "0:00:00.11",
	111 * ass.Millisecond:  "0:00:00.11",
	999 * ass.Millisecond:  "0:00:00.99",
	1001 * ass.Millisecond: "0:00:01.00",

	1 * ass.Second:  "0:00:01.00",
	10 * ass.Second: "0:00:10.00",
	61 * ass.Second: "0:01:01.00",

	1 * ass.Minute:  "0:01:00.00",
	10 * ass.Minute: "0:10:00.00",
	11 * ass.Minute: "0:11:00.00",
	59 * ass.Minute: "0:59:00.00",
	61 * ass.Minute: "1:01:00.00",

	1 * ass.Hour: "1:00:00.00",
	9 * ass.Hour: "9:00:00.00",

	1*ass.Hour + 23*ass.Minute + 45*ass.Second + 678*ass.Millisecond: "1:23:45.67",
}

var expectedTimeStrings3 = map[ass.Time]string{
	1 * ass.Millisecond:    "0:00:00.001",
	9 * ass.Millisecond:    "0:00:00.009",
	10 * ass.Millisecond:   "0:00:00.010",
	99 * ass.Millisecond:   "0:00:00.099",
	100 * ass.Millisecond:  "0:00:00.100",
	110 * ass.Millisecond:  "0:00:00.110",
	111 * ass.Millisecond:  "0:00:00.111",
	999 * ass.Millisecond:  "0:00:00.999",
	1001 * ass.Millisecond: "0:00:01.001",

	1 * ass.Second:  "0:00:01.000",
	10 * ass.Second: "0:00:10.000",
	61 * ass.Second: "0:01:01.000",

	1 * ass.Minute:  "0:01:00.000",
	10 * ass.Minute: "0:10:00.000",
	11 * ass.Minute: "0:11:00.000",
	59 * ass.Minute: "0:59:00.000",
	61 * ass.Minute: "1:01:00.000",

	1 * ass.Hour: "1:00:00.000",
	9 * ass.Hour: "9:00:00.000",

	1*ass.Hour + 23*ass.Minute + 45*ass.Second + 678*ass.Millisecond: "1:23:45.678",
}

func ExampleTime() {
	t := 3*ass.Minute + 43*ass.Second

	fmt.Printf("Time: %s\n", t)
	fmt.Printf("Seconds: %d\n", t/ass.Second) // It's ass.Second, not time.Second
	// Output:
	// Time: 0:03:43.00
	// Seconds: 223
}

func ExampleParseTime() {
	t, err := ass.ParseTime("3:43")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Time: %s\n", t)
	fmt.Printf("Seconds: %d\n", t/ass.Second) // It's ass.Second, not time.Second
	// Output:
	// Time: 0:03:43.00
	// Seconds: 223
}

func TestTime(t *testing.T) {
	for rawTimestamp, expectedValue := range expectedTimeValues {
		ts, err := ass.ParseTime(rawTimestamp)
		if err != nil {
			t.Fatal(err)
		}
		if ts != expectedValue {
			t.Fatalf("invalid timestamp %d (expecting %d)", ts, expectedValue)
		}
	}

	for timeValue, expectedString := range expectedTimeStrings {
		var (
			timeStr string
		)

		timeStr = timeValue.String()
		if timeStr != expectedString {
			t.Fatalf("invalid time string %s (expecting %s)", timeStr, expectedString)
		}

		timeValue = -timeValue
		timeStr = timeValue.String()
		expectedString = "-" + expectedString
		if timeStr != expectedString {
			t.Fatalf("invalid time string %s (expecting %s)", timeStr, expectedString)
		}
	}

	for timeValue, expectedString := range expectedTimeStrings3 {
		var (
			timeStr string
		)

		timeStr = timeValue.String3()
		if timeStr != expectedString {
			t.Fatalf("invalid time string %s (expecting %s)", timeStr, expectedString)
		}

		timeValue = -timeValue
		timeStr = timeValue.String3()
		expectedString = "-" + expectedString
		if timeStr != expectedString {
			t.Fatalf("invalid time string %s (expecting %s)", timeStr, expectedString)
		}
	}

}
