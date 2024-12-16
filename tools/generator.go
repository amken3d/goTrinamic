package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Struct to store parsed field data
type Field struct {
	Name  string
	Value string
}

// Struct to store register field information during parsing
type RegisterFieldRaw struct {
	Name     string
	Mask     string
	Shift    string
	Register string
	IsSigned bool
}

// Struct to store register field information for code generation
type RegisterField struct {
	Name     string
	Mask     uint32
	Shift    uint32
	Register uint32
	IsSigned bool
}

// Regex patterns for parsing
var (
	defineRegex = regexp.MustCompile(`(?m)^#define\s+(\w+)\s+(.+)$`)
	//registerFieldRegex = regexp.MustCompile(`(?m)#define\s+(\w+)\s+\(\(\s*RegisterField\s*\)\s*\{\s*(\w+),\s*(\w+),\s*(\w+),\s*(true|false)\s*\}\)`)
	registerFieldRegex = regexp.MustCompile(`(?m)#define\s+(\w+)\(motor\s*\)\s*\(\(RegisterField\)\s*\{\s*(\w+),\s*(\w+),\s*(\w+)\(motor\),\s*(true|false)\s*\}\)`)
	hFilterRegex       = regexp.MustCompile(`_H_`)
)

// Parse C header file
func parseHeaderFile(filename string) ([]Field, []RegisterFieldRaw, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	var fields []Field
	var registerFields []RegisterFieldRaw

	content := strings.Builder{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Skip any line containing "_H_"
		if hFilterRegex.MatchString(line) {
			continue
		}
		content.WriteString(line + "\n")
	}

	lines := content.String()

	// Match #define constants
	defineMatches := defineRegex.FindAllStringSubmatch(lines, -1)
	for _, matches := range defineMatches {
		fields = append(fields, Field{
			Name:  matches[1],
			Value: matches[2],
		})
	}

	// Match RegisterField definitions
	fieldMatches := registerFieldRegex.FindAllStringSubmatch(lines, -1)
	for _, matches := range fieldMatches {
		registerFields = append(registerFields, RegisterFieldRaw{
			Name:     matches[1],
			Mask:     matches[2],
			Shift:    matches[3],
			Register: matches[4],
			IsSigned: matches[5] == "true",
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("failed to read file: %v", err)
	}

	// Debug: Log parsed results
	fmt.Printf("Extracted Constants: %d\n", len(fields))
	for _, f := range fields {
		fmt.Printf("Constant: %s = %s\n", f.Name, f.Value)
	}

	fmt.Printf("Extracted Register Fields: %d\n", len(registerFields))
	for _, rf := range registerFields {
		fmt.Printf("RegisterField: Name: %s, Mask: %s, Shift: %s, Register: %s, IsSigned: %v\n",
			rf.Name, rf.Mask, rf.Shift, rf.Register, rf.IsSigned)
	}

	return fields, registerFields, nil
}

// Generate Go code
func generateGoCode(fields []Field, registerFields []RegisterFieldRaw) string {
	var sb strings.Builder

	sb.WriteString("// Generated Go code\n\npackage main\n\n")

	// Filter out register fields from constants
	sb.WriteString("// Constants\nconst (\n")
	for _, field := range fields {
		// Skip if this constant is part of a RegisterField
		isRegisterField := false
		for _, regField := range registerFields {
			if field.Name == regField.Name {
				isRegisterField = true
				break
			}
		}
		if !isRegisterField {
			sb.WriteString(fmt.Sprintf("\t%s = %s\n", field.Name, field.Value))
		}
	}
	sb.WriteString(")\n\n")

	// Generate RegisterField structs
	sb.WriteString("// Register Fields\nvar (\n")
	for _, regField := range registerFields {
		sb.WriteString(fmt.Sprintf("\t%s = RegisterField{\n", regField.Name))
		sb.WriteString(fmt.Sprintf("\t\tMask: %s,\n", regField.Mask))
		sb.WriteString(fmt.Sprintf("\t\tShift: %s,\n", regField.Shift))
		sb.WriteString(fmt.Sprintf("\t\tRegister: %s,\n", regField.Register))
		sb.WriteString(fmt.Sprintf("\t\tIsSigned: %v,\n", regField.IsSigned))
		sb.WriteString("\t}\n")
	}
	sb.WriteString(")\n")

	return sb.String()
}

// Main function
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run generator.go <header-file>")
		return
	}

	filename := os.Args[1]
	fields, registerFields, err := parseHeaderFile(filename)
	if err != nil {
		fmt.Printf("Error parsing header file: %v\n", err)
		return
	}

	if len(fields) == 0 && len(registerFields) == 0 {
		fmt.Println("No constants or register fields found. Ensure the input file matches the expected format.")
		return
	}

	// Generate Go code
	goCode := generateGoCode(fields, registerFields)

	outputFilename := "generated.go"
	err = os.WriteFile(outputFilename, []byte(goCode), 0644)
	if err != nil {
		fmt.Printf("Error writing Go file: %v\n", err)
		return
	}

	fmt.Printf("Generated Go code saved to %s\n", outputFilename)
}
