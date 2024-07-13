package main

import (
	"fmt"
	"os"
)

type TokenType int

const (
	LEFT_PAREN TokenType = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	STAR
	EOF
)

type Token struct {
	Type   TokenType
	Lexeme string
}

func (t TokenType) String() string {
	tokenNames := []string{"LEFT_PAREN", "RIGHT_PAREN", "LEFT_BRACE", "RIGHT_BRACE", "COMMA", "DOT", "MINUS", "PLUS", "SEMICOLON", "STAR", "EOF"}
	if t < LEFT_PAREN || t > EOF {
		return "Unknown"
	}
	return tokenNames[t]
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Fprintln(os.Stderr, "Logs from your program will appear here!")

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]

	if command != "tokenize" {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	// Uncomment this block to pass the first stage
	//
	filename := os.Args[2]
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	if len(fileContents) > 0 {

		for _, val := range fileContents {
			char := string(val)
			var t Token
			switch char {
			case "(":
				t = Token{
					Type:   TokenType(0),
					Lexeme: "(",
				}

			case ")":
				t = Token{
					Type:   TokenType(1),
					Lexeme: ")",
				}
			case "{":
				t = Token{
					Type:   TokenType(2),
					Lexeme: "{",
				}
			case "}":
				t = Token{
					Type:   TokenType(3),
					Lexeme: "}",
				}
			case ",":
				t = Token{
					Type:   TokenType(4),
					Lexeme: ",",
				}
			case ".":
				t = Token{
					Type:   TokenType(5),
					Lexeme: ".",
				}
			case "-":
				t = Token{
					Type:   TokenType(6),
					Lexeme: "-",
				}
			case "+":
				t = Token{
					Type:   TokenType(7),
					Lexeme: "+",
				}
			case ";":
				t = Token{
					Type:   TokenType(8),
					Lexeme: ";",
				}
			case "*":
				t = Token{
					Type:   TokenType(9),
					Lexeme: "*",
				}
			}

			fmt.Println(t.Type.String(), t.Lexeme, "null")
		}
		fmt.Println("EOF  null")
	} else {
		fmt.Println("EOF  null") // Placeholder, remove this line when implementing the scanner
	}
}
