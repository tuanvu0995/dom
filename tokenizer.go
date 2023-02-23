package main

type Tokenizer struct {
	program *Program
}

func NewTokenizer(program *Program) *Tokenizer {
	return &Tokenizer{
		program: program,
	}
}

func (tokenizer *Tokenizer) GenerateTokens() chan *Token {
	out := make(chan *Token)
	lines := tokenizer.program.GenerateLines()

	go func() {
		defer close(out)

		for line := range lines {
			out <- tokenizer.ParseToken(line)
		}
	}()
	return out
}

func (tokenizer *Tokenizer) ParseToken(line *LineOfCode) *Token {
	return &Token{}
}
