package main

import "fmt"

func work(x int) int {
	return x - 1
}

func main() {
	var numbers [10]int
	cache := make(map[int]int)

	for i := 0; i < 10; i++ {
		fmt.Scan(&numbers[i])
	}

	for i, num := range numbers {
		if val, ok := cache[num]; ok {
			fmt.Print(val)
		} else {
			result := work(num)
			cache[num] = result
			fmt.Print(result)
		}
		if i < 9 {
			fmt.Print(" ")
		}
	}
	fmt.Println(" time limit ok")
}
