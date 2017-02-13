package macromeasures

import (
	"fmt"
	"strconv"
	"time"
)

// Time extends time.Time
type Time struct {
	Time time.Time
}

// MarshalJSON takes time.Time and turns into an int64 unix timestamp
func (t Time) MarshalJSON() ([]byte, error) {
	var unix int64
	if result := t.Time.Unix(); result > 0 {
		unix = result
	}
	return []byte(fmt.Sprintf(`"%d"`, unix)), nil
}

// UnmarshalJSON takes an int64 and turns into time.Time
func (t *Time) UnmarshalJSON(b []byte) error {
	str, err := strconv.Unquote(string(b))
	if err != nil {
		str = string(b)
	}
	unix, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return err
	}
	t.Time = time.Unix(unix, 0).UTC()
	return nil
}
