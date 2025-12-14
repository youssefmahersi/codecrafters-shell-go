package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading input:", err)
			return
		}

		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		args := strings.Fields(input)

		switch args[0] {
		case "exit":
			return

		case "echo":
			if len(args) > 1 {
				fmt.Println(strings.Join(args[1:], " "))
			} else {
				fmt.Println()
			}

		default:
			fmt.Printf("%s: command not found\n", args[0])
		}
	}
}
