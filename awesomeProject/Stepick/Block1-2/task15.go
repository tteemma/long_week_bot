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

	if len(text) < 5 {
		fmt.Println("Wrong password")
		return
	}

	reguralExpr, _ := regexp.MatchString("^[a-zA-Z0-9]+$", text)
	if !reguralExpr {
		fmt.Println("Wrong password")
		return
	}
	fmt.Println("Ok")

}
