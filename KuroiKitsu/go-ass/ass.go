package ass

import (
	"errors"
	"io"
)

// ASS wraps the content of a `.ass` file.
type ASS struct {
	Sections []*Section
}

// Section returns the first section with that name, or nil.
func (ass *ASS) Section(name string) *Section {
	if len(ass.Sections) == 0 {
		return nil
	}

	for _, s := range ass.Sections {
		if s.Name == name {
			return s
		}
	}

	return nil
}

// Dump dumps the ASS data.
func (ass *ASS) Dump(w io.Writer) (err error) {
	if len(ass.Sections) == 0 {
		err = errors.New("no sections")
		return
	}

	for i, section := range ass.Sections {
		if i > 0 {
			w.Write([]byte{'\n'})
		}

		sectionName := section.String()
		rawSectionName := []byte(sectionName)

		w.Write(rawSectionName)
		w.Write([]byte{'\n'})

		if len(section.Pairs) == 0 {
			continue
		}

		for _, pair := range section.Pairs {
			pairStr := pair.String()
			rawPair := []byte(pairStr)

			w.Write(rawPair)
			w.Write([]byte{'\n'})
		}
	}

	return
}

// addSection appends a new section.
func (ass *ASS) addSection(name string) {
	if ass.Sections == nil {
		ass.Sections = make([]*Section, 0)
	} else {

		for _, s2 := range ass.Sections {
			if s2.Name == name {
				return
			}
		}

	}

	s := &Section{
		Name: name,
	}

	ass.Sections = append(ass.Sections, s)
}

// lastSections returns the last section.
//
// When no section is found, ok is false.
func (ass *ASS) lastSection() (section *Section, ok bool) {
	ok = false

	if len(ass.Sections) == 0 {
		return
	}

	section = ass.Sections[len(ass.Sections)-1]
	ok = true

	return
}

// New creates and initializes an *ASS.
func New() *ASS {
	ass := new(ASS)
	ass.Sections = make([]*Section, 0)

	return ass
}
