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

func TestMapsRemove(t *testing.T)     {}
func TestChansRemove(t *testing.T)    {}
func TestStructsRemove(t *testing.T)  {}
func TestPointersRemove(t *testing.T) {}
