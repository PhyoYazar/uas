package mark

import "errors"

// Set of possible types for a mark.
var (
	TypeExam      = Type{"EXAM"}
	TypePractical = Type{"PRACTICAL"}
)

// Set of known types.
var types = map[string]Type{
	TypeExam.name:      TypeExam,
	TypePractical.name: TypePractical,
}

// Type represents a type of mark in the system.
type Type struct {
	name string
}

// ParseRole parses the string value and returns a type if one exists.
func ParseType(value string) (Type, error) {
	typ, exists := types[value]
	if !exists {
		return Type{}, errors.New("invalid type")
	}

	return typ, nil
}

// MustParseRole parses the string value and returns a type if one exists. If
// an error occurs the function panics.
func MustParseMarkType(value string) Type {
	typ, err := ParseType(value)
	if err != nil {
		panic(err)
	}

	return typ
}

// Name returns the name of the type.
func (t Type) Name() string {
	return t.name
}

// UnmarshalText implement the unmarshal interface for JSON conversions.
func (t *Type) UnmarshalText(data []byte) error {
	t.name = string(data)
	return nil
}

// MarshalText implement the marshal interface for JSON conversions.
func (t Type) MarshalText() ([]byte, error) {
	return []byte(t.name), nil
}

// Equal provides support for the go-cmp package and testing.
func (t Type) Equal(t2 Type) bool {
	return t.name == t2.name
}
