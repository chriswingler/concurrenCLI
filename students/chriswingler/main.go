package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
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

	// set val with type string to handle cmd line input
	var val string

	// loop through each line and get input / compare to answer
	for i := range line {
		fmt.Println("What is", line[i][0], "?")

		addLineBreak(1)

		fmt.Scanf("%s", &val)

		addLineBreak(1)

		var greeting string

		if val == line[i][1] {
			greeting = ":) Congrats!"
		} else {
			greeting = ":( Aww.."
		}

		addLineBreak(5)

		fmt.Println(greeting, "You entered:", val, "and the answer is", line[i][1])

		addLineBreak(1)

	}
}

func addLineBreak(n int) {
	for i := 0; i <= n; i++ {
		fmt.Println()
	}
}
