package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.TrimSpace(text)

	if len(text) == 0 {
		fmt.Println("Wrong")
		return
	}

	if unicode.IsUpper(rune(text[0])) && text[len(text)-1] == '.' {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}
