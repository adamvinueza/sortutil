package sortutil

import (
	"math"
	"reflect"
	"sort"
	"strings"
	"time"
)

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

func (s stringAscending) Less(i, j int) bool {
	return s.sorter.vals[i].String() < s.sorter.vals[j].String()
}

func (s stringDescending) Less(i, j int) bool {
	return s.sorter.vals[i].String() > s.sorter.vals[j].String()
}

func (s stringInsensitiveAscending) Less(i, j int) bool {
	return strings.ToLower(s.sorter.vals[i].String()) < strings.ToLower(s.sorter.vals[j].String())
}

func (s stringInsensitiveDescending) Less(i, j int) bool {
	return strings.ToLower(s.sorter.vals[i].String()) > strings.ToLower(s.sorter.vals[j].String())
}

func (s boolAscending) Less(i, j int) bool {
	return !s.sorter.vals[i].Bool() && s.sorter.vals[j].Bool()
}

func (s boolDescending) Less(i, j int) bool {
	return s.sorter.vals[i].Bool() && !s.sorter.vals[j].Bool()
}

func (s intAscending) Less(i, j int) bool {
	return s.sorter.vals[i].Int() < s.sorter.vals[j].Int()
}

func (s intDescending) Less(i, j int) bool {
	return s.sorter.vals[i].Int() > s.sorter.vals[j].Int()
}

func (s uintAscending) Less(i, j int) bool {
	return s.sorter.vals[i].Uint() < s.sorter.vals[j].Uint()
}

func (s uintDescending) Less(i, j int) bool {
	return s.sorter.vals[i].Uint() > s.sorter.vals[j].Uint()
}

func (s floatAscending) Less(i, j int) bool {
	a := s.sorter.vals[i].Float()
	b := s.sorter.vals[j].Float()
	return a < b || math.IsNaN(a) && !math.IsNaN(b)
}

func (s floatDescending) Less(i, j int) bool {
	a := s.sorter.vals[i].Float()
	b := s.sorter.vals[j].Float()
	return a > b || !math.IsNaN(a) && math.IsNaN(b)
}

func (s timeAscending) Less(i, j int) bool {
	return s.sorter.vals[i].Interface().(time.Time).Before(s.sorter.vals[j].Interface().(time.Time))
}

func (s timeDescending) Less(i, j int) bool {
	return s.sorter.vals[i].Interface().(time.Time).After(s.sorter.vals[j].Interface().(time.Time))
}

func (s reverser) Len() int {
	return s.sorter.slice.Len()
}

func (s reverser) Less(i, j int) bool {
	return i < j
}

func new(slice interface{}, gtr getter, ord Order) *sorter {
	v := reflect.ValueOf(slice)
	return &sorter{
		slice:  v,
		getter: gtr,
		order:  ord,
	}
}

func Sort(slice interface{}, gtr getter, ord Order) error {
	return new(slice, gtr, ord).Sort()
}

func Asc(slice interface{}) error {
	return new(slice, nil, ascending).Sort()
}

func Desc(slice interface{}) error {
	return new(slice, nil, descending).Sort()
}

func CiAsc(slice interface{}) error {
	return new(slice, nil, ascendingCaseInsensitive).Sort()
}

func CiDesc(slice interface{}) error {
	return new(slice, nil, descendingCaseInsensitive).Sort()
}

func ByField(slice interface{}, name string, ord Order) error {
	return new(slice, fieldGetter(name), ord).Sort()
}

func ByCiField(slice interface{}, name string, ord Order) error {
	return new(slice, ciFieldGetter(name), ord).Sort()
}

func ByIndex(slice interface{}, index int, ord Order) error {
	return new(slice, indexGetter(index), ord).Sort()
}

func Reverse(slice interface{}) {
	s := reverser{new(slice, nil, 0)}
	if s.Len() < 2 {
		return
	}
	s.itemType = s.slice.Index(0).Type()
	ReverseInterface(s)
}

func ReverseInterface(s sort.Interface) {
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		s.Swap(i, j)
	}
}

func SortReverseInterface(s sort.Interface) {
	sort.Sort(s)
	ReverseInterface(s)
}
