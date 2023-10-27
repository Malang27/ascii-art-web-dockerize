package asciifs

import (
	"fmt"
	"os"
	"strings"
)

// Fonction pour lire les fichiers de bannière et les stocker dans une map
func Read(file string) (map[rune][]string, error) {
	input, err := os.ReadFile("utils/" + file)
	if err != nil {
		return nil, fmt.Errorf("erreur de lecture %v", err)
	}
	tab := map[rune][]string{}

	rep := strings.ReplaceAll(string(input), "\r", "")
	div := strings.Split(rep, "\n")

	key := ' '
	lignes := []string{}
	i := 1
	for j, ligne := range div {
		if j == 0 {
			continue
		}
		if i%9 == 0 {
			lignes = append(lignes, ligne)
			tab[key] = lignes
			key++
			lignes = nil
		}
		if ligne != "" {
			lignes = append(lignes, ligne)
		}
		i++
	}

	return tab, nil
}

// Fonction pour générer l'ASCII art en utilisant les fichiers de bannière et le texte
func PrintArt(s string, tab map[rune][]string, result *string) {
	lines := strings.Split(s, "\n")

	for _, line := range lines {
		ma := []rune(s)
		if len(ma) != len(s) {
			*result += ""
			return
		}
		// Parcourir chaque ligne de la map avec une taille maximale de 8
		for i := 0; i < 8; i++ {
			// Parcourir chaque ligne du caractère
			for x := 0; x < len(line); x++ {
				if i < len(tab[rune(line[x])]) {
					*result += tab[rune(line[x])][i]
				} else {
					*result += " "
				}
			}
			*result += "\n"

		}
	}
}
