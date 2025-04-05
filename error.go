package main

import "fmt"

// error handling is essential for production level compilers
// hanlding errors gracefully is vital

func LoxError(l *Lox, line int, message string) {
	report(line, "", message)
	l.hadError = true
}

func report(line int, where, message string) error {
	err := fmt.Errorf("[line %d] Error %s: %s", line, where, message)
	return err
}
