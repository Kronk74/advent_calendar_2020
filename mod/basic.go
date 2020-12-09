package basic

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var result []string

//GetInput Store input variable into an array
func GetInput(input string) []string {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i := scanner.Text()
		result = append(result, i)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}

//ShowInput Print data from array
func ShowInput(input []string) {
	for i := 0; i < len(input); i++ {
		fmt.Println(input[i])
	}
}
