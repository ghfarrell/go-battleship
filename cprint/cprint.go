package cprint

import (
	"fmt"

	"github.com/jwalton/go-supportscolor"
)

var Printf map[string]func(string, ...interface{})

func init() {
	colors := []string{
		"Red",
		"Blue",
		"Yellow",
		"White",
		"Magenta",
		"Cyan",
		"Green",
	}
	Printf = make(map[string]func(s string, a ...interface{}))
	if supportscolor.Stdout().SupportsColor {
		Printf["Yellow"] = func(s string, a ...interface{}) {
			p := "\u001b[33;1m" + s + "\u001b[0m"
			fmt.Printf(p, a...)
		}
		Printf["Red"] = func(s string, a ...interface{}) {
			p := "\u001b[31;1m" + s + "\u001b[0m"
			fmt.Printf(p, a...)
		}
		Printf["Cyan"] = func(s string, a ...interface{}) {
			p := "\u001b[36;1m" + s + "\u001b[0m"
			fmt.Printf(p, a...)
		}
		Printf["Magenta"] = func(s string, a ...interface{}) {
			p := "\u001b[35;1m" + s + "\u001b[0m"
			fmt.Printf(p, a...)
		}
		Printf["White"] = func(s string, a ...interface{}) {
			p := "\u001b[37;1m" + s + "\u001b[0m"
			fmt.Printf(p, a...)
		}
		Printf["Green"] = func(s string, a ...interface{}) {
			p := "\u001b[32;1m" + s + "\u001b[0m"
			fmt.Printf(p, a...)
		}
		Printf["Blue"] = func(s string, a ...interface{}) {
			p := "\u001b[34;1m" + s + "\u001b[0m"
			fmt.Printf(p, a...)
		}
	} else {
		for _, c := range colors {
			Printf[c] = func(s string, a ...interface{}) {
				fmt.Printf(s, a...)
			}
		}
	}
}
