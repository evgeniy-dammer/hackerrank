package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func rotate(r rune, base, delta int) rune {
	temp := int(r) - base

	temp = (temp + delta) % 26

	return rune(temp + base)
}

func cipher(r rune, delta int) rune {
	if r >= 'A' && r <= 'Z' {
		return rotate(r, 'A', delta)
	}

	if r >= 'a' && r <= 'z' {
		return rotate(r, 'a', delta)
	}

	return r
}

func main() {
	inputFile := flag.String("input", "caesar.in", "Input file with Caesar Cipher string")
	flag.Parse()

	f, err := os.Open(*inputFile)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var input []string
	var result []rune

	for scanner.Scan() {
		input = append(input, string(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	str := input[0]
	delta, _ := strconv.Atoi(input[1])

	for _, char := range str {
		result = append(result, cipher(char, delta))
	}

	fmt.Println(string(result))
}
