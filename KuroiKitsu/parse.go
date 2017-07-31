package ass

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// Parse parses an io.Reader in ASS format.
func Parse(r io.Reader) (ass *ASS, err error) {
	ass = New()

	s := bufio.NewScanner(r)
	isFirstLine := true

	for s.Scan() {
		line := s.Text()

		if isFirstLine {
			isFirstLine = false

			// Unicode BOM.
			line = strings.TrimPrefix(line, "\ufeff")
		}

		// Ignore empty lines.
		if strings.TrimSpace(line) == "" {
			continue
		}

		// Comments.
		if strings.HasPrefix(line, ";") {
			continue
		}

		// Sections.
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			sectionName := line[1 : len(line)-1]
			if sectionName == "" {
				continue
			}

			ass.addSection(sectionName)

			continue
		}

		// Pairs (Format, Style, Dialogue...).
		if chunks := strings.SplitN(line, ": ", 2); len(chunks) == 2 {
			section, ok := ass.lastSection()
			if !ok {
				continue
			}

			key := chunks[0]
			value := chunks[1]

			section.addPair(key, value)

			continue
		}

		// Silently ignore unrecognized lines.
	}

	if err = s.Err(); err != nil {
		return
	}

	return
}

// ParseFile opens a file and parses it with Parse.
func ParseFile(file string) (ass *ASS, err error) {
	var (
		f *os.File
	)

	f, err = os.Open(file)
	if err != nil {
		return
	}
	defer f.Close()

	ass, err = Parse(f)
	return
}
