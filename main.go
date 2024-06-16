package main

import (
	"log"

	"github.com/Gealber/huff-back/parser"
)

func main() {
	err := parser.Parse("code", "code.huff")
	if err != nil {
		log.Fatal(err)
	}
}
