package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	// whitespace formatting
	addLineBreak(1)

	fmt.Println("It's quiz time~!")

	addLineBreak(2)

	// open csv file and handle potential errors
	csvFile, err := os.Open("../../problems.csv")

	if err != nil {
		log.Fatal(err)
	}

	// instantiate a new csv reader
	csvReader := csv.NewReader(csvFile)

	// read all lines in csv and handle potential errors
	line, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	// set inputVal with type string to handle cmd line input
	var inputVal string

	traverseCSVLines(inputVal, line)
}

// whitespace formatting
func addLineBreak(n int) {
	for i := 0; i <= n; i++ {
		fmt.Println()
	}
}

// loop through each line and get input / compare to answer
func traverseCSVLines(inputVal string, line [][]string) {
	// counter
	var questionsAnsweredCorrectly int

	for i := range line {

		start := time.Now().Unix()

		fmt.Println("What is", line[i][0], "?")

		addLineBreak(1)

		fmt.Scanf("%s", &inputVal)

		addLineBreak(1)

		var greeting string

		if inputVal == line[i][1] {
			greeting = ":) Congrats!"
			questionsAnsweredCorrectly++
		} else {
			greeting = ":( Aww.."
		}

		addLineBreak(5)

		fmt.Println(greeting, "You entered:", inputVal, "and the answer is", line[i][1])
		fmt.Println("It took you", time.Now().Unix()-start, "seconds")

		addLineBreak(1)
	}
}
