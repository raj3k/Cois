package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter csv file path: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("An error occured while reading input. Please try again.")
	}

	// remove delimiter from the string
	input = strings.TrimSuffix(input, "\n")

	records, err := readData(input)

	if err != nil {
		log.Fatal(err)
	}

	var cost int

	for _, record := range records {
		if contains(record, "COIS") {
			re := regexp.MustCompile("[0-9]+")
			c, err := strconv.Atoi(re.FindAllString(record[81], -1)[0])
			if err != nil {
				log.Fatal(err)
			}
			cost += c
		}
	}

	fmt.Printf("W Coisie wydałem %d złoty.", cost)
}

func readData(filename string) ([][]string, error) {
	f, err := os.Open(filename)

	if err != nil {
		return [][]string{}, err
	}

	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = ';'
	r.Comment = '#'

	// skip first line
	if _, err := r.Read(); err != nil {
		return [][]string{}, err
	}

	records, err := r.ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}

func contains(s []string, x string) bool {
	for _, a := range s {
		if strings.Contains(a, x) {
			return true
		}
	}
	return false
}
