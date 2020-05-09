package product

import (
	"strconv"
	"time"
)

// Timestamp struct
type Timestamp struct {
	time.Time
}

// ToString to time func
func (ts Timestamp) String() string {

	if ts.IsZero() {
		return "unknown"
	}
	// Mon Jan 2 15:04:05 -0700 MST 2006

	const layout = "2006/01"
	return ts.Format(layout)

}

// ToTimestamp func
func ToTimestamp(v interface{}) (ts Timestamp) {
	var t int
	switch v := v.(type) { // v starts as an empty interface
	case int:
		t = v // v becomes int var
	case string:
		t, _ = strconv.Atoi(v) // v becomes string var
	}

	ts.Time = time.Unix(int64(t), 0)
	return ts
}
