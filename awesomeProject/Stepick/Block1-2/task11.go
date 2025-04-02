package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	runes := []rune(text)

	l := len(runes)
	p := l - 1
	ind := 0

	if l == 0 {
		fmt.Println("Нет")
		return
	}

	for ind < p {
		if runes[ind] != runes[p] {
			fmt.Println("Нет")
			return
		}
		ind++
		p--
	}
	fmt.Println("Палиндром")

}
