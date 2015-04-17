package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestInt64sSort(t *testing.T) {
	s := []int64{5, 3, 1, 4, 2}
	ss := int64s(s)
	sort.Sort(ss)
	s = []int64(ss)
	expected := []int64{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected: %q, but got: %q", expected, s)
	}
}
