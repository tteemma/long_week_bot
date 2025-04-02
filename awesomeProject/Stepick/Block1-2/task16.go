package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	if len(text) == 0 {
		fmt.Println("Wrong string")
		return
	}

	var result strings.Builder
	for i, char := range text {
		reguralExpr, _ := regexp.MatchString("^[0-9]+$", string(char))
		if reguralExpr {
			return
		}
		if i > 0 {
			result.WriteRune('*')
		}
		result.WriteString(string(char))
	}
	fmt.Println(result.String())

}
