package main

import (
	"bufio"
	"fmt"
	"os"
)

func repl() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome To The Pokedex!\nType help or exit to quit")

	for {
		scanner.Scan()
		text := scanner.Text()

		if text == "exit" {
			fmt.Println("Bye")
			break
		}
		if text == "help" {
			for _, command := range commands {
				fmt.Println(command.name, ":", command.description)
			}
		} else {
			fmt.Println("You said:", text)
		}

	}
}
