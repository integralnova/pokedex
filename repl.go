package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
		case "map":

			if len(input) > 1 {
				num, notNum := strconv.Atoi(input[1])
				if notNum != nil || num == 1 {
					fmt.Println("Interesting... heres 10...")
					fmt.Println(mapgetter(10)[1])
				} else {
					fmt.Println(mapgetter(num)[1])
				}
			} else {
				fmt.Println(mapgetter(10)[1])
			}

		case "mapb":
			fmt.Println(mapgetter(10))
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
