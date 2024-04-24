package getconsole

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"unicode/utf8"
)

func AlignValid(opt string) bool {
	switch opt {
	case "center":
		return true
	case "left":
		return true
	case "right":
		return true
	case "justify":
		return true
	}
	return false
}

func GetConsoleSize() int {
	cmd := exec.Command("stty", "size") // stty (shell command) returns the screen dimensions.
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	s := string(out)
	s = strings.TrimSpace(s)
	sArr := strings.Split(s, " ")

	// height, err := strconv.Atoi(sArr[0])
	// if err != nil {
	// 	log.Fatal(err)
	// }
	width, err := strconv.Atoi(sArr[1])
	if err != nil {
		log.Fatal(err)
	}
	return width
}

// NCenter centers the string to the column width.
func NCenter(width int, s string) *bytes.Buffer {
	const half, space = 2, "\u0020"
	var b bytes.Buffer
	n := (width - utf8.RuneCountInString(s)) / half
	if n < 1 {
		fmt.Fprintf(&b, s)
		return &b
	}

	fmt.Fprintf(&b, "%s%s", strings.Repeat(space, int(n)), s)

	return &b
}

// Center the string to the width of the terminal.
// When the width is unknown, the string is left-aligned.
func Center(s string, w int) *bytes.Buffer {
	if w <= 0 {
		return NCenter(0, s)
	}
	return NCenter(w, s)
}
