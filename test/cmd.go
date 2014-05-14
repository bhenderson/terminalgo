package main

import (
	"fmt"
	t "github.com/bhenderson/terminalgo"
	"os"
)

func main() {
	if t.IsTerminal(os.Stdin.Fd()) {
		fmt.Println("stdin is a terminal")
	}
	if t.IsTerminal(os.Stdout.Fd()) {
		fmt.Println("stdout is a terminal")
	}
}
