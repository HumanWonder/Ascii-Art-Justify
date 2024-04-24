package reverse

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	getstruct "ascii_art/utils/struct"
)

func Ascii_Reverse_Function(file string, AsciiArr []getstruct.ASCIIChar) {
	// variable qui servira à couper les lignes du fichier example afin d'identifier les lettres une par une
	new_start := 0

	var res string
	// fmt.Println(file)

	// Read file and verify if exists, verify also extension ?
	if ExtensionOK(file) {
		// On récupère la structure avec les caractères ascii
		// envoie le tab de struct en paramètres dans le fichier func avant l'appel de cette fiction ici

		// label qui sert à recommencer la boucle range du début
	loop:
		// On range la structure
		for i := range AsciiArr {
			//line correspond au dessin ascii d'un caractère. Split au "\n" pour gérer en 8 étages
			line := strings.Split(AsciiArr[i].Ascii, "\n")
			// test := strings.Split(AsciiArr[101-32].Ascii, "\n")

			// En parcourant la structure on lance un bool pour vérifier si la lettre correspond à celle du fichier
			// return true seulement après avoir checké les 8 lignes verticalement
			if IsMatch(line, file, new_start) {
				// fmt.Println("got match")
				// fmt.Println()
				// On ajoute la valeur ascii associée à la lettre (dans la structure) dans une variable qui sera printée à la fin du programme
				res += AsciiArr[i].Caractere
				// Nouveau départ de ligne pour le fichier pour cibler la prochaine lettre
				new_start += len(line[0])
				// Commande label pour relancer la boucle. En itération classique (i:= 0), faudrait juste remettre i à 0.
				goto loop
			} else {
				// Si bool est false, on descend d'une case dans la structure
				i++
			}

		}
		fmt.Println(res)

	} else {
		log.Fatal("Not a .txt file.")
	}
}

func IsMatch(line_ascii []string, filename string, new_start int) bool {
	data, err := ioutil.ReadFile("examples/" + filename)
	if err != nil {
		log.Fatal("Erreur lors de la lecture du fichier :", err)
	}
	// Var qui permet d'éviter le out of range car len de ligne du fichier sera toujours supérieure à la len d'une ligne de notre police
	var new_end int
	// Split pour individualiser les lignes du fichier après l'avoir lu
	lines_file := strings.Split(string(data), "\n")
	// i correspond au nombre de lignes du fichier qu'on peut calquer avec les lignes de la police
	for i, line := range lines_file {
		// fmt.Println("len file line : ", len(line))
		// Avance la balise de fin avec la valeur de new_start(elle-même égale à la longueur de la ligne à check)
		new_end = len(line_ascii[0]) + new_start
		// Vérif pour éviter le out of range de la fin de ligne (quand on arrive à la dernière lettre)
		if new_end > len(line) {
			return false
		}

		// fmt.Println("police : ", line_ascii[i])
		// fmt.Println("file : ", line[new_start:new_end])

		// Vérif qui permet de sortir de la boucle, c'est le OK final
		if line_ascii[i] == line[new_start:new_end] && i == 7 {
			return true
		}
		// Vérif de base, si la ligne police == ligne fichier (limitée par la longueur de la ligne de police qu'on vérifie)
		if line_ascii[i] == line[new_start:new_end] {
			// fmt.Println("Match...")
			i++
		} else {
			return false
		}
		//}
	}
	return false
}

func ExtensionOK(file string) bool {
	return filepath.Ext(strings.TrimSpace(file)) == ".txt"
}
