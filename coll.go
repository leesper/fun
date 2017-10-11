package fun

import (
	"errors"
	"reflect"
)

// Remove removes elements from an array/slice, returns a new slice.
func Remove(series interface{}, removes ...interface{}) (interface{}, error) {
	st := reflect.TypeOf(series)
	sv := reflect.ValueOf(series)
	if st.Kind() != reflect.Array && st.Kind() != reflect.Slice {
		return nil, errors.New("not an array or slice")
	}

	filter := reflect.MakeMapWithSize(reflect.MapOf(st.Elem(), reflect.TypeOf(true)), len(removes))
	for _, r := range removes {
		filter.SetMapIndex(reflect.ValueOf(r), reflect.ValueOf(true))
	}

	removed := reflect.MakeSlice(reflect.SliceOf(st.Elem()), 0, 0)
	for i := 0; i < sv.Len(); i++ {
		if !filter.MapIndex(sv.Index(i)).IsValid() {
			removed = reflect.Append(removed, sv.Index(i))
		}
	}
	return removed.Interface(), nil
}

// Map accepts an array/slice and a function, call function on every element
// and returns a new slice of results.
func Map(series interface{}, f interface{}) (interface{}, error) {
	return nil, errors.New("not implemented")
}

// ForEach calls function f on every element of an array/slice.
func ForEach(series interface{}, f interface{}) error {
	return errors.New("not implemented")
}

// In checks whether an element is in collection.
func In(collection interface{}, element interface{}) bool {
	return false
}

// push pop splice
