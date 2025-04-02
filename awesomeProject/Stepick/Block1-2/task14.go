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

	freq := make(map[rune]int)
	for _, char := range text {
		freq[char]++
	}

	var stringBuilder strings.Builder

	for _, char := range text {
		if freq[char] == 1 {
			stringBuilder.WriteRune(char)
		}
	}
	fmt.Println(stringBuilder.String())

}
