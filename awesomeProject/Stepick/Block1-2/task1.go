package main

import "fmt"

func main() {
	workArray := make([]int, 10)
	for i := 0; i < len(workArray); i++ {
		fmt.Scan(&workArray[i])
	}
	fmt.Println(workArray)
	swapCount := 3
	for i := 0; i < swapCount; i++ {
		var ind1, ind2 int
		fmt.Scan(&ind1, &ind2)

		if ind1 < 0 || ind2 >= len(workArray) || ind2 < 0 || ind2 >= len(workArray) {
			fmt.Printf("%d is out of range", ind1)
			continue
		}

		workArray[ind1], workArray[ind2] = workArray[ind2], workArray[ind1]
	}
	fmt.Println(workArray)
}
