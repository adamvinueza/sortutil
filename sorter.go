package sortutil

import (
	"fmt"
	"reflect"
	"sort"
	"time"
)

const (
	stringType = "strings"
	intType    = "ints"
	uintType   = "uints"
	boolType   = "booleans"
	floatType  = "floats"
	timeType   = "type time.Time"
)

func invalidOrderForType(o Order, dtype string) error {
	return fmt.Errorf("invalid order %s for %s", o, dtype)
}

var (
	time_type = reflect.TypeOf(time.Time{})
)

type sorter struct {
	slice    reflect.Value
	getter   getter
	order    Order
	itemType reflect.Type
	vals     []reflect.Value
	valKind  reflect.Kind
	valType  reflect.Type
}

func (s *sorter) Len() int {
	return len(s.vals)
}

func (s *sorter) Swap(i, j int) {
	x := s.slice.Index(i)
	y := s.slice.Index(j)
	tmp := reflect.New(s.itemType).Elem()
	tmp.Set(x)
	x.Set(y)
	y.Set(tmp)
}

// Sort sorts its slice of structs by passing the sorter into an appropriate
// struct implementing sort.Interface, and calling sort.Sort on it.
//
// For example, to sort a struct by a field whose type is int64, Sort passes the
// sorter to a struct implementing a Less function comparing int64 values.
//
// Any issues are passed along via the error.
func (s *sorter) Sort() (err error) {
	// Sorting structs use reflect functions to retrieve the relevant values as
	// their presumed types; those functions panic if the values aren't of those
	// types. To keep the calling program from crashing, we'll recover and send
	// back the error that caused the panic.
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	if s.slice.Len() < 2 {
		return
	}
	if s.getter == nil {
		s.getter = simpleGetter()
	}
	s.itemType = s.slice.Index(0).Type()
	if s.vals, err = s.getter(s.slice); err != nil {
		return err
	}
	one := s.vals[0]
	s.valType = one.Type()
	s.valKind = one.Kind()
	switch s.valKind {
	default:
		switch s.valType {
		default:
			return fmt.Errorf("cannot sort by type %v", s.valType)
		case time_type:
			switch s.order {
			default:
				return invalidOrderForType(s.order, timeType)
			case ascending:
				sort.Sort(timeAscending{s})
			case descending:
				sort.Sort(timeDescending{s})
			}
		}
	case reflect.String:
		switch s.order {
		default:
			return invalidOrderForType(s.order, stringType)
		case ascending:
			sort.Sort(stringAscending{s})
		case descending:
			sort.Sort(stringDescending{s})
		case ascendingCaseInsensitive:
			sort.Sort(stringInsensitiveAscending{s})
		case descendingCaseInsensitive:
			sort.Sort(stringInsensitiveDescending{s})
		}
	case reflect.Bool:
		switch s.order {
		default:
			return invalidOrderForType(s.order, boolType)
		case ascending:
			sort.Sort(boolAscending{s})
		case descending:
			sort.Sort(boolDescending{s})
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		switch s.order {
		default:
			return invalidOrderForType(s.order, intType)
		case ascending:
			sort.Sort(intAscending{s})
		case descending:
			sort.Sort(intDescending{s})
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		switch s.order {
		default:
			return invalidOrderForType(s.order, uintType)
		case ascending:
			sort.Sort(intAscending{s})
		case descending:
			sort.Sort(intDescending{s})
		}
	case reflect.Float32, reflect.Float64:
		switch s.order {
		default:
			return invalidOrderForType(s.order, floatType)
		case ascending:
			sort.Sort(floatAscending{s})
		case descending:
			sort.Sort(floatDescending{s})
		}
	}
	return nil
}