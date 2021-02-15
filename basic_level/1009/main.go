package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var input string

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
	}

	output := strings.Split(input, " ")
	for i := len(output) - 1; i >= 0; i-- {
		fmt.Print(output[i])
		if i != 0 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
