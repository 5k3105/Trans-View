package ass

import (
	"fmt"
	"strings"
)

// Pair is a key-value pair.
type Pair struct {
	Key   string
	Value string

	Format []string
}

// String returns a string in the form "Key: Value".
func (p *Pair) String() string {
	return fmt.Sprintf("%s: %s", p.Key, p.Value)
}

// setFormatString sets the format of the pair from the raw value of
// a Format line.
func (p *Pair) setFormatString(formatStr string) {
	components := strings.Split(formatStr, ",")
	for i, c := range components {
		components[i] = strings.TrimSpace(c)
	}

	p.Format = components
}

// FormatHas returns true when the pair format contains fkey.
//
// fkey is a component of the section format, e.g. Fontname, Start, End, Text... etc.
func (p *Pair) FormatHas(fkey string) bool {
	if fkey == "" {
		return false
	}

	for _, formatKey := range p.Format {
		if formatKey == fkey {
			return true
		}
	}

	return false
}

// IsValid returns true when both p.Key and p.Value are valid; false otherwise.
func (p *Pair) IsValid() bool {
	return p.KeyIsValid() && p.ValueIsValid()
}

// KeyIsValid returns true when p.Key is valid; false otherwise.
func (p *Pair) KeyIsValid() bool {
	if strings.Contains(p.Key, "\n") {
		return false
	}
	if strings.Contains(p.Key, ":") {
		return false
	}

	return true
}

// ValueIsValid returns true when p.Value is valid; false otherwise.
func (p *Pair) ValueIsValid() bool {
	if strings.Contains(p.Value, "\n") {
		return false
	}

	// No format means that anything is valid as long as there are no
	// invalid characters.
	if len(p.Format) == 0 {
		return true
	}

	// Value does not match format.
	if components := strings.SplitN(p.Value, ",", len(p.Format)); len(components) != len(p.Format) {
		return false
	}

	return true
}

// Get returns the fkey component of the value.
//
// fkey is a component of the section format, e.g. Fontname, Start, End, Text... etc.
//
// Example:
//
//     p.Get("Start") // "0:05:26.31"
func (p *Pair) Get(fkey string) string {
	var (
		idx int = -1
	)

	if len(p.Format) == 0 {
		return ""
	}

	for i, formatKey := range p.Format {
		if formatKey == fkey {
			idx = i
			break
		}
	}

	if idx == -1 {
		return ""
	}

	components := strings.SplitN(p.Value, ",", len(p.Format))
	if idx >= len(components) {
		return ""
	}

	return components[idx]
}

// Set changes the value associated with fkey and sets it to fvalue.
//
// fkey is a component of the section format, e.g. Fontname, Start, End, Text... etc.
//
// Example:
//
//     p.Set("Start", "0:05:26.31")
func (p *Pair) Set(fkey, fvalue string) bool {
	var (
		idx int = -1
	)

	components := strings.SplitN(p.Value, ",", len(p.Format))

	if len(p.Format) == 0 {
		return false
	}
	if len(components) != len(p.Format) {
		return false
	}

	for i, formatKey := range p.Format {
		if formatKey == fkey {
			idx = i
			break
		}
	}

	if idx == -1 {
		return false
	}

	if idx >= len(components) {
		return false
	}

	components[idx] = fvalue

	p.Value = strings.Join(components, ",")

	return true
}
