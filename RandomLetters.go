package main

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomLetters() {
	// La chaîne de caractères à partir de laquelle vous souhaitez extraire les lettres aléatoires
	maChaine := "votre_chaine_de_caracteres"

	// Initialiser le générateur de nombres aléatoires avec une graine (seed) basée sur le temps actuel
	rand.Seed(time.Now().UnixNano())

	// Générer deux indices aléatoires pour extraire deux lettres aléatoires de la chaîne
	indice1 := rand.Intn(len(maChaine))
	indice2 := rand.Intn(len(maChaine))

	// Créer une nouvelle chaîne de caractères avec des tirets pour cacher les autres caractères
	chaineCachee := ""
	for i := 0; i < len(maChaine); i++ {
		if i == indice1 || i == indice2 {
			chaineCachee += string(maChaine[i])
		} else {
			chaineCachee += "---"
		}
	}

	// Afficher la chaîne cachée avec les deux lettres aléatoires
	fmt.Println(chaineCachee)
}
