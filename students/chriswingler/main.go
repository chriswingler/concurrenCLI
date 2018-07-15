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

	// boolean variable that is set / tells us when question is finished
	finished := make(chan bool)

	go traverseCSVLines(inputVal, line, finished)

	msg := <-finished
	fmt.Println(msg)
}

// whitespace formatting
func addLineBreak(n int) {
	for i := 0; i <= n; i++ {
		fmt.Println()
	}
}

// loop through each line and get input / compare to answer
func traverseCSVLines(inputVal string, line [][]string, finished chan bool) {
	// counter
	var questionsAnsweredCorrectly int

	for i := range line {

		fmt.Println("What is", line[i][0], "?")

		time.Sleep(time.Second * 5)

		finished <- true

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

		addLineBreak(1)
	}
}
