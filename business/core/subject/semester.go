package subject

import "errors"

// Set of possible semesters for a semester.
var (
	SemesterFirst  = Semester{"first"}
	SemesterSecond = Semester{"second"}
)

// Set of known semesters.
var semesters = map[string]Semester{
	SemesterFirst.name:  SemesterFirst,
	SemesterSecond.name: SemesterSecond,
}

// Semester represents a type of semester in the system.
type Semester struct {
	name string
}

// ParseSemester parses the string value and returns a semester if one exists.
func ParseSemester(value string) (Semester, error) {
	semester, exists := semesters[value]
	if !exists {
		return Semester{}, errors.New("invalid semester")
	}

	return semester, nil
}

// MustParseSemester parses the string value and returns a semeter if one exists. If
// an error occurs the function panics.
func MustParseSemester(value string) Semester {
	semester, err := ParseSemester(value)
	if err != nil {
		panic(err)
	}

	return semester
}

// Name returns the name of the type.
func (s Semester) Name() string {
	return s.name
}

// UnmarshalText implement the unmarshal interface for JSON conversions.
func (s *Semester) UnmarshalText(data []byte) error {
	s.name = string(data)
	return nil
}

// MarshalText implement the marshal interface for JSON conversions.
func (s Semester) MarshalText() ([]byte, error) {
	return []byte(s.name), nil
}

// Equal provides support for the go-cmp package and testing.
func (s Semester) Equal(s2 Semester) bool {
	return s.name == s2.name
}
