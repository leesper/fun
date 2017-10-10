package fun

import (
	"errors"
	"reflect"
)

// Remove removes elements from an array/slice, returns a new slice.
func Remove(series interface{}, removes ...interface{}) (interface{}, error) {
	t := reflect.TypeOf(series)
	if t.Kind() != reflect.Array && t.Kind() != reflect.Slice {
		return nil, errors.New("not an array or slice")
	}
	// m := reflect.MakeMapWithSize(t, len(removes))
	// m.SetMapIndex(key, val)
	return nil, errors.New("not implemented")
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
