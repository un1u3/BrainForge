package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if a Brainfuck program file is provided as an argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run brainfuck_interpreter.go <filename>")
		return
	}

	// Read the Brainfuck program from the file
	filename := os.Args[1]
	program, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Initialize memory (30000 cells) and pointer
	memory := make([]byte, 30000)
	pointer := 0 // memory pointer

	// Execute the Brainfuck program
	for i := 0; i < len(program); i++ {
		switch program[i] {
		case '>':
			// Move pointer to the right
			pointer++
			if pointer >= len(memory) {
				pointer = 0
			}
		case '<':
			// Move pointer to the left
			pointer--
			if pointer < 0 {
				pointer = len(memory) - 1
			}
		case '+':
			// Increment the value at the pointer
			memory[pointer]++
		case '-':
			// Decrement the value at the pointer
			memory[pointer]--
		case '.':
			// Output the value at the pointer as a character
			fmt.Printf("%c", memory[pointer])
		case ',':
			// Read a single byte of input and store it at the pointer
			var input byte
			_, err := os.Stdin.Read([]byte{input})
			if err != nil {
				fmt.Println("Error reading input:", err)
			}
			memory[pointer] = input
		case '[':
			// Jump forward to the matching ']' if the value at the pointer is zero
			if memory[pointer] == 0 {
				// Find the matching ']'
				depth := 1
				for depth > 0 {
					i++
					if i >= len(program) {
						fmt.Println("Error: unmatched '['")
						return
					}
					if program[i] == '[' {
						depth++
					} else if program[i] == ']' {
						depth--
					}
				}
			}
		case ']':
			// Jump back to the matching '[' if the value at the pointer is nonzero
			if memory[pointer] != 0 {
				// Find the matching '['
				depth := 1
				for depth > 0 {
					i--
					if i < 0 {
						fmt.Println("Error: unmatched ']'")
						return
					}
					if program[i] == ']' {
						depth++
					} else if program[i] == '[' {
						depth--
					}
				}
			}
		}
	}
}
