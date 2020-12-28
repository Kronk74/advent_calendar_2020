package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func getID(line string) int {
	id := 0
	row := 0
	columns := 0
	min := 0
	max := 128
	st := strings.Split(line, "")
	for i := 0; i < 6; i++ {
		if st[i] == "B" {
			min = min + (max-min)/2
		} else if st[i] == "F" {
			max = min + (max-min)/2
		}
	}
	max = max - 1
	if st[6] == "B" {
		row = max
	} else if st[6] == "F" {
		row = min
	}
	min = 0
	max = 8
	for i := 7; i < 9; i++ {
		if st[i] == "R" {
			min = min + (max-min)/2
		} else if st[i] == "L" {
			max = min + (max-min)/2
		}
	}
	max = max - 1
	if st[9] == "R" {
		columns = max
	} else if st[9] == "L" {
		columns = min
	}
	id = (row * 8) + columns

	return id
}

func part1(input string) int {
	id := 0
	result := 0
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	//for each line
	for scanner.Scan() {
		id = getID(scanner.Text())
		if id > result {
			result = id
		}
	}
	return result
}

func part2(input string) int {
	id := 0
	result := 0
	var ids []int
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	//for each line
	for scanner.Scan() {
		id = getID(scanner.Text())
		ids = append(ids, id)
	}
	sort.Ints(ids)
	//fmt.Println(ids)
	seat := ids[0]
	for i := 0; i < len(ids); i++ {
		if seat != ids[i] && (seat-1) == ids[i-1] && (seat+1) == ids[i] {
			result = seat
			seat++
		}
		seat++
	}
	return result
}

func main() {
	validpass := part2("input")
	fmt.Println(validpass)
}
