# BrainForge
## Brainfuck Interpreter in Go
```
 ____  _____            _____ _   _   ______ ____  _____   _____ ______ 
 |  _ \|  __ \     /\   |_   _| \ | | |  ____/ __ \|  __ \ / ____|  ____|
 | |_) | |__) |   /  \    | | |  \| | | |__ | |  | | |__) | |  __| |__   
 |  _ <|  _  /   / /\ \   | | | . ` | |  __|| |  | |  _  /| | |_ |  __|  
 | |_) | | \ \  / ____ \ _| |_| |\  | | |   | |__| | | \ \| |__| | |____ 
 |____/|_|  \_\/_/    \_\_____|_| \_| |_|    \____/|_|  \_\\_____|______|
```
 

This project is a simple **BrainForge** written in Go. It reads a Brainfuck program from a file and executes it. The interpreter simulates the Brainfuck language's memory and pointer manipulation to output the result of the program.

## Requirements

- Go 1.18 or higher

## Usage

To use the Brainfuck interpreter, follow these steps:

1. Clone the repository or download the code.

2. Create a Brainfuck program file (e.g., `hello_world.bf`).

3. Run the interpreter with the following command:

   ```bash
   go run compiler.go <filename>
