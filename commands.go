package main

type command struct {
	name        string
	description string
	//callback    func() error
}

var commands = map[string]command{
	"help": {
		name:        "help",
		description: "Displays a help message",
	},
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
	},
	"map": {
		name:        "map",
		description: "Displays the names of 20 location areas in the Pokemon world",
	},
}
