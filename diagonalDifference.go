package main

import (
	"fmt"
	"math"
)

/*
 * Complete the 'diagonalDifference' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

func diagonalDifference(arr [][]int32) int32 {
	var diagonalOne int32 = 0
	var diagonalTwo int32 = 0

	for index, value := range arr {
		diagonalOne += value[index]
		diagonalTwo += value[len(value)-index-1]
	}

	return int32(math.Abs(float64(diagonalOne - diagonalTwo)))
}

func main() {
	arr := [][]int32{{11, 2, 4}, {3, 5, 6}, {10, 8, -12}}

	result := diagonalDifference(arr)

	fmt.Println(result)
}
