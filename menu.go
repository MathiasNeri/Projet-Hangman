package main

import (
	"fmt"
)

func Menu() {

	fmt.Println("║═══════════════════════════════════════════════════║")
	fmt.Println("║                                                   ║")
	fmt.Println("║           BIENVENUE DANS HORROR!PENDU             ║")
	fmt.Println("║                                                   ║")
	fmt.Println("║═══════════════════════════════════════════════════║")
	fmt.Println("║                                                   ║")
	fmt.Println("║             1.  Jouer                             ║")
	fmt.Println("║             2.  Quitter                           ║")
	fmt.Println("║                                                   ║")
	fmt.Println("╚═══════════════════════════════════════════════════╝")

	var choice int

	fmt.Scanln(&choice)

	switch choice {

	case 1:
		ClearConsole()
		fmt.Println("QUE LA PARTIE COMMENCE !")
		
	case 2:
		Menu()
		break

	default:
		fmt.Println("Choix invalide. Veuillez sélectionner une autre option.")
		Menu()
		break
	}
}
