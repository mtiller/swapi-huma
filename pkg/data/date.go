package data

import (
	"encoding/json"
	"strings"
	"time"
)

type StrictDate time.Time

// Implement Marshaler and Unmarshaler interface
func (j *StrictDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*j = StrictDate(t)
	return nil
}

func (j StrictDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(j))
}
