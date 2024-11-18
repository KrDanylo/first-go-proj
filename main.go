package main

import (
	"fmt"
	"log"

	"MyFirstApp/greetings"
)

func main() {
	// Configure the logger with a custom prefix and flags.
	setupLogger()

	// List of names to greet.
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Generate greetings for the list of names.
	messages, err := greetings.Hellos(names)
	if err != nil {
		logErrorAndExit(err)
	}

	// Print the greetings to the console.
	printGreetings(messages)
}

// setupLogger configures the logger with a custom prefix and disables additional flags.
func setupLogger() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
}

// logErrorAndExit logs an error and terminates the program.
func logErrorAndExit(err error) {
	log.Fatal(err)
}

// printGreetings prints a map of greetings to the console in a formatted way.
func printGreetings(messages map[string]string) {
	fmt.Println("Greeting messages:")
	for name, message := range messages {
		fmt.Printf("- %s: %s\n", name, message)
	}
}
