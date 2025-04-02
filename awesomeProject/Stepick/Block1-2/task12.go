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

	text1, _ := reader.ReadString('\n')
	text1 = strings.TrimSpace(text1)

	index := strings.Index(text, text1)

	fmt.Println(index)
}
