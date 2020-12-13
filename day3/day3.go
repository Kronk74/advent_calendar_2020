package main

import (
	basic "basic/mod"
	"fmt"
	"strings"
)

func fillArray(input []string, mappy *[][]string) {
	for i := 0; i < len(input); i++ {
		line := strings.Split(input[i], "")
		*mappy = append(*mappy, line)
	}
	//return mappy
}

func getTree(mappy [][]string, right int, down int) int {

	nbr := 0
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

func part1(file string) int {
	nbr := 0
	right := 3
	down := 1
	input := basic.GetInput(file)
	mappy := [][]string{}

	fillArray(input, &mappy)

	nbr = getTree(mappy, right, down)
	return nbr
}

func part2(file string) int {
	nbr := 1
	right := 1
	down := 1
	input := basic.GetInput(file)
	mappy := [][]string{}

	//Fill 2D array
	fillArray(input, &mappy)

	//Find trees
	for right := 1; right <= 7; right = right + 2 {
		trees := getTree(mappy, right, down)
		nbr = trees * nbr
		fmt.Printf("Trees : %v, Right: %v, Down: %v\n", trees, right, down)
		fmt.Println(nbr)
	}
	down = 2
	trees := getTree(mappy, right, down)
	nbr = trees * nbr
	fmt.Printf("Trees : %v, Right: %v, Down: %v\n", trees, right, down)
	return nbr
}

func main() {
	validpass := part2("input")
	fmt.Println(validpass)
}
