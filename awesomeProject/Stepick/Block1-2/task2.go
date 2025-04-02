package main

import "fmt"

func main() {

	var num int
	fmt.Scan(&num)
	if num < 4 {
		fmt.Printf("%d is less than 4\n", num)
	}
	arr := make([]int, num)
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Println(arr[3])

}
