package main

import (
	"fmt"

	"gopkg.in/russross/blackfriday.v2"
)

func main() {
	unsafe := blackfriday.Run([]byte("### hello"), func(m *blackfriday.Markdown) {
	})

	fmt.Print(string(unsafe))
}
