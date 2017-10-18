package fun

import (
	"fmt"
	"math"
	"reflect"
	"strings"
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
	if err != ErrNotArrayOrSlice {
		t.Errorf("returned %v, expected: %v", err, ErrNotArrayOrSlice)
	}
}

func TestFunctions(t *testing.T) {
	_, err := Remove([]func(){func() {}, func() {}}, func() {})
	if err != ErrNotSupported {
		t.Errorf("returned: %v, expected: %v", err, ErrNotSupported)
	}
}

func TestMapSquare(t *testing.T) {
	ints := []int{1, 3, 5, 7, 9}
	expected := []int{1, 9, 25, 49, 81}

	result, err := Map(ints, func(x int) int { return x * x })
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("returned: %v, expected: %v", result, expected)
	}
}

func TestMapSqrt(t *testing.T) {
	floats := []float64{0.25, 1.0, 4.0, 9.0, 16.0}
	expected := []float64{0.5, 1.0, 2.0, 3.0, 4.0}

	result, err := Map(floats, math.Sqrt)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("returned: %v, expected: %v", result, expected)
	}
}

func TestMapUppercase(t *testing.T) {
	strs := []string{"foo", "bar", "baz"}
	expected := []string{"FOO", "BAR", "BAZ"}

	result, err := Map(strs, strings.ToUpper)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(reflect.DeepEqual([]string{"foo", "bar", "baz"}, []string{"foo", "bar", "baz"}))
	// if fmt.Sprintf("%v", result) != fmt.Sprintf("%v", expected) {
	res := result.([]string)
	fmt.Println(res)
	if reflect.DeepEqual(result, expected) {
		t.Errorf("returned: %v, expected: %v", result, expected)
	}
}

func TestMapNotSeries(t *testing.T) {
	_, err := Map(2, func() {})
	if err != ErrNotArrayOrSlice {
		t.Errorf("returned: %v, expected: %v", err, ErrNotArrayOrSlice)
	}
}

func TestMapNotFunc(t *testing.T) {
	_, err := Map([]int{1, 2, 3}, 3)
	if err != ErrNotFunc {
		t.Errorf("returned: %v, expected: %v", err, ErrNotFunc)
	}
}

func TestMapErrFuncInParam(t *testing.T) {
	_, err := Map([]int{1, 2, 3}, func(i, j int) int { return 0 })
	if err != ErrFuncParam {
		t.Errorf("returned: %v, expected: %v", err, ErrFuncParam)
	}
}

func TestMapErrFuncOutParam(t *testing.T) {
	_, err := Map([]int{1, 2, 3}, func(i int) (int, int) { return 0, 1 })
	if err != ErrFuncParam {
		t.Errorf("returned: %v, expected: %v", err, ErrFuncParam)
	}
}

func TestMapNotCompatible(t *testing.T) {
	_, err := Map([]int{1, 2, 3}, func(s string) int { return len(s) })
	if err != ErrNotCompatible {
		t.Errorf("returned: %v, expected: %v", err, ErrNotCompatible)
	}
}

func TestCapitalize(t *testing.T) {
	s := "foo,bar,baz"
	expected := "Foo,Bar,Baz"

	result := Capitalize(s, ",")
	if result != expected {
		t.Errorf("returned: %s, expected: %s", result, expected)
	}
}

func TestStringsIn(t *testing.T)      {}
func TestFunctionsIn(t *testing.T)    {}
func TestIntsNotIn(t *testing.T)      {}
func TestStructsIn(t *testing.T)      {}
func TestPtrsIn(t *testing.T)         {}
func TestFilterEven(t *testing.T)     {}
func TestFilterPositive(t *testing.T) {}
