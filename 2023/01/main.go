package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	re := regexp.MustCompile(`\d`)

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var subtotal int

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		results := re.FindAll([]byte(scanner.Text()), -1)

		//fmt.Printf("%q\n", results)

		if len(results) == 1 {
			sum := string(results[0]) + string(results[0])
			//fmt.Println(sum)
			i, _ := strconv.Atoi(sum)
			subtotal = subtotal + i
		}

		if len(results) == 2 {
			sum := string(results[0]) + string(results[1])
			//fmt.Println(sum)
			i, _ := strconv.Atoi(sum)
			subtotal = subtotal + i
		}

		if len(results) >= 3 {
			sum := string(results[0]) + string(results[len(results)-1])
			//fmt.Println(sum)
			i, _ := strconv.Atoi(sum)
			subtotal = subtotal + i
		}
	}

	fmt.Printf("The total sum of all calibration values is: %d\n", subtotal)
}
