package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Read argumants from command line and check the number of argumants
	Args := os.Args
	if len(Args) < 2 || len(Args) > 4 {
		fmt.Println("please enter a valid arguments <program Name> <Input test> ")
		fmt.Println("Or:")
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		fmt.Println("EX: go run . --output=<fileName.txt> something standard")
		fmt.Println("Or:")
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		fmt.Println("EX: go run . something standard")
		return
	}
	// get the argumant and check if it is valid
	var argsValue string
	if len(Args) == 2 || len(Args) == 3 {
		argsValue = os.Args[1]
	} else {
		argsValue = os.Args[2]
	}

	for _, v := range argsValue {
		if v < 32 || v > 126 {
			fmt.Println("Please enter a valid input!!")
			return
		}
	}
	// count how much of new lines
	count := strings.Count(argsValue, "\\n")
	// split the argumant by new lines
	testLines := strings.Split(argsValue, "\\n")
	if argsValue == "" {
		// Empty input string, print nothing
		return
	}
	// get the content of banner
	var bannerType string
	if len(Args) > 2 {
		bannerType = os.Args[2]
	}
	if len(Args) == 4 {
		bannerType = os.Args[3]
	}

	var banner []byte
	var asciiChars []string
	if bannerType == "standard" || len(Args) == 2 {
		content, err := os.ReadFile("standard.txt")
		banner = content
		if err != nil {
			fmt.Println(err)
			return
		}
	} else if bannerType == "shadow" {
		content, err := os.ReadFile("shadow.txt")
		banner = content
		if err != nil {
			fmt.Println(err)
			return
		}
	} else if bannerType == "thinkertoy" {
		content, err := os.ReadFile("thinkertoy.txt")
		banner = content
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		fmt.Println("EX: go run . something standard")
		return
	}
	// fmt.Println(banner)
	phrase := string(banner)
	// split the banner by double new line
	if bannerType == "thinkertoy" {
		asciiChars = strings.Split(phrase, "\r\n\r\n")
	} else {
		asciiChars = strings.Split(phrase, "\n\n")
	}
	characters := make([][]string, len(asciiChars))
	// stock every line of instance of banner in a case
	for i, char := range asciiChars {
		if bannerType == "thinkertoy" {
			characters[i] = strings.Split(char, "\r\n")
		} else {
			characters[i] = strings.Split(char, "\n")
		}
	}
	// print a result with all requirements
	result := ""
	counter := 1
	for _, line := range testLines {
		if line == "" {
			if counter <= count {
				result += "\n"
			}
			counter++
			continue
		}
		for l := 0; l < 8; l++ {
			for _, char := range line {
				if char == ' ' { // a space character
					result += characters[char-32][l+1]
				} else {
					index := char - 32

					result += characters[index][l]
				}
			}
			result += "\n"
		}
	}
	if len(Args) != 4 {
		fmt.Print(result)
	}

	if len(Args) == 4 {
		flag := "--output="
		if os.Args[1][:len(flag)] == flag && len(os.Args[1][len(flag):]) > 0 && os.Args[1][len(os.Args[1])-4:] == ".txt" {
			err := os.WriteFile(os.Args[1][len(flag):], []byte(result), 0o644)
			if err != nil {
				fmt.Printf("Error writing output file: %v\n", err)
				return
			}
		} else {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
			fmt.Println("EX: go run . --output=<fileName.txt> something standard")
			return
		}
	}
}
