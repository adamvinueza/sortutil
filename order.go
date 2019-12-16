package sortutil

// Order enumerates orders.
type Order int

//revive:disable
const (
	Ascending Order = iota
	Descending
	AscendingCaseInsensitive
	DescendingCaseInsensitive
)
//revive:enable

// The string representation of an order.
func (o Order) String() string {
	orders := [...]string{
		"Ascending",
		"Descending",
		"AscendingCaseInsensitive",
		"DescendingCaseInsensitive",
	}
	if o < Ascending || o > DescendingCaseInsensitive {
		return "Unknown"
	}
	return orders[o]
}
