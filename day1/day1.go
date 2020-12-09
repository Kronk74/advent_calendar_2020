package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var result []int

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		result = append(result, i)
		if err != nil {
			log.Fatal(err)
		}
	}
	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result); j++ {
			for k := 0; k < len(result); k++ {
				res := result[i] + result[j] + +result[k]
				if res == 2020 {
					fmt.Printf("%v \n", result[i]*result[j]*result[k])
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
