package sortutil

import (
	"math"
	"reflect"
	"sort"
	"strings"
	"time"
)

// type-specific sorters
type stringAscending struct{ *sorter }
type stringDescending struct{ *sorter }
type stringInsensitiveAscending struct{ *sorter }
type stringInsensitiveDescending struct{ *sorter }
type boolAscending struct{ *sorter }
type boolDescending struct{ *sorter }
type intAscending struct{ *sorter }
type intDescending struct{ *sorter }
type uintAscending struct{ *sorter }
type uintDescending struct{ *sorter }
type floatAscending struct{ *sorter }
type floatDescending struct{ *sorter }
type timeAscending struct{ *sorter }
type timeDescending struct{ *sorter }
type reverser struct{ *sorter }

// Returns true if string i is less than string j.
func (s stringAscending) Less(i, j int) bool {
	return s.sorter.vals[i].String() < s.sorter.vals[j].String()
}

// Returns true if string j is less than string i.
func (s stringDescending) Less(i, j int) bool {
	return s.sorter.vals[i].String() > s.sorter.vals[j].String()
}

// Returns true if, ignoring case, string i is less than string j.
func (s stringInsensitiveAscending) Less(i, j int) bool {
	return strings.ToLower(s.sorter.vals[i].String()) < strings.ToLower(s.sorter.vals[j].String())
}

// Returns true if, ignoring case, string j is less than string i.
func (s stringInsensitiveDescending) Less(i, j int) bool {
	return strings.ToLower(s.sorter.vals[i].String()) > strings.ToLower(s.sorter.vals[j].String())
}

// Returns true if bool i is false and bool j is true.
func (s boolAscending) Less(i, j int) bool {
	return !s.sorter.vals[i].Bool() && s.sorter.vals[j].Bool()
}

// Returns true if bool j is false and bool i is true.
func (s boolDescending) Less(i, j int) bool {
	return s.sorter.vals[i].Bool() && !s.sorter.vals[j].Bool()
}

// Returns true if int i is less than int j.
func (s intAscending) Less(i, j int) bool {
	return s.sorter.vals[i].Int() < s.sorter.vals[j].Int()
}

// Returns true if int j is less than int i.
func (s intDescending) Less(i, j int) bool {
	return s.sorter.vals[i].Int() > s.sorter.vals[j].Int()
}

// Returns true if uint i is less than uint j.
func (s uintAscending) Less(i, j int) bool {
	return s.sorter.vals[i].Uint() < s.sorter.vals[j].Uint()
}

// Returns true if uint j is less than uint i.
func (s uintDescending) Less(i, j int) bool {
	return s.sorter.vals[i].Uint() > s.sorter.vals[j].Uint()
}

// Returns true if float i is less than float j, or only float i isNaN.
func (s floatAscending) Less(i, j int) bool {
	a := s.sorter.vals[i].Float()
	b := s.sorter.vals[j].Float()
	return a < b || math.IsNaN(a) && !math.IsNaN(b)
}

// Returns true if float j is less than float i, or only float j isNaN.
func (s floatDescending) Less(i, j int) bool {
	a := s.sorter.vals[i].Float()
	b := s.sorter.vals[j].Float()
	return a > b || !math.IsNaN(a) && math.IsNaN(b)
}

// Returns true if time i is before time j.
func (s timeAscending) Less(i, j int) bool {
	return s.sorter.vals[i].Interface().(time.Time).Before(s.sorter.vals[j].Interface().(time.Time))
}

// Returns true if time j is before time i.
func (s timeDescending) Less(i, j int) bool {
	return s.sorter.vals[i].Interface().(time.Time).After(s.sorter.vals[j].Interface().(time.Time))
}

// Returns the number of items in the slice to be reversed.
func (s reverser) Len() int {
	return s.sorter.slice.Len()
}

// Less does nothing useful, as reverse does not actually sort.
func (s reverser) Less(i, j int) bool {
	return i < j
}

// Returns a sorter with the specified getter and order for the specified slice.
func new(slice interface{}, gtr getter, ord Order) *sorter {
	v := reflect.ValueOf(slice)
	return &sorter{
		slice:  v,
		getter: gtr,
		order:  ord,
	}
}

// Sort the specified slice in the specified order, using the specified getter.
func Sort(slice interface{}, gtr getter, ord Order) error {
	return new(slice, gtr, ord).Sort()
}

// Asc sorts the specified slice in ascending order.
func Asc(slice interface{}) error {
	return new(slice, nil, Ascending).Sort()
}

// Desc sorts the specified slice in descending order.
func Desc(slice interface{}) error {
	return new(slice, nil, Descending).Sort()
}

// CiAsc the specified slice in ascending order, ignoring case.
func CiAsc(slice interface{}) error {
	return new(slice, nil, AscendingCaseInsensitive).Sort()
}

// CiDesc sorts the specified slice in descending order, ignoring case.
func CiDesc(slice interface{}) error {
	return new(slice, nil, DescendingCaseInsensitive).Sort()
}

// ByField sorts the specified slice of structs in the specified order, by the
// values in the named field.
func ByField(slice interface{}, name string, ord Order) error {
	return new(slice, fieldGetter(name), ord).Sort()
}

// ByCiField sorts the specified slice of structs in the specified order, by the
// values in the named field. The named field is searched ignoring case.
func ByCiField(slice interface{}, name string, ord Order) error {
	return new(slice, ciFieldGetter(name), ord).Sort()
}

// ByIndex sorts the specified slice of structs in the specified order, by the
// values in the indexed field.
func ByIndex(slice interface{}, index int, ord Order) error {
	return new(slice, indexGetter(index), ord).Sort()
}

// Reverse reverses the order of items in the specified slice.
func Reverse(slice interface{}) {
	s := reverser{new(slice, nil, 0)}
	if s.Len() < 2 {
		return
	}
	s.itemType = s.slice.Index(0).Type()
	reverseInterface(s)
}

func reverseInterface(s sort.Interface) {
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		s.Swap(i, j)
	}
}

// SortReverseInterface sorts and reverses a sortable thing.
func SortReverseInterface(s sort.Interface) {
	sort.Sort(s)
	reverseInterface(s)
}
