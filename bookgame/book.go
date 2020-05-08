package bookgame

import (
	"fmt"
	"strconv"
	"time"
)

// receiver name should be consistent with previous receiver name b for Book

// Book struct
type Book struct {
	Product
	Published interface{}
}

// Print func
func (b *Book) Print() {
	b.Product.Print()
	p := format(b.Published)
	fmt.Printf("\t - (%v) \n", p)
}

func format(v interface{}) string {
	var t int
	switch v := v.(type) { // v starts as an empty interface
	case int:
		t = v // v becomes int var
	case string:
		t, _ = strconv.Atoi(v) // v becomes string var
	default:
		return "unknown" // v stays as an empty interface
	}
	// Mon Jan 2 15:04:05 -0700 MST 2006
	const layout = "02/01/2006"
	u := time.Unix(int64(t), 0)
	return u.Format(layout)

}
