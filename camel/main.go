package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	inputFile := flag.String("input", "camel.in", "Input file with Camel case string")
	flag.Parse()

	input, err := os.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}

	answer := 1

	for _, char := range input {
		str := string(char)

		if strings.ToUpper(str) == str {
			answer++
		}
	}
	fmt.Println(answer)
}
