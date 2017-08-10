package handlers

import (
	"log"
)

var handlers = make(map[string]Command)

// Register is called to register a command handler for use by the program.
func Register(commandName string, commandHandler Command) {
	if _, exists := handlers[commandName]; exists {
		log.Fatalln(commandName, "Handler already registered")
	}

	log.Println("Registered", commandName, "command handler")
	handlers[commandName] = commandHandler
}

// Get is called to retrieve a command handler by it command
func Get(commandName string) Command {
	commandHandler, exists := handlers[commandName]

	if !exists {
		var empty emptyHandler
		return empty
	}

	return commandHandler
}
