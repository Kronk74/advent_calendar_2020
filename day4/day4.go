package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func testPassport(requiredField []string, res map[string]string) bool {
	valid := true
	for _, s := range requiredField {
		if _, ok := res[s]; !ok {
			valid = false
		}
	}
	return valid
}

func part1(input string) int {

	requiredField := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	var rex = regexp.MustCompile("(\\w+):(#?\\w+)")
	res := make(map[string]string)
	var validpass int = 0

	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	//for each line
	for scanner.Scan() {
		data := rex.FindAllStringSubmatch(scanner.Text(), -1)
		for _, kv := range data {
			res[kv[1]] = kv[2]
		}
		//For each passport
		if scanner.Text() == "" {
			if testPassport(requiredField, res) {
				validpass++
			}

			for k := range res {
				delete(res, k)
			}
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	if testPassport(requiredField, res) {
		validpass++
	}
	return validpass
}

func main() {
	validpass := part1("input")
	fmt.Println(validpass)
}
