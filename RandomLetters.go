package projethangman

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func RandomLetters() {
	// La chaîne de caractères à partir de laquelle vous souhaitez extraire les lettres aléatoires
	maChaine := "votre_chaine_de_caracteres"

	for i := 0; i < 2; i++ {
		// Générer deux indices aléatoires pour extraire deux lettres aléatoires de la chaîne
		indice1, _ := rand.Int(rand.Reader, big.NewInt(int64(len(maChaine)))
		indice2, _ := rand.Int(rand.Reader, big.NewInt(int64(len(maChaine)))

		// Obtenir les indices en utilisant Int64
		i1 := indice1.Int64()
		i2 := indice2.Int64()

		// Créer une nouvelle chaîne de caractères avec des tirets pour cacher les autres caractères
		chaineCachee := ""
		for j := 0; j < len(maChaine); j++ {
			if j == int(i1) || j == int(i2) {
				chaineCachee += string(maChaine[j])
			} else {
				chaineCachee += "---"
			}
		}

		// Afficher la chaîne cachée avec les lettres aléatoires
		fmt.Println(chaineCachee)
	}
}