package main

import (
	"fmt"
	"math"
)

func fibonachiSearch(array []int, elem int) int {

	fibMMm2 := 0              // (m-2)'th Fibonacci No.
	fibMMm1 := 1              // (m-1)'th Fibonacci No.
	fibM := fibMMm2 + fibMMm1 // m'th Fibonacci

	for fibM < len(array) {
		fibMMm2 = fibMMm1
		fibMMm1 = fibM
		fibM = fibMMm2 + fibMMm1
	}

	var offset = -1

	fmt.Println(fibM, fibMMm1, fibMMm2)

	for fibM > 1 {
		i := int(math.Min(float64(offset+fibMMm2), float64(len(array)-1)))
		if array[i] < elem {
			fibM = fibMMm1
			fibMMm1 = fibMMm2
			fibMMm2 = fibM - fibMMm1
			offset = i
		} else if array[i] > elem {
			fibM = fibMMm2
			fibMMm1 = fibMMm1 - fibMMm2
			fibMMm2 = fibM - fibMMm1
		} else {
			return i
		}

		if array[len(array)-1] == elem {
			return len(array) - 1
		}
	}
	return -1
}

func main() {
	array := []int{10, 22, 35, 40, 45, 50, 80, 82, 85, 90, 100, 235}
	elem := 82

	index := fibonachiSearch(array, elem)
	fmt.Println(index)
}
