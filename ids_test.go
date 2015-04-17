package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestIDsSort(t *testing.T) {
	s := []int64{5, 3, 1, 4, 2}
	ss := IDs(s)
	sort.Sort(ss)
	s = []int64(ss)
	expected := []int64{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected: %q, but got: %q", expected, s)
	}
}

func TestIDsDiff(t *testing.T) {
	assert := func(a, b, expected []int64) {
		l := IDs(a)
		r := IDs(b)
		l.Diff(r)
		e := IDs(expected)
		if e.Eq(r) {
			t.Fatalf("expected: %q, but got %q", e, l)
		}
	}

	assert(
		[]int64{1, 2, 3, 5, 10, 11},
		[]int64{2, 8, 11},
		[]int64{1, 3, 5, 10},
	)
	assert(
		[]int64{1, 2, 3, 4, 5},
		[]int64{1, 2, 3, 4, 5},
		[]int64{},
	)
	assert(
		[]int64{},
		[]int64{1, 2, 3, 4, 5},
		[]int64{},
	)
	assert(
		[]int64{1, 2, 3, 4, 5},
		[]int64{},
		[]int64{1, 2, 3, 4, 5},
	)
}
