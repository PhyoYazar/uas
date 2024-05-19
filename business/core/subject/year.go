package subject

import "errors"

// Set of possible years for a mark.
var (
	YearFirst  = Year{"First Year"}
	YearSecond = Year{"Second Year"}
	YearThird  = Year{"Third Year"}
	YearFourth = Year{"Fourth Year"}
	YearFifth  = Year{"Fifth Year"}
	YearSixth  = Year{"Sixth Year"}
)

// Set of known years.
var years = map[string]Year{
	YearFirst.name:  YearFirst,
	YearSecond.name: YearSecond,
	YearThird.name:  YearThird,
	YearFourth.name: YearFourth,
	YearFifth.name:  YearFifth,
	YearSixth.name:  YearSixth,
}

// Year represents a type of year in the system.
type Year struct {
	name string
}

// ParseYear parses the string value and returns a year if one exists.
func ParseYear(value string) (Year, error) {
	year, exists := years[value]
	if !exists {
		return Year{}, errors.New("invalid year")
	}

	return year, nil
}

// MustParseYear parses the string value and returns a year if one exists. If
// an error occurs the function panics.
func MustParseYear(value string) Year {
	year, err := ParseYear(value)
	if err != nil {
		panic(err)
	}

	return year
}

// Name returns the name of the type.
func (y Year) Name() string {
	return y.name
}

// UnmarshalText implement the unmarshal interface for JSON conversions.
func (y *Year) UnmarshalText(data []byte) error {
	y.name = string(data)
	return nil
}

// MarshalText implement the marshal interface for JSON conversions.
func (y Year) MarshalText() ([]byte, error) {
	return []byte(y.name), nil
}

// Equal provides support for the go-cmp package and testing.
func (y Year) Equal(y2 Year) bool {
	return y.name == y2.name
}
