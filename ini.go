package ini

import (
	"bufio"
	"os"
	"strings"

	"github.com/grimdork/str"
)

// INI file base structure.
type INI struct {
	// Sections with settings.
	Sections map[string]*INISection
	// Order sections were loaded or added in.
	Order []string
}

const (
	// Bool type
	Bool = iota
	// Int type
	Int
	// INIFloat type
	Float
	// String type
	String
)

// New returns an empty INI structure.
func New() *INI {
	return &INI{
		Sections: make(map[string]*INISection),
	}
}

// Load INI from file and take a guess at the types of each value.
func Load(filename string) (*INI, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer f.Close()
	ini := New()
	r := bufio.NewReader(f)
	loop := true
	for loop {
		l, err := r.ReadString('\n')
		if err != nil {
			loop = false
		} else {
			l = l[:len(l)-1]
			// This automatically skips comments, and really anything else
			// unknown that isn't after the first section header.
			if strings.HasPrefix(l, "[") && strings.HasSuffix(l, "]") {
				name := l[1 : len(l)-1]
				s := ini.AddSection(name)
				s.parse(r)
			}
		}
	}
	return ini, err
}

// Save outputs the INI to a file.
// If tabbed is true, the fields will be saved with a tab character prepended.
func (ini *INI) Save(filename string, tabbed bool) error {
	b := str.NewStringer()
	count := 0
	for _, secname := range ini.Order {
		if count > 0 {
			b.WriteString("\n")
		}
		count++
		b.WriteStrings("[", secname, "]\n")
		for _, key := range ini.Sections[secname].Order {
			f := ini.Sections[secname].Fields[key]
			if tabbed {
				b.WriteRune('\t')
			}
			b.WriteStrings(key, "=", f.Value, "\n")
		}
	}
	return os.WriteFile(filename, []byte(b.String()), 0600)
}

// AddSection to INI structure.
func (ini *INI) AddSection(name string) *INISection {
	sec := &INISection{
		Fields: make(map[string]*INIField),
	}
	ini.Sections[name] = sec
	ini.Order = append(ini.Order, name)
	return sec
}
