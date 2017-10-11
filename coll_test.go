package fun

import (
	"reflect"
	"testing"
)

func TestIntsRemove(t *testing.T) {
	originalArr := [7]int{0, 1, 2, 3, 4, 5, 6}
	originalSli := []int{0, 1, 2, 3, 4, 5, 6}
	removed := []int{2, 4, 5}

	result, err := Remove(originalArr, 0, 1, 3, 6)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(result, removed) {
		t.Errorf("returned: %v, expected: %v", result, removed)
	}

	result, err = Remove(originalSli, 0, 1, 3, 6)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(result, removed) {
		t.Errorf("returned: %v, expected: %v", result, removed)
	}
}

func TestStringsRemove(t *testing.T) {
	originalArr := [3]string{"foo", "bar", "baz"}
	originalSli := []string{"foo", "bar", "baz"}
	removed := []string{}

	result, err := Remove(originalArr, "foo", "bar", "baz")
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(result, removed) {
		t.Errorf("returned: %v, expected: %v", result, removed)
	}

	result, err = Remove(originalSli, "foo", "bar", "baz")
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(result, removed) {
		t.Errorf("returned: %v, expected: %v", result, removed)
	}
}

func TestFloatsRemove(t *testing.T) {
	originalArr := [5]float64{3.1415926, 2.71828, 0.57721566490153286060651209, 2.6854520010653064453, 1.6180339887498948482}
	originalSli := []float64{3.1415926, 2.71828, 0.57721566490153286060651209, 2.6854520010653064453, 1.6180339887498948482}
	removed := []float64{0.57721566490153286060651209, 2.6854520010653064453, 1.6180339887498948482}

	result, err := Remove(originalArr, 3.1415926, 2.71828)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(result, removed) {
		t.Errorf("returned: %v, expected: %v", result, removed)
	}

	result, err = Remove(originalSli, 3.1415926, 2.71828)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(result, removed) {
		t.Errorf("returned: %v, expected: %v", result, removed)
	}
}

func TestMapsRemove(t *testing.T) {
	maps := []map[string]int{
		map[string]int{
			"foo": 1,
		},
		map[string]int{
			"bar": 2,
		},
		map[string]int{
			"baz": 3,
		},
	}
	removed := []map[string]int{
		map[string]int{
			"baz": 3,
		},
	}

	result, err := Remove(maps, map[string]int{"foo": 1}, map[string]int{"bar": 2})
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(result, removed) {
		t.Errorf("returned: %v, expected: %v", result, removed)
	}
}

func TestIncompatible(t *testing.T) {
	strs := []string{"foo", "bar", "baz"}
	_, err := Remove(strs, 1, 2)
	if err == nil {
		t.Errorf("should return %v", ErrNotCompatible)
	}
}

func TestNoRemove(t *testing.T) {
	ints := []int{1, 2, 3}
	result, err := Remove(ints)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(ints, result) {
		t.Errorf("returned: %v, expected: %v", result, ints)
	}
}

type testStruct struct {
	foo string
	bar int
	baz float64
}

func TestStructsRemove(t *testing.T) {
	structs := []testStruct{
		testStruct{"foo", 1, 1.2},
		testStruct{"bar", 2, 2.2},
		testStruct{"baz", 3, 3.2},
	}
	removed := []testStruct{
		testStruct{"bar", 2, 2.2},
	}

	result, err := Remove(structs, testStruct{"foo", 1, 1.2}, testStruct{"baz", 3, 3.2})
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(result, removed) {
		t.Errorf("returned: %v, expected: %v", result, removed)
	}
}

func TestPointersRemove(t *testing.T) {
	foo := &testStruct{"foo", 1, 1.2}
	bar := &testStruct{"bar", 2, 2.2}
	baz := &testStruct{"baz", 3, 3.2}

	ptrs := []*testStruct{
		foo,
		bar,
		baz,
	}
	removed := []*testStruct{
		bar,
	}

	result, err := Remove(ptrs, foo, baz)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(result, removed) {
		t.Errorf("returned: %v, expected: %v", result, removed)
	}
}

func TestNotArraySlice(t *testing.T) {
	_, err := Remove(map[string]int{}, 2, 3)
	if err != ErrNotArraySlice {
		t.Errorf("returned %v, expected: %v", err, ErrNotArraySlice)
	}
}

func TestFunctions(t *testing.T) {
	_, err := Remove([]func(){func() {}, func() {}}, func() {})
	if err != ErrNotSupported {
		t.Errorf("returned: %v, expected: %v", err, ErrNotSupported)
	}
}
