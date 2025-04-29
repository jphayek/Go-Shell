package main

import (
	"fmt"
	"os"
	"os/exec"
)

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

func runExternalCommand(name string, args []string) {
	cmd := exec.Command(name, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("Commande échouée :", err)
	}
}
