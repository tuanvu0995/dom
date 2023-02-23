package main

import (
	"io"
	"os"
)

type TermiUI struct {
	Writer io.Writer
	Reader io.Reader
}

func NewTermiUI() *TermiUI {
	return &TermiUI{
		Writer: os.Stdout,
		Reader: os.Stdin,
	}
}

func (termiUI *TermiUI) ReadInputByte() (string, error) {
	var rBuf []byte
	var err error

	for {
		var buf [1]byte
		_, err = termiUI.Reader.(*os.File).Read(buf[:])

		if err != nil && err != io.EOF {
			break
		}

		rBuf = append(rBuf, buf[0])

		if len(rBuf) > 0 {
			break
		}
	}

	return string(rBuf), nil
}
