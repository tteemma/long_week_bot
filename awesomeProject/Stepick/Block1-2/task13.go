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

	var stringBuilder strings.Builder
	for i := 1; i < len(text); i += 2 {
		stringBuilder.WriteByte(text[i])
	}
	fmt.Println(stringBuilder.String())

}
