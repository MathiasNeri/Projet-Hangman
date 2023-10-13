package main

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

func RandomWords(nomFichier string) (string, error) {
	fichier, err := os.Open(nomFichier)
	if err != nil {
		return "", err
	}
	defer fichier.Close()
	var mots []string
	scanner := bufio.NewScanner(fichier)
	for scanner.Scan() {
		mots = append(mots, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	indexAleatoire := rng.Intn(len(mots))
	return mots[indexAleatoire], nil
}

func main() {
	Menu()
}
