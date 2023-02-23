package main

type TokenType int

const (
	DELIM TokenType = iota
)

type TokenValue string

type Token struct {
	Type  TokenType
	Value TokenValue
}
