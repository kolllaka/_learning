package main

import (
	"fmt"
)

func main() {
	arr := []int{10, 2, 3, 4, 1, 7, 23, 3, 0}

	fmt.Printf("arr: %v\n", arr)
	fmt.Printf("sort(arr): %v\n", sort(arr))
	fmt.Printf("arr: %v\n", arr)
}

func smallest(arr []int) (int, int) {
	small := arr[0]
	var indexSmall int

	for i, k := range arr {
		if small > k {
			small = k
			indexSmall = i
		}
	}

	return small, indexSmall
}

func sort(arr []int) []int {
	var newArr []int

	for range arr {
		small, indexSmall := smallest(arr)
		pop(indexSmall, &arr)
		newArr = append(newArr, small)
	}

	return newArr
}

func pop(i int, arr *[]int) {
	(*arr)[i] = (*arr)[len(*arr)-1]
	(*arr) = (*arr)[:len(*arr)-1]
}
