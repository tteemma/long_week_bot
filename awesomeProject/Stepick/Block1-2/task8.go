package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)
	var toRemove int
	fmt.Scan(&toRemove)
	fmt.Println(detectToDelete(n, toRemove))

}

func detectToDelete(num, digitToRemove int) int {
	strNum := strconv.Itoa(num)
	strDigNum := strconv.Itoa(digitToRemove)
	result := ""

	for _, i := range strNum {
		if string(i) != strDigNum {
			result += string(i)
		}
	}

	if result == "" {
		return 0
	}

	resultNum, _ := strconv.Atoi(result)

	return resultNum
}
