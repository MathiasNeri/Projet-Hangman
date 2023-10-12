package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func choisirMotAleatoire(nomFichier string) (string, error) {
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
	nomFichier := "easy.txt"
	mot, err := choisirMotAleatoire(nomFichier)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
		return
	}
	fmt.Printf("Le mot est : %s\n", mot)
}
