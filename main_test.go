package main

import (
	"reflect"
	"testing"
)

func TestDiffInt64(t *testing.T) {
	assert := func(a, b, expected []int64) {
		d := DiffInt64(a, b)
		if !reflect.DeepEqual(expected, d) {
			t.Errorf("expected: %+q, but got %+q", expected, d)
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
