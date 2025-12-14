package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

var builtinKeys = []string{"echo", "type", "exit"}

func checkCommandPermission(cmd string) (*string, bool) {
	pathEnv := os.Getenv("PATH")
	dirs := strings.Split(pathEnv, string(os.PathListSeparator))

	for _, dir := range dirs {
		if dir == "" {
			continue
		}

		fullPath := filepath.Join(dir, cmd)

		info, err := os.Stat(fullPath)
		if err != nil {
			continue
		}

		if !info.IsDir() && info.Mode()&0111 != 0 {
			fp := fullPath
			return &fp, true
		}
	}

	return nil, false
}

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
		content := ""
		if len(args) > 1 {
			content = strings.Join(args[1:], " ")
		}

		switch args[0] {
		case "exit":
			return

		case "echo":
			fmt.Println(content)

		case "type":
			if len(args) < 2 || strings.TrimSpace(args[1]) == "" {
				fmt.Println()
				continue
			}

			target := args[1]

			if slices.Contains(builtinKeys, target) {
				fmt.Printf("%s is a shell builtin\n", target)
				continue
			}

			foundFile, exists := checkCommandPermission(target)
			if exists {
				fmt.Printf("%s is %s\n", target, *foundFile)
			} else {
				fmt.Printf("%s: not found\n", target)
			}

		default:
			fmt.Printf("%s: command not found\n", args[0])
		}
	}
}
