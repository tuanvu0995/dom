package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/term"
)

type Buffer struct {
	termiUI  *TermiUI
	filePath *string
	data     []string
}

func NewBuffer(termiUI *TermiUI, filePath *string) *Buffer {
	return &Buffer{
		termiUI:  termiUI,
		filePath: filePath,
		data:     make([]string, 0),
	}
}

func (buffer *Buffer) Load() {
	fh, err := os.Open(*buffer.filePath)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(fh)

	for scanner.Scan() {
		buffer.data = append(buffer.data, scanner.Text())
	}

	fmt.Println(buffer.GetData())
}

func (buffer *Buffer) SetFocus() {
	oldState, err := term.MakeRaw(0)

	if err != nil {
		panic(err)
	}

	defer term.Restore(0, oldState)

	for {
		str, _ := buffer.termiUI.ReadInputByte()
		if str == "q" {
			break
		}

		fmt.Println(str)
	}
}

func (buffer *Buffer) GetData() string {
	return strings.Join(buffer.data, "\n")
}
