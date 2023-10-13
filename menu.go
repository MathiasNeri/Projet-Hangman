package projethangman

import (
	"fmt"
	"time"
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
		time.Sleep(2000 * time.Millisecond)
		ClearConsole()
		fmt.Println("║═══════════════════════════════════════════════════║")
		fmt.Println("║             1.  Easy                              ║")
		fmt.Println("║             2.  Medium                            ║")
		fmt.Println("║             3.  Hard                              ║")
		fmt.Println("╚═══════════════════════════════════════════════════╝")

		var ChoiceDifficulty int

		fmt.Scanln(&ChoiceDifficulty)

		switch ChoiceDifficulty {

		case 1:

			ClearConsole()
			fmt.Println("Vous avez choisi la difficultée easy, Bonne chance !")
			time.Sleep(2000 * time.Millisecond)
			ClearConsole()
			nomFichier := "easy.txt"
			RandomWords(nomFichier)
			mot, err := RandomWords(nomFichier)
			if err != nil {
				fmt.Println("Erreur lors ded la lecture du fichier :", err)
				return
			}
			fmt.Printf("Le mot est : %s\n", mot)

		case 2:
			ClearConsole()
			fmt.Println("Vous avez choisi la difficultée medium, Bonne chance !")
			time.Sleep(2000 * time.Millisecond)
			ClearConsole()
			nomFichier := "medium.txt"
			RandomWords(nomFichier)
			mot, err := RandomWords(nomFichier)
			if err != nil {
				fmt.Println("Erreur lors ded la lecture du fichier :", err)
				return
			}
			fmt.Printf("Le mot est : %s\n", mot)

		case 3:
			ClearConsole()
			fmt.Println("Vous avez choisi la difficultée hard, Bonne chance !")
			time.Sleep(2000 * time.Millisecond)
			ClearConsole()
			nomFichier := "hard.txt"
			RandomWords(nomFichier)
			mot, err := RandomWords(nomFichier)
			if err != nil {
				fmt.Println("Erreur lors ded la lecture du fichier :", err)
				return
			}
			fmt.Printf("Le mot est : %s\n", mot)

		default:
			fmt.Println("Choix invalide. Veuillez sélectionner une autre option.")
			return
		}
	case 2:
		break

	default:
		fmt.Println("Choix invalide. Veuillez sélectionner une autre option.")
		Menu()
		break
	}
}
