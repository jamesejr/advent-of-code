package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Match struct {
	Target string
	Index  int
}

// isNumber checks if the string is a number
func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

// getNumber converts a string to an int
func getNumber(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

// wordToNum converts a word to a number
func wordToNum(s string) string {
	numbers := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	return numbers[s]
}

// findMatches finds all valid digit matches in a string
func findMatches(s string) []Match {
	var matches []Match

	targets := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	for _, target := range targets {
		startIndex := 0
		for {
			index := strings.Index(s[startIndex:], target)
			if index == -1 {
				break
			}

			match := s[startIndex:][index : index+len(target)]
			matches = append(matches, Match{Target: match, Index: startIndex + index})
			startIndex += index + len(target)
		}
	}

	// sort matches by index
	sort.Slice(matches, func(i, j int) bool {
		return matches[i].Index < matches[j].Index
	})
	return matches
}

func main() {
	var subtotal int

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		matches := findMatches(scanner.Text())

		var results []string

		// convert all words to numbers
		for _, match := range matches {
			if isNumber(match.Target) {
				results = append(results, match.Target)
			} else {
				results = append(results, wordToNum(match.Target))
			}
		}

		if len(results) == 1 {
			sum := results[0] + results[0]
			subtotal += getNumber(sum)
		} else if len(results) >= 2 {
			sum := results[0] + results[len(results)-1]
			subtotal += getNumber(sum)
		}
	}

	fmt.Printf("The total sum of all calibration values is: %d\n", subtotal)
}
