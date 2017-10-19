package fun

import (
	"errors"
	"reflect"
	"unsafe"
)

// errors defined
var (
	ErrNotImplemented  = errors.New("not implemented")
	ErrNotSupported    = errors.New("not supported")
	ErrNotArrayOrSlice = errors.New("not an array or slice")
	ErrNotCompatible   = errors.New("types not compatible")
	ErrNotFunc         = errors.New("not a function")
	ErrFuncParam       = errors.New("invalid function parameters")
)

// Remove removes elements from an array/slice and returns a new slice in comlexity
// of O(len(series)). Notice that it falls back to O(len(series) * len(removes))
// when it comes to slice of maps/slice.
func Remove(series interface{}, removes ...interface{}) (interface{}, error) {
	st := reflect.TypeOf(series)
	sv := reflect.ValueOf(series)

	switch {
	case st.Kind() != reflect.Array && st.Kind() != reflect.Slice:
		return nil, ErrNotArrayOrSlice
	case st.Elem().Kind() == reflect.Func:
		return nil, ErrNotSupported
	case len(removes) == 0:
		return series, nil
	case st.Elem().Kind() != reflect.TypeOf(removes[0]).Kind():
		return nil, ErrNotCompatible
	}

	removed := reflect.MakeSlice(reflect.SliceOf(st.Elem()), 0, 0)
	switch st.Elem().Kind() {
	case reflect.Map, reflect.Slice:
		for i := 0; i < sv.Len(); i++ {
			found := false
			for _, r := range removes {
				if reflect.DeepEqual(sv.Index(i).Interface(), r) {
					found = true
					break
				}
			}
			if !found {
				removed = reflect.Append(removed, sv.Index(i))
			}
		}
		return removed.Interface(), nil
	default:
		filter := reflect.MakeMapWithSize(reflect.MapOf(st.Elem(), reflect.TypeOf(true)), len(removes))
		for _, r := range removes {
			filter.SetMapIndex(reflect.ValueOf(r), reflect.ValueOf(true))
		}
		for i := 0; i < sv.Len(); i++ {
			if !filter.MapIndex(sv.Index(i)).IsValid() {
				removed = reflect.Append(removed, sv.Index(i))
			}
		}
		return removed.Interface(), nil
	}
}

// Map accepts an array/slice and a function, call function on every element
// and returns a new slice of results.
func Map(series interface{}, f interface{}) (interface{}, error) {
	st := reflect.TypeOf(series)
	sv := reflect.ValueOf(series)
	ft := reflect.TypeOf(f)
	fv := reflect.ValueOf(f)

	switch {
	case st.Kind() != reflect.Slice && st.Kind() != reflect.Array:
		return nil, ErrNotArrayOrSlice
	case ft.Kind() != reflect.Func:
		return nil, ErrNotFunc
	case ft.NumIn() != 1 || ft.NumOut() != 1:
		return nil, ErrFuncParam
	case ft.In(0).Kind() != st.Elem().Kind():
		return nil, ErrNotCompatible
	}

	mapped := reflect.MakeSlice(reflect.SliceOf(ft.Out(0)), 0, 0)
	for i := 0; i < sv.Len(); i++ {
		mapped = reflect.Append(mapped, fv.Call([]reflect.Value{sv.Index(i)})...)
	}

	return mapped.Interface(), nil
}

// In checks whether an element is in collection.
func In(collection interface{}, element interface{}) bool {
	return false
}

// Filter returns a new slice with all elements that pass the test implemented by
// the provided function filterFunc.
func Filter(series interface{}, filterFunc interface{}) (interface{}, error) {
	return nil, ErrNotImplemented
}

// Find returns the value of the first element in series taht satisfies the
// provided testing function.
func Find(series interface{}, findFunc interface{}) (interface{}, error) {
	return nil, ErrNotImplemented
}

// Every returns true if all elements in the array pass the test implemented by
// the provided function.
func Every(series interface{}, everyFunc interface{}) bool {
	return false
}

// Some returns true if at least one element in the series passes the test
// implemented by someFunc.
func Some(series interface{}, someFunc interface{}) bool {
	return false
}

// Reduce applies a reduceFunc against an accumulator and each element in the series
// frome left to right to reduce it to a single value, reduceFunc must be func(acc, val) val.
func Reduce(series interface{}, reduceFunc interface{}, initial interface{}) (interface{}, error) {
	return nil, ErrNotImplemented
}

// push pop splice
type visit struct {
	a1  unsafe.Pointer
	a2  unsafe.Pointer
	typ reflect.Type
}

func DeepEqual(x, y interface{}) bool {
	if x == nil || y == nil {
		return x == y
	}
	v1 := reflect.ValueOf(x)
	v2 := reflect.ValueOf(y)
	if v1.Type() != v2.Type() {
		return false
	}
	return DeepValueEqual(v1, v2, make(map[visit]bool), 0)
}

func DeepValueEqual(v1, v2 reflect.Value, visited map[visit]bool, depth int) bool {
	if !v1.IsValid() || !v2.IsValid() {
		return v1.IsValid() == v2.IsValid()
	}
	if v1.Type() != v2.Type() {
		return false
	}

	// if depth > 10 { panic("deepValueEqual") }	// for debugging

	// We want to avoid putting more in the visited map than we need to.
	// For any possible reference cycle that might be encountered,
	// hard(t) needs to return true for at least one of the types in the cycle.
	hard := func(k reflect.Kind) bool {
		switch k {
		case reflect.Map, reflect.Slice, reflect.Ptr, reflect.Interface:
			return true
		}
		return false
	}

	if v1.CanAddr() && v2.CanAddr() && hard(v1.Kind()) {
		addr1 := unsafe.Pointer(v1.UnsafeAddr())
		addr2 := unsafe.Pointer(v2.UnsafeAddr())
		if uintptr(addr1) > uintptr(addr2) {
			// Canonicalize order to reduce number of entries in visited.
			// Assumes non-moving garbage collector.
			addr1, addr2 = addr2, addr1
		}

		// Short circuit if references are already seen.
		typ := v1.Type()
		v := visit{addr1, addr2, typ}
		if visited[v] {
			return true
		}

		// Remember for later.
		visited[v] = true
	}

	switch v1.Kind() {
	case reflect.Slice:
		if v1.IsNil() != v2.IsNil() {
			return false
		}
		if v1.Len() != v2.Len() {
			return false
		}
		if v1.Pointer() == v2.Pointer() {
			return true
		}
		for i := 0; i < v1.Len(); i++ {
			if !DeepValueEqual(v1.Index(i), v2.Index(i), visited, depth+1) {
				return false
			}
		}
		return true
	default:
		// Normal equality suffices
		return true
	}
}
