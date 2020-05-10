package product

import (
	"fmt"
	"sort"
	"strings"
)

// List type
type List []*Product

// Print func
func (l List) String() string {

	if len(l) == 0 {
		return "Sorry, We're waiting for a delivery.\n"
	}

	// - print
	var str strings.Builder
	for _, p := range l {
		str.WriteString("* ")
		str.WriteString(p.String())
		str.WriteRune('\n')
	}
	return str.String()
}

// Discount func
func (l List) Discount(discount float64) {
	fmt.Printf("Discounting by %v  \n", discount)

	// "it" here is Item
	for _, p := range l {
		p.Discount(discount)
	}
}

// by default `list` sorts by `title`.
func (l List) Len() int {
	return len(l)
}
func (l List) Less(i, j int) bool {
	return l[i].Title < l[j].Title
}
func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

// ByRelease order by release date
type ByRelease struct {
	List
}

// Less for ByRelease
func (b ByRelease) Less(i, j int) bool {
	return b.List[i].Released.Before(b.List[j].Released.Time)
}

// ByReleaseDate func
func ByReleaseDate(l List) sort.Interface {
	return &ByRelease{l}
}
