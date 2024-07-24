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
	mapPage := 0

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
			getMap(mapPage)
			mapPage++

		case "mapb":
			mapPage--
			getMap(mapPage)

		default:
			fmt.Println("HUH?")
		}

		fmt.Println("Type in a command or type help or exit to quit")
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
func getMap(mapPage int) {
	numlocations := 20
	fmt.Println("page", 1+(mapPage*numlocations))
	go func() {
		for i := 1; i < numlocations+1; i++ {
			location := getLocation(i + (mapPage * numlocations))
			fmt.Println(location.ID, ": ", location.Name)

		}
	}()

}
