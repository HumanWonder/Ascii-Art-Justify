package main

import (
	"os"

	ascii "ascii_art/utils/algo"
)

func main() {
	ascii.Aiguillage(os.Args[1:])
}
