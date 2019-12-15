package sortutil

import (
	"fmt"
	"reflect"
	"strings"
)

type getter func(reflect.Value) ([]reflect.Value, error)

func valueSlice(l int) []reflect.Value {
	s := make([]reflect.Value, l, l)
	return s
}

func toArray(slice reflect.Value, g func(reflect.Value) reflect.Value) (vals []reflect.Value, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
			vals = make([]reflect.Value, 0)
		}
	}()
	vals = valueSlice(slice.Len())
	for i := range vals {

		if v := g(slice.Index(i)); v.IsValid() {
			vals[i] = v
		} else {
			return vals, fmt.Errorf("reflected value is invalid")
		}
	}
	return vals, nil
}

func simpleGetter() getter {
	return func(s reflect.Value) (vals []reflect.Value, err error) {
		return toArray(s, func(r reflect.Value) reflect.Value {
			return reflect.Indirect(reflect.Indirect(r))
		})
	}
}

func handleFieldError(s reflect.Value, name string, e error) error {
	first := s.Index(0)
	t := first.Type()
	names := []string{}
	for i := 0; i < first.NumField(); i++ {
		names = append(names, fmt.Sprintf(`"%s"`, t.Field(i).Name))
	}
	return fmt.Errorf(
		`error getting values for field "%s" from struct with fields %s: %s`,
		name,
		strings.Join(names, ","),
		e.Error(),
	)

}

func fieldGetter(name string) getter {
	return func(s reflect.Value) (vals []reflect.Value, err error) {
		v, e := toArray(s, func(r reflect.Value) reflect.Value {
			return reflect.Indirect(reflect.Indirect(r).FieldByName(name))
		})
		if e != nil {
			e = handleFieldError(s, name, e)
		}
		return v, e
	}
}

func ciFieldGetter(name string) getter {
	return func(s reflect.Value) (vals []reflect.Value, err error) {
		v, e := toArray(s, func(r reflect.Value) reflect.Value {
			v := reflect.Indirect(r).FieldByNameFunc(
				func(n string) bool {
					return strings.EqualFold(n, name)
				})
			return reflect.Indirect(v)
		})
		if e != nil {
			e = handleFieldError(s, name, e)
		}
		return v, e
	}
}

func indexGetter(index int) getter {
	return func(s reflect.Value) (vals []reflect.Value, err error) {
		return toArray(s, func(r reflect.Value) reflect.Value {
			return reflect.Indirect(reflect.Indirect(r).Index(index))
		})
	}
}
