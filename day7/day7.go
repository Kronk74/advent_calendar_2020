package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type bag struct {
	name    string
	contain map[string]int
}

func getListBags(input string) []bag {
	var bags []bag
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bags = append(bags, getBagInfos(scanner.Text()))
	}
	return bags
}

func getBagInfos(line string) bag {
	var bag bag
	re := regexp.MustCompile(`(?P<bag>\w+ \w+ \w+) (?:contain) (?P<contain>.*)(?:\.)`)
	match := re.FindStringSubmatch(line)
	bag.name = match[1]
	canContain := make(map[string]int)
	capacity := strings.Split(match[2], ",")
	for c := 0; c < len(capacity); c++ {
		rey := regexp.MustCompile(`(?P<nb>\d+)\s(?P<name>.\w+ \w+ \w+)`)
		if rey.Match([]byte(capacity[0])) {
			nbrBags, _ := strconv.Atoi(rey.FindStringSubmatch(capacity[c])[1])
			nameBags := rey.FindStringSubmatch(capacity[c])[2]
			canContain[nameBags] = nbrBags
		}
	}
	bag.contain = canContain
	return bag
}

func part1(input string) int {
	result := 0
	var bags []bag

	bags = getListBags(input)

	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		testBag := getBagInfos(scanner.Text())

		//for name, count := range testBag.contain {
		//}
		if testBag.contain["shiny gold bags"] > 0 || testBag.contain["shiny gold bag"] > 0 {
			fmt.Println(testBag.name)
			bags = append(bags, testBag)
			result++
		}

		//for k, v := range bags {
		//	fmt.Println(k)
		//	fmt.Println(v)
		//}
	}

	file.Close()
	return result
}

func main() {
	validpass := part1("input")
	fmt.Println(validpass)
}
