package board

import (
	"fmt"

	"github.com/ghfarrell/go-battleship/cprint"
)

// clearscreen and printlogo functions have to be in this package so that I can use
// them in the ship placement loop :/
func PrintLogo() {
	cprint.Printf["Magenta"](" /$$$$$$$   /$$$$$$  /$$$$$$$$/$$$$$$$$/$$       /$$$$$$$$  \n")
	cprint.Printf["Magenta"]("| $$__  $$ /$$__  $$|__  $$__/__  $$__/ $$      | $$_____/ \n")
	cprint.Printf["Magenta"]("| $$  \\ $$| $$  \\ $$   | $$     | $$  | $$      | $$     \n")
	cprint.Printf["Magenta"]("| $$$$$$$ | $$$$$$$$   | $$     | $$  | $$      | $$$$$   \n")
	cprint.Printf["Magenta"]("| $$__  $$| $$__  $$   | $$     | $$  | $$      | $$__/    \n")
	cprint.Printf["Magenta"]("| $$  \\ $$| $$  | $$   | $$     | $$  | $$      | $$       \n")
	cprint.Printf["Magenta"]("| $$$$$$$/| $$  | $$   | $$     | $$  | $$$$$$$$| $$$$$$$$| \n")
	cprint.Printf["Magenta"]("|_______/ |__/  |__/   |__/     |__/  |________/|________/  \n")
	cprint.Printf["Magenta"]("  /$$$$$$  /$$   /$$ /$$$$$$ /$$$$$$$\n")
	cprint.Printf["Magenta"](" /$$__  $$| $$  | $$|_  $$_/| $$__  $$\n")
	cprint.Printf["Magenta"]("| $$  \\__/| $$  | $$  | $$  | $$  \\ $$\n")
	cprint.Printf["Magenta"]("|  $$$$$$ | $$$$$$$$  | $$  | $$$$$$$/\n")
	cprint.Printf["Magenta"](" \\____  $$| $$__  $$  | $$  | $$____/ \n")
	cprint.Printf["Magenta"](" /$$  \\ $$| $$  | $$  | $$  | $$     \n")
	cprint.Printf["Magenta"]("|  $$$$$$/| $$  | $$ /$$$$$$| $$      \n")
	cprint.Printf["Magenta"](" \\______/ |__/  |__/|______/|__/  \n")
	cprint.Printf["Red"]("Press ctrl + c at any time to quit.\n\n")
}
func ClearScreen() {
	fmt.Print("\033[H\033[2J")
	PrintLogo()
}
