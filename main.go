package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) == 1 || os.Args[1] == "--help" {
		fmt.Println("cregex - commandline Regex processor [version 1.0]")
		fmt.Println(" Usage: cregex [string] [expression]")
		fmt.Println("\ncregex is a tool for processing Regex expressions, in the commandline")
		fmt.Println("\nExample: cregex \"Hello World!\" \"[A-z]\"")
		return
	} else if len(os.Args) < 3 {
		fmt.Println("cregex: Not enough arguments")
		fmt.Println(" Usage: cregex [string] [expression]")
		return
	}

	fmt.Println(regexp.MatchString(os.Args[2], os.Args[1]))
}
