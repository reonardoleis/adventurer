package terminal

import (
	"fmt"
	"os"
)

type Color struct {
	start string
	end   string
}

// terminal colors scheme
var (
	Red    = Color{"\033[31m", "\033[0m"}
	Green  = Color{"\033[32m", "\033[0m"}
	Blue   = Color{"\033[34m", "\033[0m"}
	Yellow = Color{"\033[33m", "\033[0m"}
	White  = Color{"\033[37m", "\033[0m"}
)

func show(val any, color Color, br bool) {
	f := ""
	if br {
		f = "\n"
	}

	fmt.Print(color.start + fmt.Sprintf("%v", val) + color.end + f)
}

func askFor(text string, color Color) string {
	fmt.Print(text + ": ")
	buff := make([]byte, 512)
	n, _ := os.Stdin.Read(buff)
	return string(buff[:n-2])
}
