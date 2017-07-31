package ass

import (
	"regexp"
	"strings"
)

// A backslash escapes the following character.
//
// The only exceptions are `\n` and `\N`, since they have
// a special meaning.
var escapingRegexp = regexp.MustCompile(`\\([^Nn])`)

// Text is the value of "Text" in a Dialogue line.
type Text string

// IsValid returns true if the text is valid.
func (text Text) IsValid() bool {
	// Text can't contain new line characters. New line characters.
	if strings.Contains(string(text), "\n") {
		return false
	}

	for _, token := range text.Split() {
		if len(token) < 3 {
			continue
		}
		if !strings.HasPrefix(token, "{") {
			continue
		}

		// Nested braces are considered invalid.
		if strings.Contains(token[1:len(token)-1], "{") {
			return false
		}
	}

	return true
}

// Split separates readable parts from style override tags.
func (text Text) Split() []string {
	breakpoints := make([]int, 0)
	breakpoints = append(breakpoints, 0)

	insideBraces := false
	escaping := false

	// Find break points.
	for i := 0; i < len(text); i++ {
		c := text[i]

		if escaping {
			escaping = false
			continue
		}

		if insideBraces {
			// A closing brace can *not* be escaped.
			if c == '}' {
				insideBraces = false
				breakpoints = append(breakpoints, i+1)
			}
			continue
		}

		if c == '\\' {
			escaping = true
			continue
		}

		if c == '{' {
			insideBraces = true

			// We might already have a breakpoint at this position. It
			// happens because the '{' is either at text[0] or after a
			// closing '}'.
			if len(breakpoints) > 0 && breakpoints[len(breakpoints)-1] == i {
				continue
			}

			breakpoints = append(breakpoints, i)
			continue
		}
	}

	if breakpoints[len(breakpoints)-1] != len(text) {
		breakpoints = append(breakpoints, len(text))
	}

	lastIndex := len(breakpoints) - 1
	parts := make([]string, lastIndex)

	for n := 0; n < lastIndex; n++ {
		i := breakpoints[n]
		j := breakpoints[n+1]

		parts[n] = string(text[i:j])
	}

	return parts
}

// Readable returns only the visible text, i.e. without style override
// tags.
func (text Text) Readable() string {
	readable := ""

	for _, chunk := range text.Split() {

		// Ignore style override tags.
		if strings.HasPrefix(chunk, "{") {
			continue
		}

		// Escape characters.
		readable += escapingRegexp.ReplaceAllString(chunk, "$1")
	}

	readable = strings.TrimSpace(readable)

	return readable
}
