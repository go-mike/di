package di

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapSlice(t *testing.T) {
	source := rangeSlice(1, 5)
	actual := mapSlice(source, func(item int) string {
		return strconv.FormatInt(int64(item), 10)
	})
	expected := []string{"1", "2", "3", "4", "5"}
	assert.Equal(t, expected, actual)
}

func TestFlattenSlice(t *testing.T) {
	source := [][]int{rangeSlice(1, 3), rangeSlice(4, 3)}
	actual := flattenSlice(source)
	expected := rangeSlice(1, 6)
	assert.Equal(t, expected, actual)
}

func TestBindSlice(t *testing.T) {
	source := rangeSlice(1, 5)
	actual := bindSlice(source, func(item int) []int {
		return rangeSlice(1, item)
	})
	expected := []int{1, 1, 2, 1, 2, 3, 1, 2, 3, 4, 1, 2, 3, 4, 5}
	assert.Equal(t, expected, actual)
}

func TestFilterSlice(t *testing.T) {
	source := rangeSlice(1, 5)
	actual := filterSlice(source, func(item int) bool {
		return item % 2 == 0
	})
	expected := []int{2, 4}
	assert.Equal(t, expected, actual)
}

func TestFindSlice(t *testing.T) {
	source := rangeSlice(1, 5)
	actual := findSlice(source, func(item int) bool {
		return item == 3
	})
	expected := 3
	assert.Equal(t, expected, *actual)
}

func TestFindSlice_NotFound(t *testing.T) {
	source := rangeSlice(1, 5)
	actual := findSlice(source, func(item int) bool {
		return item == 6
	})
	assert.Equal(t, (*int)(nil), actual)
}

func TestFilterSliceIndices(t *testing.T) {
	source := rangeSlice(1, 5)
	actual := filterSliceIndices(source, func(item int) bool {
		return item % 2 == 0
	})
	expected := []int{1, 3}
	assert.Equal(t, expected, actual)
}

func TestFilterSliceAll(t *testing.T) {
	source := rangeSlice(1, 5)
	actual, indices := filterSliceAll(source, func(item int) bool {
		return item % 2 == 0
	})
	expected := []int{2, 4}
	assert.Equal(t, expected, actual)
	assert.Equal(t, []int{1, 3}, indices)
}

func TestPartitionSlice(t *testing.T) {
	source := rangeSlice(1, 5)
	trueItems, falseItems := partitionSlice(source, func(item int) bool {
		return item % 2 == 0
	})
	expectedTrueItems := []int{2, 4}
	assert.Equal(t, expectedTrueItems, trueItems)
	expectedFalseItems := []int{1, 3, 5}
	assert.Equal(t, expectedFalseItems, falseItems)
}

func TestRangeSlice(t *testing.T) {
	actual := rangeSlice(1, 5)
	expected := []int{1, 2, 3, 4, 5}
	assert.Equal(t, expected, actual)
}

func TestRangeSliceWithStep(t *testing.T) {
	actual := rangeSliceWithStep(1, 5, 2)
	expected := []int{1, 3, 5, 7, 9}
	assert.Equal(t, expected, actual)
}

func TestRangeMapSlice(t *testing.T) {
	actual := rangeMapSlice(1, 5, func(item int) int {
		return item * 2
	})
	expected := []int{2, 4, 6, 8, 10}
	assert.Equal(t, expected, actual)
}
