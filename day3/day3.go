package main

import (
	basic "basic/mod"
	"fmt"
	"strings"
)

func part1(file string) int {
	nbr := 0
	right := 3
	down := 1
	input := basic.GetInput(file)
	mappy := [][]string{}

	//Fill 2d array from input
	for i := 0; i < len(input); i++ {
		line := strings.Split(input[i], "")
		mappy = append(mappy, line)
	}

	j := right

	for i := down + 1; i < len(mappy); i = i + down {
		j = j + right
		if j >= len(mappy[i]) {
			j = j - len(mappy[i])
		}
		if mappy[i][j] == "#" {
			nbr++
		}
	}
	return nbr
}

func part2(file string) int {
	nbr := 0
	right := 3
	down := 1
	input := basic.GetInput(file)
	mappy := [][]string{}

	//Fill 2d array from input
	for i := 0; i < len(input); i++ {
		line := strings.Split(input[i], "")
		mappy = append(mappy, line)
	}

	j := right

	for i := down + 1; i < len(mappy); i = i + down {
		j = j + right
		if j >= len(mappy[i]) {
			j = j - len(mappy[i])
		}
		if mappy[i][j] == "#" {
			nbr++
		}
	}
	return nbr
}

func main() {
	validpass := part1("input")
	fmt.Println(validpass)
}
