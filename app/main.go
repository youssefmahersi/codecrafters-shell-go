package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var stop = false
	for stop != true {

		fmt.Print("$ ")
		command, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		args := strings.Split(command, " ")
		if command[:len(command)-1] == "exit" {
			stop = true
			break
		}
		switch args[0] {
		case "echo":
			fmt.Printf("%s", args[1])
		default:
			fmt.Println(command[:len(command)-1] + ": command not found")
		}

	}
}
