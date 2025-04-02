package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	result_num := sum_digits(num)
	for result_num > 9 {
		result_num = sum_digits(result_num)
	}
	fmt.Println(result_num)

}
func sum_digits(n int) int {
	sum := 0
	if n < 0 {
		n = -n
	}
	for n > 0 {
		sum += n % 10
		n = n / 10
	}
	return sum
}
