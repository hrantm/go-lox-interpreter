package main

import (
	"bufio"
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
	error := false
	// if len(fileContents) > 0 {
	for scanner.Scan() {
		line := scanner.Text()

		for _, val := range line {
			// line_num := i + 1
			line_num := 1
			// fmt.Println(line_num, string(val))
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
			default:
				error = true
				message := fmt.Sprintf("[line %v] Error: Unexpected character: %v", line_num, char)

				fmt.Fprintln(os.Stderr, message)
			}
			if t.Lexeme != "" {
				fmt.Fprintln(os.Stdout, t.Type.String(), t.Lexeme, "null")
			}

		}
	}
	fmt.Println("EOF  null")
	if error {
		os.Exit(65)
	}

}
