package main

import (
	"fmt"
	"strings"
)

// Regular expressions
//var (
//	defineRegex = regexp.MustCompile(`(?m)^#define\s+(\w+)\s+(.+)$`)
//	maskRegex   = regexp.MustCompile(`(?m)#define\s+(\w+)_MASK\s+(0x[0-9A-Fa-f]+)`)
//	shiftRegex  = regexp.MustCompile(`(?m)#define\s+(\w+)_SHIFT\s+(\d+)`)
//	//fieldRegex         = regexp.MustCompile(`(?m)#define\s+(\w+)_FIELD\s+\(\(RegisterField\)\s+\{\s*(\w+),\s*(0x[0-9A-Fa-f]+),\s*(\d+),\s*(true|false)\s*}\)`)
//	cTypeRegex   = regexp.MustCompile(`\((uint\d+_t)\)`)
//	hFilterRegex = regexp.MustCompile(`_H_`)
//	//registerFieldRegex = regexp.MustCompile(`(?m)#define\s+(\w+)_FIELD\s+\(\(RegisterField\)\s+\{\s*(\w+),\s*(\d+),\s*(\w+),\s*(true|false)\s*}\)`)
//	fieldRegex         = regexp.MustCompile(`(?m)#define\s+(\w+)_FIELD\s+\(\(\s*RegisterField\s*\)\s*\{\s*(\w+)\,\s*(\w+),\s*(\w+),\s*(true|false)\s*}`)
//	registerFieldRegex = regexp.MustCompile(`(?m)#define\s+(\w+)_FIELD\s+\(\(\s*RegisterField\s*\)\s*\{\s*(\w+)\,\s*(\w+),\s*(\w+),\s*(true|false)\s*}`)
//)

// Sample C header content
const sampleHeader = `
#ifndef TMC2240_H_
#define TMC2240_H_

#define TMC2240_REGISTER_COUNT 128
##define TMC2240_FAST_STANDSTILL_MASK          0x00000002
#define TMC2240_FAST_STANDSTILL_SHIFT         1
#define TMC2240_FAST_STANDSTILL_FIELD         ((RegisterField) {TMC2240_FAST_STANDSTILL_MASK, TMC2240_FAST_STANDSTILL_SHIFT, TMC2240_GCONF, false})
#define TMC2240_EN_PWM_MODE_MASK              0x00000004
#define TMC2240_EN_PWM_MODE_SHIFT             2
#define TMC2240_EN_PWM_MODE_FIELD             ((RegisterField) {TMC2240_EN_PWM_MODE_MASK, TMC2240_EN_PWM_MODE_SHIFT, TMC2240_GCONF, false})
#define TMC2240_MULTISTEP_FILT_MASK           0x00000008
#define TMC2240_MULTISTEP_FILT_SHIFT          3
#define TMC2240_MULTISTEP_FILT_FIELD          ((RegisterField) {TMC2240_MULTISTEP_FILT_MASK, TMC2240_MULTISTEP_FILT_SHIFT, TMC2240_GCONF, false})
#define TMC2240_SHAFT_MASK                    0x00000010
#define TMC2240_SHAFT_SHIFT                   4
#define TMC2240_SHAFT_FIELD                   ((RegisterField) {TMC2240_SHAFT_MASK, TMC2240_SHAFT_SHIFT, TMC2240_GCONF, false})
#define TMC2240_DIAG0_ERROR_MASK              0x00000020
#define TMC2240_DIAG0_ERROR_SHIFT             5
#define TMC2240_DIAG0_ERROR_FIELD             ((RegisterField) {TMC2240_DIAG0_ERROR_MASK, TMC2240_DIAG0_ERROR_SHIFT, TMC2240_GCONF, false})

#endif // TMC2240_H_
`

func testRegex() {
	lines := strings.Split(sampleHeader, "\n")

	// Test each regex
	fmt.Println("Testing defineRegex...")
	for _, match := range defineRegex.FindAllStringSubmatch(sampleHeader, -1) {
		fmt.Printf("Match: %s = %s\n", match[1], match[2])
	}

	fmt.Println("\nTesting maskRegex...")
	for _, match := range maskRegex.FindAllStringSubmatch(sampleHeader, -1) {
		fmt.Printf("Mask Name: %s, Value: %s\n", match[1], match[2])
	}

	fmt.Println("\nTesting shiftRegex...")
	for _, match := range shiftRegex.FindAllStringSubmatch(sampleHeader, -1) {
		fmt.Printf("Shift Name: %s, Value: %s\n", match[1], match[2])
	}

	fmt.Println("\nTesting fieldRegex...")
	for _, match := range fieldRegex.FindAllStringSubmatch(sampleHeader, -1) {
		fmt.Printf("Field Name: %s, Mask: %s, Shift: %s, Register: %s, Signed: %s\n", match[1], match[2], match[3], match[4], match[5])
	}

	fmt.Println("\nTesting cTypeRegex...")
	for _, line := range lines {
		if cTypeRegex.MatchString(line) {
			fmt.Printf("C Type Match: %s\n", line)
		}
	}

	fmt.Println("\nTesting hFilterRegex...")
	for _, line := range lines {
		if hFilterRegex.MatchString(line) {
			fmt.Printf("H Filter Match: %s\n", line)
		}
	}

	fmt.Println("\nTesting registerFieldRegex...")
	for _, match := range registerFieldRegex.FindAllStringSubmatch(sampleHeader, -1) {
		fmt.Printf("Register Field Name: %s, Mask: %s, Shift: %s, Register: %s, Signed: %s\n", match[1], match[2], match[3], match[4], match[5])
	}
}

func main() {
	testRegex()
}
