package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	var result strings.Builder
	for _, char := range text {
		reguralExpr, _ := regexp.MatchString("^[0-9]+$", string(char))
		if !reguralExpr {
			return
		}
		digit, _ := strconv.Atoi(string(char))

		res := digit * digit

		result.WriteString(strconv.Itoa(res))
	}
	fmt.Println(result.String())
}
