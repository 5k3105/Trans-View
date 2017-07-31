package ass

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

var (
	fullTimeRegexp    = regexp.MustCompile(`^(((?P<hours>[0-9]{1,2}):)?(?P<minutes>[0-9]{1,2}):)?(?P<seconds>[0-9]{1,2})(\.(?P<milliseconds>[0-9]{1,3}))?`)
	secondsTimeRegexp = regexp.MustCompile(`^(?P<seconds>[0-9]+)(\.(?P<milliseconds>[0-9]+))?$`)
)

// Time is a moment in the video/audio, or a duration.
type Time int32

// String returns the textual representation of the time.
//
// You can use this value directly in the "Start" and "End" properties
// of dialogues.
func (t Time) String() string {
	sign := ""

	if t < 0 {
		sign = "-"
	}

	return fmt.Sprintf("%s%d:%02d:%02d.%02d", sign, t.Hour(), t.Minute(), t.Second(), t.Millisecond()/10)
}

// String3 is the same as String, but returns 3 decimals.
func (t Time) String3() string {
	sign := ""

	if t < 0 {
		sign = "-"
	}

	return fmt.Sprintf("%s%d:%02d:%02d.%03d", sign, t.Hour(), t.Minute(), t.Second(), t.Millisecond())
}

// Millisecond returns the millisecond part of the textual representation.
func (t Time) Millisecond() int {
	if t < 0 {
		t = -t
	}
	return int(t % 1000)
}

// Second returns the seconds part of the textual representation.
func (t Time) Second() int {
	if t < 0 {
		t = -t
	}
	return int(t/Second) % 60
}

// Minute returns the minute part of the textual representation.
func (t Time) Minute() int {
	if t < 0 {
		t = -t
	}
	return int(t/Minute) % 60
}

// Hour returns the hour part of the textual representation.
func (t Time) Hour() int {
	if t < 0 {
		t = -t
	}
	return int(t / Hour)
}

const (
	Millisecond Time = 1
	Second           = Millisecond * 1000
	Minute           = Second * 60
	Hour             = Minute * 60
)

// ParseTime returns a Time from its textual representation.
func ParseTime(rawTimestamp string) (t Time, err error) {
	var (
		m []string
	)

	if m = fullTimeRegexp.FindStringSubmatch(rawTimestamp); m != nil {
		t, err = buildTime(fullTimeRegexp.SubexpNames(), m)
		return
	}

	if m = secondsTimeRegexp.FindStringSubmatch(rawTimestamp); m != nil {
		t, err = buildTime(secondsTimeRegexp.SubexpNames(), m)
		return
	}

	err = errors.New("invalid timestamp")

	return
}

// buildTime returns a Time from values matched via regexp.
func buildTime(names, matches []string) (t Time, err error) {
	if len(names) != len(matches) {
		panic("`names` and `matches` have different lengths")
	}

	var (
		n int
	)

	t = 0

	for i := 0; i < len(names); i++ {
		if names[i] != "hours" && names[i] != "minutes" && names[i] != "seconds" && names[i] != "milliseconds" {
			continue
		}

		if matches[i] == "" {
			continue
		}

		if names[i] == "milliseconds" {
			for len(matches[i]) < 3 {
				matches[i] += "0"
			}
		}

		n, err = strconv.Atoi(matches[i])
		if err != nil {
			return
		}

		if names[i] == "hours" {
			t += Time(n) * Hour
		} else if names[i] == "minutes" {
			t += Time(n) * Minute
		} else if names[i] == "seconds" {
			t += Time(n) * Second
		} else if names[i] == "milliseconds" {
			t += Time(n) * Millisecond
		}
	}

	return
}
