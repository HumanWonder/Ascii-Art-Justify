package getstruct

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type ASCIIChar struct {
	Ascii     string
	Caractere string
	Height    int
	Width     int
}

func FindAsciiChars(filename string) []ASCIIChar {
	var buffer string

	// ouvre le fichier
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	// lit les données
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	// tableau de struct pour associer les caractères aux ascii
	AsciiChars := make([]ASCIIChar, 0)
	// ID du tableau ascii, on démarre par l'espace
	AsciiNumber := 32
	startline := false
	height := 0

	// Split pour lire ligne par ligne
	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		line = strings.ReplaceAll(line, "\r", "")
		if countLeadingSpaces(line) > 0 {
			buffer += line + "\n"
			height++
			startline = true
		} else if startline && countLeadingSpaces(buffer) > 0 {
			asciiChar := ASCIIChar{
				Ascii:     buffer,
				Caractere: string(rune(AsciiNumber)),
				Height:    height,
				Width:     len(buffer) / 8,
			}
			AsciiNumber++
			AsciiChars = append(AsciiChars, asciiChar)
			buffer = ""
			height = 0
			startline = false
		}
	}
	return AsciiChars
}

func countLeadingSpaces(line string) int {
	return len(line) + len(strings.TrimLeft(line, " "))
}
