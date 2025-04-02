package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var num int
	reader := bufio.NewReader(os.Stdin)

	fmt.Scan(&num)
	arr := make([]int, num)

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	fields := strings.Fields(input)

	cnt := 0
	for i, f := range fields {
		arr[i], _ = strconv.Atoi(f)
		if arr[i] > 0 {
			cnt++
		}
	}
	fmt.Println(arr)
	fmt.Println(cnt)
}
