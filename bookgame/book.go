package bookgame

import (
	"fmt"
	m "goMethodsOOP/money"
	"strconv"
	"time"
)

// receiver name should be consistent with previous receiver name b for Book

// Book struct
type Book struct {
	Title     string
	Price     m.Money
	Published interface{}
}

// Print func
func (b *Book) Print() {
	p := format(b.Published)
	fmt.Printf("%-15s: %s - (%v)\n", b.Title, b.Price.ToString(), p)
}

// Discount price
func (b *Book) Discount(ratio float64) {
	b.Price *= m.Money(1 - ratio)
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
