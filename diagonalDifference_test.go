package main

import (
	"reflect"
	"testing"
)

/*
 * Complete the 'diagonalDifference' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

func TestDiagonalDifference(t *testing.T) {

	t.Run("Base case", func(t *testing.T) {
		arr := [][]int32{{11, 2, 4}, {3, 5, 6}, {10, 8, -12}}

		got := diagonalDifference(arr)
		want := int32(15)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d given", got, want)
		}
	})

	t.Run("Array of zeros", func(t *testing.T) {
		arr := [][]int32{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}

		got := diagonalDifference(arr)
		want := int32(0)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d given", got, want)
		}
	})
}
