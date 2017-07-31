package ass_test

import (
	"strings"
	"testing"

	"github.com/KuroiKitsu/go-ass"
)

func TestPair_FormatHas(t *testing.T) {
	pair := getPair()

	formatKeys := []string{
		"Layer",
		"Start",
		"End",
		"Style",
		"Name",
		"MarginL",
		"MarginR",
		"MarginV",
		"Effect",
		"Text",
	}

	for _, fkey := range formatKeys {
		if !pair.FormatHas(fkey) {
			t.Fatalf("format has a %#v key, but FormatHas is returning false", fkey)
		}

		lowerFKey := strings.ToLower(fkey)
		if pair.FormatHas(lowerFKey) {
			t.Fatalf("format does not have a %#v key, but FormatHas is returning true", fkey)
		}
	}
}

func TestPair_KeyIsValid(t *testing.T) {
	pair := getPair()

	if !pair.KeyIsValid() {
		t.Fatalf("%#v is a valid key, but KeyIsValid is returning false", pair.Key)
	}

	pair.Key = pair.Key + ":"
	if pair.KeyIsValid() {
		t.Fatalf("%#v is an invalid key, but KeyIsValid is returning true", pair.Key)
	}
}

func TestPair_ValueIsValid(t *testing.T) {
	pair := getPair()

	if !pair.ValueIsValid() {
		t.Fatalf("%#v is a valid value, but ValueIsValid is returning false", pair.Value)
	}

	// This value is invalid, since the format says there are 10 tokens
	// but the value has only 9.
	//
	// ------> "0,0:00:30.92,0:00:32.42,main,,0,0,0,,It's morning."
	pair.Value = "0:00:30.92,0:00:32.42,main,,0,0,0,,It's morning."

	if pair.ValueIsValid() {
		t.Fatalf("%#v is an invalid value, but ValueIsValid is returning true", pair.Value)
	}

	pair.Format = nil

	if !pair.ValueIsValid() {
		t.Fatalf("value %#v should be valid when the pair has no format", pair.Value)
	}

	pair = getFormatPair()
	if !pair.ValueIsValid() {
		t.Fatalf("value %#v should be valid", pair.Value)
	}
}

func TestPair_Get(t *testing.T) {
	pair := getPair()

	cases := map[string]string{
		"Start": "0:00:30.92",
		"End":   "0:00:32.42",
		"Text":  "It's morning.",
	}

	for fkey, expectedValue := range cases {
		if val := pair.Get(fkey); val != expectedValue {
			t.Fatalf("invalid token for %#v (expecting %#v, got %#v)", fkey, expectedValue, val)
		}
	}
}

func TestPair_Set(t *testing.T) {
	var (
		expectingText string
	)

	pair := getPair()

	expectingText = "It's morning."

	if text := pair.Get("Text"); text != expectingText {
		t.Fatalf("invalid Text (expecting %#v, got %#v)", expectingText, text)
	}

	expectingText = "Onii-chan!"
	pair.Set("Text", expectingText)

	if text := pair.Get("Text"); text != expectingText {
		t.Fatalf("invalid Text (expecting %#v, got %#v)", expectingText, text)
	}
}

func getPair() *ass.Pair {
	return &ass.Pair{
		Key:   "Dialogue",
		Value: "0,0:00:30.92,0:00:32.42,main,,0,0,0,,It's morning.",
		Format: []string{
			"Layer",
			"Start",
			"End",
			"Style",
			"Name",
			"MarginL",
			"MarginR",
			"MarginV",
			"Effect",
			"Text",
		},
	}
}

func getFormatPair() *ass.Pair {
	return &ass.Pair{
		Key:    "Format",
		Value:  "Layer, Start, End, Style, Name, MarginL, MarginR, MarginV, Effect, Text",
		Format: nil,
	}
}
