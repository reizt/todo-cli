package utils

import "fmt"

func PrintRed(str string) {
	fmt.Printf("\x1b[31m%s\x1b[0m", str)
}

func PrintGreen(str string) {
	fmt.Printf("\x1b[32m%s\x1b[0m", str)
}

func PrintYellow(str string) {
	fmt.Printf("\x1b[33m%s\x1b[0m", str)
}
