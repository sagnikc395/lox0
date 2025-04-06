//tool for generating a class definition , field declaration
//consutrutor and initializer
// it will have a description of each tree type - its name and fields
// and print out the Go code needed to define a class with that name and state

package tool

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: generate_ast <output_directory>")
		os.Exit(64)
	}

	outputDir := os.Args[1]
	err := defineAST(outputDir, "Expr", []string{
		"Binary		: Expr left, Token operator, Expr right",
		"Grouping	: Expr expression",
		"Literal	: interface{} value",
		"Unary		: Token operator, Expr right",
	})
	if err != nil {
		log.Fatal(err)
	}
}

func defineAST(outputDir, baseName string, types []string) error {
	path := filepath.Join(outputDir, baseName+".go")
	file, err := os.Create(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating file :%v\n", err)
		return err
	}

	defer file.Close()

	fmt.Fprintln(file, "package lox0")
	fmt.Fprintln(file)
	fmt.Fprintln(file, "type", baseName, "interface {")
	fmt.Fprintln(file, "}")

	for _, t := range types {
		parts := strings.SplitN(t, ":", 2)
		typeName := strings.TrimSpace(parts[0])
		fields := strings.TrimSpace(parts[1])
		fmt.Fprintf(file, "\ntype %s struct {\n", typeName)
		for _, field := range strings.Split(fields, ",") {
			field = strings.TrimSpace(field)
			fieldParts := strings.SplitN(field, " ", 2)
			fmt.Fprintf(file, "\t%s %s\n", strings.Title(fieldParts[1]), fieldParts[0])
		}
		fmt.Fprintln(file, "}")
	}

	return nil
}

func defineType(file *os.File, baseName, className, fields string) {
	fmt.Fprintf(file, "\ntype %s struct {\n", className)
	for _, field := range strings.Split(fields, ",") {
		field = strings.TrimSpace(field)
		fieldParts := strings.SplitN(field, " ", 2)
		fmt.Fprintln(file, "\t%s %s\n", strings.Title(fieldParts[1]), fieldParts[0])
	}

	fmt.Fprintln(file, "}")
}
