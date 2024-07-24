# TinyBASIC Compiler
The TinyBASIC Compiler, written in Go, is a hobby project inspired by Austin Z. Henley's three-part article ["Let's make a Teeny Tiny compiler"](https://austinhenley.com/blog/teenytinycompiler1.html).

## Main Idea
The main idea of this project is to create a compiler that can compile TinyBASIC code into C code and then compile it using GCC.

The program contains a lexer, parser, and code emitter. The lexer reads the input and tokenizes it. The parser reads the tokens and generates an abstract syntax tree (AST). The emitter reads the AST and generates C code.

![img.png](documentation/program_idea_diagram.png)

## Technologies
- Go 1.22.5

## Installation
1. Make sure you have the Go language installed on your system.
2. Clone the repository.
```shell
$ git clone git@github.com:pploszczyca/GoTinyBasicCompiler.git
```

## Usage
The project has a Makefile with the following commands:
- `make build` - Build the project
- `make run ARGS="path_to_bas_file path_to_c_file"` - Run the project with the path to the BAS file to compile and the output C file as arguments
- `make clean` - Clean the project
- `make test` - Run tests
- `make buildSamples` - Build samples from the `samples` directory
- `make buildAndRunSamples` - Build and run samples from the `samples` directory
- `make format` - Format the code

## Features
- Parsing TinyBASIC code into C code
- Setting input and output files
- Support for parsing expressions
- Support for the following TinyBASIC commands:
  - PRINT
  - INPUT
  - LET
  - GOTO
  - IF ... THEN
  - END
  - FOR ... TO ... NEXT
  - WHILE ... WEND
  - GOSUB
  - RETURN

## Examples
The `samples` directory contains examples of TinyBASIC code, and the compiled versions can be found in the `results` directory.
