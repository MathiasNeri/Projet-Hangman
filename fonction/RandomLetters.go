package fonction

import (
	"crypto/rand"
	"math/big"
)

func RandomLetters(maChaine string) string {
	// Créer une nouvelle chaîne de caractères avec des tirets pour cacher les lettres
	chaineCachee := ""
	for j := 0; j < len(maChaine); j++ {
		chaineCachee += "-"
	}

	// Générer deux indices aléatoires pour extraire deux lettres aléatoires de la chaîne
	indice1, _ := rand.Int(rand.Reader, big.NewInt(int64(len(maChaine))))
	indice2, _ := rand.Int(rand.Reader, big.NewInt(int64(len(maChaine))))

	// Obtenir les indices en utilisant Int64
	i1 := indice1.Int64()
	i2 := indice2.Int64()

	// Remplacer les tirets par les lettres aléatoires aux indices correspondants
	chaineCachee = chaineCachee[:i1] + string(maChaine[i1]) + chaineCachee[i1+1:]
	chaineCachee = chaineCachee[:i2] + string(maChaine[i2]) + chaineCachee[i2+1:]

	// Afficher la chaîne cachée avec les lettres aléatoires
	return chaineCachee
}
