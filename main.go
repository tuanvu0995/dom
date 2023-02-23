package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	editFile := flag.String("e", "test.txt", "Edit a file")
	flag.Parse()

	prog := os.Args[1]

	if editFile != nil {
		termiUI := NewTermiUI()
		buffer := NewBuffer(termiUI, editFile)
		buffer.Load()
		buffer.SetFocus()
		prog = buffer.GetData()
	}

	program := NewProgram(prog)
	tokenizer := NewTokenizer(program)
	parser := NewParser(tokenizer)
	evaluator := NewEvaluator(parser)

	for output := range evaluator.GenerateOutput() {
		fmt.Println(&output)
	}
}
