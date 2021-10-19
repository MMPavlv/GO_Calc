package main

import (
	"Calc/Grammar"
	"fmt"
)

func main() {
	g := Grammar.Tokenizer{}
	var result = g.Parse("1/3")

	fmt.Println(result)
}
