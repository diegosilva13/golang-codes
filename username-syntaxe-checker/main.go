package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
)

const usernameRegex = `^@?(\w){1,15}$`

func main() {
	var usernameInput string

	flag.StringVar(&usernameInput, "username", "Gopher", "The GopherFace Username To Check")
	flag.Parse()

	fmt.Println("GopherFace Validation Checker")
	fmt.Println()
	fmt.Println("Checking Syntax for username, \"", usernameInput, "\", resulted in:\"", checkUsernameSyntax(usernameInput))
}

func checkUsernameSyntax(username string) bool {
	validationResult := false
	r, err := regexp.Compile(usernameRegex)
	if err != nil {
		log.Fatal(err)
	}
	validationResult = r.MatchString(username)
	return validationResult
}
