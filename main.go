package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		// Affiche le répertoire courant dans le prompt
		cwd, _ := os.Getwd()
		fmt.Printf("[%s] > ", cwd)

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Erreur de lecture :", err)
			continue
		}

		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		parts := strings.Fields(input)
		command := parts[0]
		args := parts[1:]

		switch command {
		case "cd":
			changeDirectory(args)
		case "exit":
			fmt.Println("Bye!")
			os.Exit(0)
		default:
			runExternalCommand(command, args)
		}
	}
}
