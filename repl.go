package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func scanNext(scanner bufio.Scanner) []string {
	scanner.Scan()
	return cleanInput(scanner.Text())
}

func repl() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome To The Pokedex!")

	for {
		fmt.Println("Type in a command or type help or exit to quit")
		input := scanNext(*scanner)
		text := input[0]

		switch text {
		case "exit":
			fmt.Println("Bye")
			return
		case "help":
			for _, command := range commands {
				fmt.Println(command.name, ":", command.description)
			}
		case "map":
			getMap()
		loop:

			for {
				fmt.Println("Enter : NEXT| BACK| EXIT")
				input := scanNext(*scanner)
				text := input[0]
				switch text {
				case "next":
					nextPage()
					getMap()
				case "back":
					previousPage()
					getMap()
				case "exit":
					fmt.Println("Exiting map")
					break loop
				default:
					fmt.Println("HUH?")
				}
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
