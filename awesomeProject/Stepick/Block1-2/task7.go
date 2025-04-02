package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	m := 1
	for m <= n {
		fmt.Print(m, " ")
		m = m * 2
	}
}
