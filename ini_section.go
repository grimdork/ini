package files

import (
	"bufio"
	"strings"
)

// INISection holds one or more fields.
type INISection struct {
	Fields map[string]*INIField
	// Order fields were loaded or added in.
	Order []string
}

// parse section properties until a new section or end of file.
func (s *INISection) parse(r *bufio.Reader) {
	loop := true
	for loop {
		next, err := r.Peek(2)
		// EOF
		if err != nil {
			return
		}

		// Skip blank lines
		if next[0] == '\n' {
			return
		}

		// New section, so this one's done
		if next[0] == '[' || next[1] == '[' {
			return
		}

		p, err := r.ReadString('\n')
		if err != nil {
			return
		}

		// Skip comments
		if strings.HasPrefix(p, "#") || strings.HasPrefix(p, ";") {
			continue
		}

		a := strings.SplitN(p, "=", 2)
		if a == nil || len(a) != 2 {
			return
		}

		a[0] = strings.TrimSpace(a[0])
		a[1] = strings.TrimSpace(a[1])
		switch a[1] {
		case "yes", "true", "1", "on", "no", "false", "0", "off":
			s.AddBool(a[0], boolValue(a[1]))
			return
		}

		// TODO: Figure out ints and floats.
		s.AddString(a[0], a[1])
	}
}

// GetBool returns a field as a bool, or the alternative.
func (s *INISection) GetBool(key string, alt bool) bool {
	v, ok := s.Fields[key]
	if !ok {
		return alt
	}

	return v.boolV
}

// AddBool adds a new bool field to the section.
func (s *INISection) AddBool(key string, value bool) {
	f := INIField{}
	f.SetBool(key, value)
	s.Fields[key] = &f
	s.Order = append(s.Order, key)
}

// GetInt returns a field as an int64, or the alternative.
func (s *INISection) GetInt(key string, alt int64) int64 {
	v, ok := s.Fields[key]
	if !ok {
		return alt
	}

	return v.intV
}

// AddInt adds a new int64 field to the section.
func (s *INISection) AddInt(key string, value int64) {
	f := INIField{}
	f.SetInt(key, value)
	s.Fields[key] = &f
	s.Order = append(s.Order, key)
}

// GetFloat returns a field as a float64, or the alternative.
func (s *INISection) GetFloat(key string, alt float64) float64 {
	v, ok := s.Fields[key]
	if !ok {
		return alt
	}

	return v.floatV
}

// AddFloat adds a new float64 field to the section.
func (s *INISection) AddFloat(key string, value float64) {
	f := INIField{}
	f.SetFloat(key, value)
	s.Fields[key] = &f
	s.Order = append(s.Order, key)
}

// GetString returns a field as a string, or the alternative.
func (s *INISection) GetString(key, alt string) string {
	v, ok := s.Fields[key]
	if !ok {
		return alt
	}

	return v.Value
}

// AddString adds a new string field to the section.
func (s *INISection) AddString(key string, value string) {
	f := INIField{}
	f.SetString(key, value)
	s.Fields[key] = &f
	s.Order = append(s.Order, key)
}

// boolValue from common strings.
func boolValue(s string) bool {
	switch s {
	case "yes", "true", "1", "on":
		return true
	}

	return false
}
