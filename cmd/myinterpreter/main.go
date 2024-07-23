package main

import (
	"bufio"
	"fmt"
	"os"
)

type TokenType int

const (
	UNKNOWN TokenType = iota
	LEFT_PAREN
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	STAR
	EQUAL
	EQUAL_EQUAL
	BANG
	BANG_EQUAL
	LESS
	LESS_EQUAL
	GREATER
	GREATER_EQUAL
	SLASH
	EOF
)

type Token struct {
	Type   TokenType
	Lexeme string
}

func (t TokenType) String() string {
	tokenNames := []string{"UNKNOWN", "LEFT_PAREN", "RIGHT_PAREN", "LEFT_BRACE", "RIGHT_BRACE", "COMMA", "DOT", "MINUS", "PLUS", "SEMICOLON", "STAR", "EQUAL", "EQUAL_EQUAL", "BANG", "BANG_EQUAL", "LESS", "LESS_EQUAL", "GREATER", "GREATER_EQUAL", "SLASH", "EOF"}
	if t < LEFT_PAREN || t > EOF {
		return "Unknown"
	}
	return tokenNames[t]
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	// fmt.Fprintln(os.Stderr, "Logs from your program will appear here!")

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
	// filename := os.Args[2]
	// fileContents, err := os.ReadFile(filename)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
	// 	os.Exit(1)
	// }
	filename := os.Args[2]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	scanner := bufio.NewScanner(file)
	const maxCapacity = 1024 * 1024 // 1 MB
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	// error := false
	// // if len(fileContents) > 0 {
	line_num := 1
	error := false

	for scanner.Scan() {
		line := scanner.Text()
		var t Token
		skipNext := 0
		for i, val := range line {
			if skipNext > 0 {
				skipNext--
				continue
			}
			t = getToken(val)

			if t.Type == 11 && i+1 < len(line) {
				nextTok := getToken(rune(line[i+1]))
				if nextTok.Type == 11 {
					t = Token{
						Type:   TokenType(12),
						Lexeme: "==",
					}
					skipNext++
				}
			}

			if t.Type == 13 && i+1 < len(line) {
				nextTok := getToken(rune(line[i+1]))
				if nextTok.Type == 11 {
					t = Token{
						Type:   TokenType(14),
						Lexeme: "!=",
					}
					skipNext++
				}
			}

			if t.Type == 15 && i+1 < len(line) {
				nextTok := getToken(rune(line[i+1]))
				if nextTok.Type == 11 {
					t = Token{
						Type:   TokenType(16),
						Lexeme: "<=",
					}
					skipNext++
				}
			}

			if t.Type == 17 && i+1 < len(line) {
				nextTok := getToken(rune(line[i+1]))
				if nextTok.Type == 11 {
					t = Token{
						Type:   TokenType(18),
						Lexeme: ">=",
					}
					skipNext++
				}
			}

			if t.Type == 19 && i+1 < len(line) {
				nextTok := getToken(rune(line[i+1]))
				if nextTok.Type == 19 {
					break
				}
			}

			if t.Type != 0 {
				fmt.Fprintln(os.Stdout, t.Type.String(), t.Lexeme, "null")
			} else {
				message := fmt.Sprintf("[line %v] Error: Unexpected character: %v", line_num, string(val))

				fmt.Fprintln(os.Stderr, message)
				error = true

			}
		}
		line_num++
	}
	fmt.Println("EOF  null")
	if error {
		os.Exit(65)
	}
}

func getToken(c rune) Token {
	var t Token
	switch c {
	case '(':
		t = Token{
			Type:   TokenType(1),
			Lexeme: "(",
		}

	case ')':
		t = Token{
			Type:   TokenType(2),
			Lexeme: ")",
		}

	case '{':
		t = Token{
			Type:   TokenType(3),
			Lexeme: "{",
		}
	case '}':
		t = Token{
			Type:   TokenType(4),
			Lexeme: "}",
		}
	case ',':
		t = Token{
			Type:   TokenType(5),
			Lexeme: ",",
		}
	case '.':
		t = Token{
			Type:   TokenType(6),
			Lexeme: ".",
		}
	case '-':
		t = Token{
			Type:   TokenType(7),
			Lexeme: "-",
		}
	case '+':
		t = Token{
			Type:   TokenType(8),
			Lexeme: "+",
		}
	case ';':
		t = Token{
			Type:   TokenType(9),
			Lexeme: ";",
		}
	case '*':
		t = Token{
			Type:   TokenType(10),
			Lexeme: "*",
		}
	case '=':
		t = Token{
			Type:   TokenType(11),
			Lexeme: "=",
		}
	case '!':
		t = Token{
			Type:   TokenType(13),
			Lexeme: "!",
		}
	case '<':
		t = Token{
			Type:   TokenType(15),
			Lexeme: "<",
		}
	case '>':
		t = Token{
			Type:   TokenType(17),
			Lexeme: ">",
		}
	case '/':
		t = Token{
			Type:   TokenType(19),
			Lexeme: "/",
		}
	}

	return t
}
