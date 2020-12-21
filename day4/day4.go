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

func hardTestPassport(res map[string]string) bool {
	var valid bool
	var validation bool = true
	requiredField := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for _, s := range requiredField {
		valid = false
		switch field := s; field {
		case "byr":
			re := regexp.MustCompile("^(19[2-9][0-9]|200[0-2])$")
			valid = re.MatchString(res[s])
		case "iyr":
			re := regexp.MustCompile("^20(1[0-9]|20)$")
			valid = re.MatchString(res[s])
		case "eyr":
			re := regexp.MustCompile("^20(2[0-9]|30)$")
			valid = re.MatchString(res[s])
		case "hgt":
			re := regexp.MustCompile("^(\\d+)(cm|in)$")
			if re.MatchString(res[s]) {
				h := re.FindAllStringSubmatch(res[s], -1)
				for _, kv := range h {
					switch metric := kv[2]; metric {
					case "cm":
						rem := regexp.MustCompile("^1([5-8][0-9]|9[0-3])")
						valid = rem.MatchString(res[s])
					case "in":
						rem := regexp.MustCompile("^(59|6[0-9]|7[0-6])")
						valid = rem.MatchString(res[s])
					}
				}
			}
		case "hcl":
			re := regexp.MustCompile("^#[0-9a-f]{6}$")
			valid = re.MatchString(res[s])
		case "ecl":
			re := regexp.MustCompile("^(amb|blu|brn|gry|grn|hzl|oth)$")
			valid = re.MatchString(res[s])
		case "pid":
			re := regexp.MustCompile("^\\d{9}$")
			valid = re.MatchString(res[s])
		default:
			valid = false
		}
		if !valid {
			validation = false
		}

	}
	return validation
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

func part2(input string) int {
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
			if hardTestPassport(res) {
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
	if hardTestPassport(res) {
		validpass++
	}
	return validpass
}

func main() {
	validpass := part2("input")
	fmt.Println(validpass)
}
