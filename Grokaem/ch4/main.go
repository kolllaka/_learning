package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 7, 5, 200, 123, 3}
	fmt.Printf("sum(arr): %v\n", sum(arr))
	fmt.Printf("count(arr): %v\n", count(arr))
	fmt.Printf("max(arr): %v\n", max(arr))
	fmt.Printf("sort(arr): %v\n", sort(arr))
}

func sum(arr []int) int {
	if len(arr) > 0 {
		return arr[0] + sum(arr[1:])
	}

	return 0
}

func count(arr []int) int {
	if len(arr) > 0 {
		return 1 + count(arr[1:])
	}

	return 0
}

func max(arr []int) int {
	maximum := arr[0]

	if len(arr) > 1 {
		if maximum < max(arr[1:]) {
			return max(arr[1:])
		}
	}
	// for _, v := range arr {
	// 	if max < v {
	// 		max = v
	// 	}
	// }

	return maximum
}

func sort(arr []int) []int {
	var less []int
	var greater []int
	var sortArray []int
	var oporaIndex = len(arr) / 2

	if len(arr) < 2 {
		return arr
	}

	opora := arr[oporaIndex]
	for i, k := range arr {
		if i == oporaIndex {
			continue
		}

		if k <= opora {
			less = append(less, k)
		} else {
			greater = append(greater, k)
		}
	}

	sortArray = append(sortArray, sort(less)...)
	sortArray = append(sortArray, opora)
	sortArray = append(sortArray, sort(greater)...)

	return sortArray
}
