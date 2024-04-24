package algo

import (
	"flag"
	"fmt"
	"log"
	"strings"

	ascii_color "ascii_art/colors"
	getconsole "ascii_art/utils/align"
	ascii_reverse "ascii_art/utils/reverse"
	getstruct "ascii_art/utils/struct"
)

// Le charactère ascii correspond à un caractère d'un fichier font
type ASCIIChar struct {
	Ascii     string
	Caractere string
	Height    int
	Width     int
}

// Struct ne se fera que si l'option align est active
type Options struct {
	Calc     int
	SpacePos []int
}

var alignment string

func Aiguillage(in []string) {
	font := "font/standard.txt"
	// Variable crééé pour le flag car sinon il ne le reconnaît pas
	var color_in string
	reverse := flag.String("reverse", "", "Usage: go run . [OPTION]\n\nEX: go run . --reverse=<fileName>")
	flag.StringVar(&color_in, "color", "", "Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
	justify := flag.String("align", "", "Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right \"something\" standard")
	flag.Parse()

	// Condition qui aiguille vers le ascii de base s'il n'y a qu'un argument
	if len(in) > 1 {
		// check si un flag est utilisé
		if color_in == "" {
			if *justify == "" {
				/*S'il ne l'est pas et que la len de l'input est supérieure à 1,
				Cela veut dire qu'il y a une police de spécifiée. On check alors la police.
				OU le flag reverse est mal renseigné (on le check en premier)*/
				if *reverse != "" {
					fmt.Println("Usage: go run . [OPTION]\n\nEX: go run . --reverse=<fileName>")
					return
				}
				if FontValid(in[1]) {
					font = "font/" + in[1] + ".txt"
				} else {
					log.Fatal("Please check the spelling of the font you're asking for.")
				}

				Ascii_Art(in[0], font)
				// s'il tombe dans le else, le flag align est validé
			} else {
				// On recheck le nombre d'arguments, si input > 2 nous avons une police de renseignée
				if *justify != "" && len(in) > 2 {
					// Merci à l'exercice d'interdire le format accepté du package flag et de nous imposer un seul format d'input
					if len(in) > 3 {
						fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right \"something\" standard")
						return
					}
					// On check en même temps la valeur du flag (s'il s'agit d'une bonne option)
					if getconsole.AlignValid(*justify) && FontValid(flag.Arg(1)) {
						font = "font/" + flag.Arg(1) + ".txt"
					} else {
						fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right \"something\" standard")
						return
					}
				} else {
					// Ici pas de police, on check juste la valeur du flag
					if !getconsole.AlignValid(*justify) {
						fmt.Println("Usage: align flag values are <left>/<right>/<center>/<justify>")
						return
					}
				}
				// c'est parti dans l'ascii yay
				Ascii_Art(flag.Arg(0), font, *justify)
			}
			// si le nombre d'arguments après le flag reste sup à 1 (ou égal), en sachant que nous avons déjà éliminé le flag de justify, il ne reste que color
		} else if len(flag.Args()) >= 1 {
			/* Pour régler le problème d'une seule écriture possible pour le flag
			on vérifie si l'argument pris en compte par le flag n'est pas l'argument suivant :
			celui qui correspondrait à la couleur.*/

			// fmt.Println(color_in, "= color_in; input1 =", in[1])
			if color_in == in[1] {
				fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
				return
			}
			// on check alors le nombre d'args après le flag. 1 = string entière
			if len(flag.Args()) == 1 {
				letters_to_COLOR := ""
				to_print := in[1]
				arg := in[1]
				Ascii_ArtForColor(arg, font, &color_in, letters_to_COLOR, to_print)
				// 3 = lettre(s) spéciales à colorier
			} else if len(flag.Args()) >= 2 {
				letters_to_COLOR := in[1]
				to_print := in[2]
				arg := in[2]
				Ascii_ArtForColor(arg, font, &color_in, letters_to_COLOR, to_print)
			}
		} else {
			fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
			return
		}
		// Ascii art de base
	} else if len(in) == 1 {
		// fmt.Println(*reverse, len(in))
		// On check si le flag est utilisé, montre erreur d'utilisation si c'est le cas.
		if color_in == "" && *reverse == "" && *justify == "" {
			Ascii_Art(in[0], "font/standard.txt")
			//*reverse = valeur string du flag =/= flag args qui est l'arg qui vient après
			//Comme ici on est à len = 1, on a pas de flag.Args
		} else if *reverse != "" {
			AsciiArr := getstruct.FindAsciiChars("font/standard.txt")
			file := *reverse
			ascii_reverse.Ascii_Reverse_Function(file, AsciiArr)
		} else if *justify != "" {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right \"something\" standard")
			return
		} else {
			fmt.Println("Usage: go run . [OPTION]\n\nEX: go run . --reverse=<fileName>")
			return
		}
	}
}

/*Choper biblio standard pour individualiser les caractères ascii.
Appliquer chaque caractère */

func Ascii_Art(arg string, font string, opt ...string) {
	// fmt.Println(opt)
	// fmt.Println(arg)
	var NewLines []int

	AsciiArr := getstruct.FindAsciiChars(font)
	ArgArr := ArgBecomesArr(arg, &NewLines)
	// check ascii :
	// fmt.Println(ArgArr)
	// fmt.Println("Ascii Char :")
	// fmt.Println(AsciiArr[33-32].Ascii)
	// fmt.Println("Utf-8 char :")
	// fmt.Println(AsciiArr[33-32].Caractere)
	// fmt.Println("Height :")
	// fmt.Println(AsciiArr[33-32].Height)
	// fmt.Println("Width :")
	// fmt.Println(AsciiArr[33-32].Width)

	AsciiArrMatched := MatchAsciiToArg(AsciiArr, ArgArr)

	start := 0
	var end int
	for i := 0; i <= len(NewLines); i++ {
		if i == len(NewLines) {
			end = len(AsciiArrMatched)
		} else {
			end = NewLines[i]
		}
		if len(AsciiArrMatched) > 0 {
			// on découpe selon newlines = bornes pour slices grâce à start et end
			ArgToPrint := SepLinesToPrint(strings.Join(AsciiArrMatched[start:end], ""))
			// opt correspond à la valeur du flag passée en paramètre
			if len(opt) == 0 {
				PrintLine(ArgToPrint)
			} else if opt[0] == "left" {
				PrintLine(ArgToPrint)
			} else if len(opt) > 0 {
				// selon valeur, PrintLine effectura un calcul différent pour l'afficher à droite ou au centre par ex.
				switch opt[0] {
				case "center":
					alignment = "center"
					PrintLine(ArgToPrint)
				case "right":
					alignment = "right"
					PrintLine(ArgToPrint)
				case "justify":
					alignment = "justify"
					// Nospace = si utilisation d'espaces dans l'argument
					// la func ci-dessous calcule le nombre d'espaces nécessaires entre chaque mot (pour l'option justify seulement)
					NoSpace, FinalOptStruct := CalcSpaces(AsciiArr[0], arg, ArgToPrint)
					if NoSpace {
						// Si pas d'escpace utilisé, le comportement redevient celui d'un alignement left
						alignment = "left"
					}
					// for i := 0; i < 8; i++ {
					PrintLine(ArgToPrint, FinalOptStruct)
					//}
					// fmt.Println("justify option")
					// w terminal - nb caractères / nb espaces

				}
			}

		}

		start = end
	}
	if len(NewLines) != 0 && arg == "\\n" {
		fmt.Println()
	}
}

func Ascii_ArtForColor(arg string, font string, color_in *string, letters_to_COLOR, to_print string) {
	var NewLines []int
	// fmt.Println("Arg pris en compte : ", arg)
	// Regarde si une font est demandée par la longueur de l'input
	// Check si la font existe dans nos fichiers

	arr_color_done := ArrayColor(color_in, letters_to_COLOR, to_print)
	AsciiArr := getstruct.FindAsciiChars(font)
	ArgArr := ArgBecomesArr(arg, &NewLines)
	// fmt.Println(ArgArr)
	// fmt.Println("Ascii Char :")
	// fmt.Println(AsciiArr[90-32].Ascii)
	// fmt.Println("Utf-8 char :")
	// fmt.Println(AsciiArr[90-32].Caractere)
	// fmt.Println("Height :")
	// fmt.Println(AsciiArr[90-32].Height)

	AsciiArrMatched := MatchAsciiToArg(AsciiArr, ArgArr)
	// fmt.Println(NewLines)
	// fmt.Print(AsciiArrMatched)

	start := 0
	var end int
	for i := 0; i < len(NewLines)+1; i++ {
		if i == len(NewLines) {
			end = len(AsciiArrMatched)
		} else {
			end = NewLines[i]
		}
		if len(AsciiArrMatched) > 0 {
			ArgToPrint := SepLinesToPrint(strings.Join(AsciiArrMatched[start:end], ""))
			for j := 0; j < 8; j++ {
				PrintLineColor(ArgToPrint, j, arr_color_done)
			}
		}

		start = end
	}
	if len(NewLines) != 0 && arg == "\\n" {
		fmt.Println()
	}
}

func FontValid(input string) bool {
	switch input {
	case "standard":
		return true
	case "shadow":
		return true
	case "thinkertoy":
		return true
	case "varsity":
		return true
	}
	return false
}

func ArgBecomesArr(arg string, NewLines *[]int) []string {
	FinalArr := make([]string, 0)
	buffer := ""
	NewLinePosition := 0
	SkippedN := false

	for i := 0; i < len(arg); i++ {
		if arg[i] == 92 && arg[i+1] == 110 {
			FinalArr = append(FinalArr, buffer)
			buffer = ""
			SkippedN = true
			*NewLines = append(*NewLines, NewLinePosition)
			// fmt.Println(*NewLines)
		} else {
			if SkippedN {
				SkippedN = false
			} else {
				NewLinePosition++
				buffer += string(arg[i])
			}
		}
		if i == len(arg)-1 {
			FinalArr = append(FinalArr, buffer)
		}
	}

	return FinalArr
}

func MatchAsciiToArg(AsciiCharas []getstruct.ASCIIChar, ArgArr []string) []string {
	var FinalArr []string
	for i := 0; i < len(ArgArr); i++ {
		for k := 0; k < len(ArgArr[i]); k++ {
			// 	asciiArray = append(asciiArray, asciiChars[arrayText[i][k]-32].Ascii)
			FinalArr = append(FinalArr, AsciiCharas[ArgArr[i][k]-32].Ascii)
			// fmt.Println(FinalArr)
		}
	}
	return FinalArr
}

func SepLinesToPrint(str string) []string {
	var SeparatedAscii []string
	buffer := ""

	for i := 0; i < len(str); i++ {
		if str[i] == '\n' {
			SeparatedAscii = append(SeparatedAscii, buffer)
			buffer = ""
		} else {
			buffer += string(str[i])
		}
	}

	return SeparatedAscii
}

func PrintLine(tab []string, Options ...Options) {
	var strBuilder strings.Builder
	if len(tab) > getconsole.GetConsoleSize() {
		log.Fatal("Terminal size too small")
	}
	// fmt.Println("len tab :", len(tab))

	if alignment == "right" {
		for line := 0; line < 8; line++ {
			i := line
			for i < len(tab) {
				strBuilder.WriteString(tab[i])
				// fmt.Println(tab[i : i+7])
				i += 8
			}
			// Nb d'espace + longueur str = longueur terminal
			fmt.Println(strings.Repeat(" ", getconsole.GetConsoleSize()-(len(strBuilder.String()))) + strBuilder.String())
			strBuilder.Reset()
		}
	} else if alignment == "center" {
		for line := 0; line < 8; line++ {
			i := line
			for i < len(tab) {
				strBuilder.WriteString(tab[i])
				// fmt.Println(tab[i : i+7])
				i += 8
			}
			// fmt.Println(strBuilder.String())
			fmt.Println(strings.Repeat(" ", getconsole.GetConsoleSize()/2-(len(strBuilder.String())/2)) + strBuilder.String())
			strBuilder.Reset()
		}
	} else if alignment == "justify" {
		// fmt.Println(len(tab))
		// strbuff := strBuilder.String()
		spacebuff := strings.Repeat(" ", Options[0].Calc)
		for line := 0; line < 8; line++ {
			// fmt.Println("Reset, next line :", line)
			i := line

			for i < len(tab) {
				var check bool
				// fmt.Println("i = ", i, i/8)
				for ind := range Options[0].SpacePos {
					// fmt.Println("Checking : ", i, Options[0].SpacePos[ind]*8+line)
					if i == Options[0].SpacePos[ind]*8+line {
						check = true
					}
					// fmt.Println(Options[0].SpacePos[ind]*8 + line)
				}
				if check {
					strBuilder.WriteString(spacebuff)
				} else {
					strBuilder.WriteString(tab[i])
				}
				i += 8
			}
			// for j := range Options[0].SpacePos {
			// 	Options[0].SpacePos[j]++
			// }
			fmt.Println(strBuilder.String())
			strBuilder.Reset()
		}
		// fmt.Println(strBuilder.String())

	} else {
		for line := 0; line < 8; line++ {
			i := line
			for i < len(tab) {
				strBuilder.WriteString(tab[i])
				// fmt.Println(tab[i : i+7])
				i += 8
			}
			fmt.Println(strBuilder.String())
			strBuilder.Reset()
		}
		// fmt.Println(count)
	}
}

func CalcSpaces(SpaceStruct getstruct.ASCIIChar, arg string, ArgtoPrint []string) (bool, Options) {
	var count int
	var spacePos []int
	var NoSpace bool
	FinalOptStruct := Options{}

	for j := range arg {
		if arg[j] == 32 {
			count++
			spacePos = append(spacePos, j)
		}
	}
	if count == 0 {
		NoSpace = true
		return NoSpace, FinalOptStruct

		// log.Fatal("Please use spaces with the justify option.")
	}
	var strBuilder strings.Builder
	for i := 0; i < 8; i++ {
		for i < len(ArgtoPrint) {
			strBuilder.WriteString(ArgtoPrint[i])
			// fmt.Println(tab[i : i+7])
			i += 8
		}
	}

	// fmt.Println("spacePos : ", spacePos)
	// fmt.Println(count, SpaceStruct.Width)
	// fmt.Println(strBuilder.String())
	// fmt.Println(strBuilder.String())
	// fmt.Println("Spaces :", count)
	// fmt.Println("Terminal size :", getconsole.GetConsoleSize())
	// fmt.Println("strBuilder len :", len(strBuilder.String()))
	// fmt.Println("Terminal size unused :", getconsole.GetConsoleSize()-len(strBuilder.String()))
	// fmt.Println("Number of spaces :", (getconsole.GetConsoleSize()-len(strBuilder.String()))/count)
	res := ((getconsole.GetConsoleSize() - strBuilder.Len()) + (6 * count)) / count
	FinalOptStruct = Options{
		Calc:     res,
		SpacePos: spacePos,
	}
	return NoSpace, FinalOptStruct

	// fmt.Println(len(ArgtoPrint) - (SpaceStruct.Width * count))
}

func PrintLineColor(tab []string, i int, arr_color []string) {
	var strBuilder strings.Builder
	count := 0
	// fmt.Println(tab, arr_color, len(tab), len(arr_color))

	for i < len(tab) {
		switch arr_color[count] {
		case "red":
			strBuilder.WriteString(ascii_color.Red + tab[i] + ascii_color.Reset)
		case "green":
			strBuilder.WriteString(ascii_color.Green + tab[i] + ascii_color.Reset)
		case "yellow":
			strBuilder.WriteString(ascii_color.Yellow + tab[i] + ascii_color.Reset)
		case "blue":
			strBuilder.WriteString(ascii_color.Blue + tab[i] + ascii_color.Reset)
		case "purple":
			strBuilder.WriteString(ascii_color.Purple + tab[i] + ascii_color.Reset)
		case "cyan":
			strBuilder.WriteString(ascii_color.Cyan + tab[i] + ascii_color.Reset)
		case "white":
			strBuilder.WriteString(ascii_color.White + tab[i] + ascii_color.Reset)
		case "orange":
			strBuilder.WriteString(ascii_color.Orange + tab[i] + ascii_color.Reset)
		case "none":
			strBuilder.WriteString(ascii_color.Reset + tab[i] + ascii_color.Reset)
		}
		i += 8
		count++
	}
	// fmt.Println(count)
	fmt.Println(strBuilder.String())
}

func ArrayColor(color_in *string, letters_to_COLOR, to_print string) []string {
	var color_arr []string
	if letters_to_COLOR == "" {
		for i := 0; i < len(to_print); i++ {
			color_arr = append(color_arr, *color_in)
		}
	} else {
		for i := 0; i < len(to_print); i++ {
			if strings.Contains(letters_to_COLOR, string(to_print[i])) {
				color_arr = append(color_arr, *color_in)
			} else {
				color_arr = append(color_arr, "none")
			}
		}
	}
	// fmt.Println(*color_in, letters_to_COLOR, to_print)
	return color_arr
}
