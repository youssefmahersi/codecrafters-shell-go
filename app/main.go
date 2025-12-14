package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

var builtin_keys = []string{"echo", "type", "exit"}

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
		content := strings.Join(args[1:], " ")

		switch args[0] {
		case "exit":
			return

		case "echo":
			if len(args) > 1 {
				fmt.Println(content)
			} else {
				fmt.Println()
			}
		case "type":
			if len(args) > 1 {
				if strings.TrimSpace(args[1]) != "" {

					if slices.Contains(builtin_keys, args[1]) == true {
						fmt.Printf("%s is a shell builtin\n", args[1])

					} else {
						fmt.Printf("%s : not found\n", args[1])
					}
				} else {
					fmt.Println()
				}
			} else {
				fmt.Println()
			}

		default:
			fmt.Printf("%s: command not found\n", args[0])
		}
	}
}
