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

func format(x interface{}) string {
	// &bookgame.Book{Title: "Moby Dick", Price: 20},
	if x == nil {
		return "unknown"

	}
	// &bookgame.Book{Title: "Astrix", Price: 40, Published: 234244979},
	var t int

	if i, ok := x.(int); ok {
		t = i
	}

	// &bookgame.Book{Title: "Tintin", Price: 40, Published: "123456959"},
	if i, ok := x.(string); ok {
		t, _ = strconv.Atoi(i)
	}

	u := time.Unix(int64(t), 0)
	return u.String()

}
