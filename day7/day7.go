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

func unique(s []string) []string {
	unique := make(map[string]bool, len(s))
	us := make([]string, len(unique))
	for _, elem := range s {
		if len(elem) != 0 {
			if !unique[elem] {
				us = append(us, elem)
				unique[elem] = true
			}
		}
	}

	return us

}

func part1(input string) int {
	result := 0
	var bags []bag
	var baggy []string

	bags = getListBags(input)

	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		testBag := getBagInfos(scanner.Text())

		//Test is a bag contain one or more shiny gold bag
		if testBag.contain["shiny gold bag"] > 0 || testBag.contain["shiny gold bag"+"s"] > 0 {
			baggy = append(baggy, testBag.name)
		}
	}
	var tempBag1 []string
	var tempBag2 []string
	var tb2 string
	tempBag1 = append(tempBag1, baggy...)
	for len(tempBag1) > 0 {
		for _, b := range bags {
			for _, tb1 := range tempBag1 {
				retb := regexp.MustCompile(`s$`)
				if retb.Match([]byte(tb1)) {
					tb2 = tb1[:len(tb1)-1]
				} else {
					tb2 = tb1 + "s"
				}
				if b.contain[tb1] > 0 || b.contain[tb2] > 0 {
					tempBag2 = append(tempBag2, b.name)
				}

			}
		}
		tempBag1 = tempBag1[:0]
		tempBag1 = append(tempBag1, tempBag2...)
		baggy = append(baggy, tempBag2...)
		tempBag2 = tempBag2[:0]

	}

	baggy = unique(baggy)
	result = len(baggy)
	file.Close()
	return result
}

func main() {
	validpass := part1("input")
	fmt.Println(validpass)
}
