package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	fmt.Println(sumDigit(num))
}

func sumDigit(n int) int {
	sum := 0
	if n < 0 {
		n = -n
	}
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}
