package main

import (
	"os"

	"github.com/dobuzora/learn-words-cl/pkg/words"
)

func main() {
	var (
		reader   = os.Stdin
		writer   = os.Stdout
		filename = "list"
	)

	g := words.New(reader, writer, filename)
	g.Do()
}
