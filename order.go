package sortutil

type Order int

const (
	Ascending Order = iota
	Descending
	AscendingCaseInsensitive
	DescendingCaseInsensitive
)

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
