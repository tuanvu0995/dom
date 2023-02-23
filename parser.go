package main

type ASTNode struct {
}

type AST struct {
	nodes []ASTNode
}

type Parser struct {
	tokenizer *Tokenizer
}

func NewParser(tokenizer *Tokenizer) *Parser {
	return &Parser{
		tokenizer: tokenizer,
	}
}

func (parser *Parser) BuildAST() *AST {
	var ast AST

	for token := range parser.tokenizer.GenerateTokens() {
		_ = token
	}

	return &ast
}
