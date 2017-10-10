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
		t.Errorf("returned: %v, expected: %v")
	}

	result, err = Remove(originalSli, 0, 1, 3, 6)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(result, removed) {
		t.Errorf("returned: %v, expected: %v")
	}
}

func TestStringsRemove(t *testing.T)  {}
func TestFloatsRemove(t *testing.T)   {}
func TestMapsRemove(t *testing.T)     {}
func TestPointersRemove(t *testing.T) {}
func TestChansRemove(t *testing.T)    {}
func TestStructsRemove(t *testing.T)  {}
