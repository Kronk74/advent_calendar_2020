package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func part1(input string) int {
	answers := make(map[string]bool)
	result := 0
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	//for each line
	for scanner.Scan() {
		a := strings.Split(scanner.Text(), "")
		for i := 0; i < len(a); i++ {
			answers[a[i]] = true
		}
		if scanner.Text() == "" {
			result = result + len(answers)
			for k := range answers {
				delete(answers, k)
			}
		}
	}
	result = result + len(answers)
	return result
}

func main() {
	validpass := part1("input")
	fmt.Println(validpass)
}
