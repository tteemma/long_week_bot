package main

import (
	"fmt"
	"strconv"
)

func main() {
	fn := func(x uint) uint {
		result := ""
		for _, digit := range strconv.Itoa(int(x)) {
			if digit == '0' || (digit-'0')%2 != 0 {
				continue
			}
			result += string(digit)
		}
		if result == "" {
			return 100
		}
		resNum, _ := strconv.Atoi(result)
		return uint(resNum)
	}

	var num uint
	fmt.Scan(&num)

	fmt.Println(fn(uint(num)))
}
