package main

import "fmt"

func main() {
	var c int
	fmt.Scan(&c)

	arr := make([]int, c)
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Println(minNumInArr(arr))
}

func minNumInArr(arr []int) int {
	min_num := arr[0]
	cnt := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] < min_num {
			min_num = arr[i]
		}
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] == min_num {
			cnt++
		}
	}
	return cnt
}
