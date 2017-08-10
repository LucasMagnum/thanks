package handlers

import (
	"errors"
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
func Get(commandName string) (Command, error) {
	commandHandler, exists := handlers[commandName]

	if !exists {
		log.Fatalln(commandName, "Handler not found")
		return commandHandler, errors.New("Handler not found")
	}

	return commandHandler, nil
}
