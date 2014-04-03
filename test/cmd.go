package main

import (
	"fmt"
	t "github.com/bhenderson/terminal"
)

func main() {
	if t.IsTerminal(0) {
		fmt.Println("stdin is a terminal")
	}
	if t.IsTerminal(1) {
		fmt.Println("stdout is a terminal")
	}
}
