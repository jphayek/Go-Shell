package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		// Affiche le répertoire courant
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

		// Découpe la commande
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

// Commande interne : cd
func changeDirectory(args []string) {
	if len(args) < 1 {
		fmt.Println("cd: chemin manquant")
		return
	}
	err := os.Chdir(args[0])
	if err != nil {
		fmt.Println("cd: erreur :", err)
	}
}

// Commandes externes : ls, rm, echo, etc.
func runExternalCommand(name string, args []string) {
	var inputFile *os.File
	var outputFile *os.File
	var err error

	// Analyse des redirections dans les args
	cleanArgs := []string{}
	for i := 0; i < len(args); i++ {
		if args[i] == ">" && i+1 < len(args) {
			outputFile, err = os.Create(args[i+1])
			if err != nil {
				fmt.Println("Erreur création fichier de sortie :", err)
				return
			}
			i++ // skip filename
		} else if args[i] == "<" && i+1 < len(args) {
			inputFile, err = os.Open(args[i+1])
			if err != nil {
				fmt.Println("Erreur ouverture fichier d'entrée :", err)
				return
			}
			i++ // skip filename
		} else {
			cleanArgs = append(cleanArgs, args[i])
		}
	}

	cmd := exec.Command(name, cleanArgs...)

	// Redirection d'entrée/sortie
	if inputFile != nil {
		cmd.Stdin = inputFile
	} else {
		cmd.Stdin = os.Stdin
	}

	if outputFile != nil {
		cmd.Stdout = outputFile
	} else {
		cmd.Stdout = os.Stdout
	}

	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Println("Erreur d'exécution :", err)
	}

	// Fermeture des fichiers
	if inputFile != nil {
		inputFile.Close()
	}
	if outputFile != nil {
		outputFile.Close()
	}
}
