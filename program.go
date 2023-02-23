package main

import (
	"bufio"
	"log"
	"os"
)

type LineOfCode string

type Program struct {
	fileHandler *os.File
}

func NewProgram(filePath string) *Program {
	fh, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	return &Program{
		fileHandler: fh,
	}
}

func (program *Program) GenerateLines() chan *LineOfCode {
	out := make(chan *LineOfCode)
	scanner := bufio.NewScanner(program.fileHandler)

	go func() {
		defer program.fileHandler.Close()
		defer close(out)

		for scanner.Scan() {
			loc := LineOfCode(scanner.Text())
			out <- &loc
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}()
	return out
}
