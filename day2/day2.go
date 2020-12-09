package main

import (
	basic "basic/mod"
	"fmt"
	"strconv"
	"strings"
)

var validpass int = 0
var errorpass int = 0

func part1(file string) int {
	input := basic.GetInput(file)
	for i := 0; i < len(input); i++ {
		s := strings.Split(input[i], " ")
		nr := strings.Split(s[0], "-")
		var test int = strings.Count(s[2], strings.Trim(s[1], ":"))
		min, _ := strconv.Atoi(nr[0])
		max, _ := strconv.Atoi(nr[1])
		if (min <= test) && test <= max {
			validpass++
		} else {
			errorpass++
		}
	}
	return validpass
}

func part2(file string) int {
	input := basic.GetInput(file)
	for i := 0; i < len(input); i++ {
		s := strings.Split(input[i], " ")
		nr := strings.Split(s[0], "-")
		letter := strings.Trim(s[1], ":")
		st := strings.Split(s[2], "")
		min, _ := strconv.Atoi(nr[0])
		max, _ := strconv.Atoi(nr[1])
		var j []int
		var m int = 0
		for i := 0; i < len(st); i++ {
			if st[i] == letter {
				j = append(j, i+1)
			}
		}
		for i := 0; i < len(j); i++ {
			if min == j[i] {
				m++
			}
			if max == j[i] {
				m++
			}
		}
		if m == 1 {
			validpass++
		} else {
			errorpass++
		}
	}

	return validpass
}

func main() {
	validpass := part2("input")
	fmt.Println(validpass)
}
