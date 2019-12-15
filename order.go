package sortutil

type Order int

const (
    ascending Order = iota
    descending
    ascendingCaseInsensitive
    descendingCaseInsensitive
)

func (o Order) String() string {
    orders := [...]string{
        "Ascending",
        "Descending",
        "AscendingCaseInsensitive",
        "DescendingCaseInsensitive",
    }
    if o < ascending || o > descendingCaseInsensitive {
        return "Unknown"
    }
    return orders[o]
}
