package fonction

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func JouerAuPendu(nomFichier, motSecret string) {
	essaisRestants := 10
	lettresRestantes := len(motSecret)
	lettresDejaEssayees := make(map[rune]bool)

	motCache := RandomLetters(motSecret)

	fmt.Println("Bienvenue au jeu du pendu!")

	var images []string
	path := "./data/image_pendu.txt"

	// Charger les images du pendu depuis le fichier
	fichierImages, err := os.Open(path)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier d'images:", err)
		return
	}
	defer fichierImages.Close()
	if err != nil {
		fmt.Println("Erreur lors de la fermeture du fichier d'images:", err)
	}

	scannerImages := bufio.NewScanner(fichierImages)
	var image string
	for scannerImages.Scan() {
		ligne := scannerImages.Text()
		if ligne == "=========" {
			if image != "" {
				images = append([]string{image}, images...) // Ajoutez l'image en haut de la pile
			}
			image = ""
		} else {
			image += ligne + "\n"
		}
	}
	if image != "" {
		images = append([]string{image}, images...) // Ajoutez la dernière image en haut de la pile
	}

	for essaisRestants > 0 && lettresRestantes > 0 {
		fmt.Printf("Mot caché: %s\n", motCache)
		fmt.Printf("Tentatives restantes: %d\n", essaisRestants)

		if essaisRestants >= 0 && essaisRestants < len(images) {
			fmt.Println(images[essaisRestants])
		} else {
			fmt.Println("Nombre d'essais non pris en charge.")
		}

		var lettre rune
		fmt.Print("Devinez une lettre: ")
		fmt.Scanf("%c\n", &lettre) // Ajoutez "\n" pour consommer le caractère de retour à la ligne
		lettre = unicode.ToLower(lettre)

		if !unicode.IsLetter(lettre) {
			fmt.Println("Veuillez entrer une lettre alphabétique.")
			continue
		}

		if lettresDejaEssayees[lettre] {
			fmt.Println("Vous avez déjà essayé cette lettre.")
			continue
		}

		lettresDejaEssayees[lettre] = true

		if strings.ContainsRune(motSecret, lettre) {
			fmt.Println("Bonne devinette!")
			for i, char := range motSecret {
				if char == lettre {
					motCache = motCache[:i] + string(char) + motCache[i+1:]
					lettresRestantes--
				}
			}
		} else {
			fmt.Println("Mauvaise devinette!")
			essaisRestants--
		}
	}

	if essaisRestants >= 0 && essaisRestants < len(images) {
		fmt.Println(images[essaisRestants])
	} else {
		fmt.Println("Nombre d'essais non pris en charge.")
	}

	if lettresRestantes == 0 {
		fmt.Println("Félicitations, vous avez gagné! Le mot était :", motSecret)
	} else {
		fmt.Println("Désolé, vous avez perdu. Le mot était :", motSecret)
	}
}
