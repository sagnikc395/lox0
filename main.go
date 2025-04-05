package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Lox struct {
	hadError bool
}

func NewLox(hadError bool) *Lox {
	return &Lox{
		hadError: hadError,
	}
}

func main() {
	args := os.Args[1:]

	if len(args) > 1 {
		fmt.Println("Usage: lox0 [script]")
		os.Exit(64)
	} else if len(args) == 1 {
		runFile(args[0])
	} else {
		runPrompt()
	}
}

// since lox is a scripting language meaning that we can directly
// execute it from source
// the interpreter supports 2 ways of running
// if we start lox0 from command line and give it a path to a file, it
// will read the file and execute it.

func (l *Lox) runFile(path string) error {
	contents, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
		return err
	}
	run(string(contents))
	if l.hadError {
		os.Exit(65)
	}
	return nil
}

//we can also run it interactively,
// fire up our interpreter and it will drop us into a prompt where we can
// enter and execute code one line at a time

func (l *Lox) runPrompt() error {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Printf("> ")
		line := scanner.Text()
		if line == "" {
			break
		}
		run(line)
		l.hadError = false
	}
	return nil
}

//runPrompt and runFile are just wrappers around this core function

func run(source string) {
	// TODO
	// ?
	// 	Scanner scanner = new Scanner(source);
	//  List<Token> tokens = scanner.scanTokens();
	//  // For now, just print the tokens.
	//  for (Token token : tokens) {
	//  System.out.println(token);
	//  }
}
