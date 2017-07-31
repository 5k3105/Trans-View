package ass

import (
	"fmt"
	"strings"
)

// Section is a section.
type Section struct {
	Name  string
	Pairs []*Pair
}

// String returns a string in the form `[Name]`.
func (s *Section) String() string {
	return fmt.Sprintf("[%s]", s.Name)
}

// Get returns the value of the first pair with that key.
//
// This is only useful for the [Script Info] section.
func (s *Section) Get(key string) string {
	if len(s.Pairs) == 0 {
		return ""
	}

	for _, pair := range s.Pairs {
		if pair.Key == key {
			return pair.Value
		}
	}

	return ""
}

// Set changes the value of the first pair with the given key. If the
// section does not have any pair with that key, a new is automatically
// created.
//
// This is only useful for the [Script Info] section.
func (s *Section) Set(key, value string) {
	if len(s.Pairs) > 0 {
		for _, pair := range s.Pairs {
			if pair.Key == key {
				pair.Value = value
				return
			}
		}
	}

	s.Pairs = append(s.Pairs, &Pair{
		Key:   key,
		Value: value,
	})
}

// IsValid returns true if the section name is valid, false otherwise.
func (s *Section) IsValid() bool {
	if s.Name == "" {
		return false
	}
	if strings.Contains(s.Name, "[") {
		return false
	}
	if strings.Contains(s.Name, "]") {
		return false
	}
	if strings.Contains(s.Name, "\n") {
		return false
	}

	return true
}

// addPair appends a new pair to the section.
//
// If the section contains a format, it is included in the new pair.
func (s *Section) addPair(key, value string) {
	p := &Pair{
		Key:   key,
		Value: value,
	}

	if len(s.Pairs) > 0 && s.Pairs[0].Key == "Format" {
		p.setFormatString(s.Pairs[0].Value)
	}

	s.Pairs = append(s.Pairs, p)
}

// NewSection creates a new *Section.
func NewSection(name string) *Section {
	return &Section{
		Name:  name,
		Pairs: make([]*Pair, 0),
	}
}
