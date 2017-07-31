package ass_test

import (
	"testing"

	"github.com/KuroiKitsu/go-ass"
)

func TestSection_IsValid(t *testing.T) {
	s := new(ass.Section)

	cases := map[string]bool{
		"":             false,
		"Events":       true,
		"Events]":      false,
		"[Events":      false,
		"[Events]":     false,
		"Events\n":     false,
		"Script Info":  true,
		"Script\nInfo": false,
	}

	for name, expectedValue := range cases {
		s.Name = name

		if value := s.IsValid(); value != expectedValue {
			var validStr string

			if expectedValue == true {
				validStr = "a valid"
			} else {
				validStr = "an invalid"
			}

			t.Fatalf("%#v is %s name, expecting Section.IsValid() to be %#v (got %#v)", name, validStr, expectedValue, value)
		}
	}
}
