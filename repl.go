package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func repl() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome To The Pokedex!\nType help or exit to quit")

	for {
		scanner.Scan()
		input := cleanInput(scanner.Text())
		text := input[0]

		switch text {
		case "exit":
			fmt.Println("Bye")
			return
		case "help":
			for _, command := range commands {
				fmt.Println(command.name, ":", command.description)
			}
		default:
			fmt.Println("HUH?")
		}

	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
