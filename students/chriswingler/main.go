package main

import (
	"encoding/csv"
	"flag"
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

	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	// set inputVal with type string to handle cmd line input
	var inputVal string

	traverseCSVLines(inputVal, line, timer)
}

// whitespace formatting
func addLineBreak(n int) {
	for i := 0; i <= n; i++ {
		fmt.Println()
	}
}

// loop through each line and get input / compare to answer
func traverseCSVLines(inputVal string, line [][]string, timer *time.Timer) {
	// counter
	var questionsAnsweredCorrectly int

	start := time.Now().Unix()

	for i := range line {
		select {
		case <-timer.C:
			fmt.Println("Out of time")

			addLineBreak(1)

			fmt.Println("You answered", questionsAnsweredCorrectly, "out of", len(line), "correctly.")

			addLineBreak(1)

			os.Exit(0)
		default:

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
				timer.Reset(time.Duration(time.Second))
			}

			addLineBreak(5)

			fmt.Println(greeting, "You entered:", inputVal, "and the answer is", line[i][1])

			addLineBreak(1)

			fmt.Println("You answered", questionsAnsweredCorrectly, "out of", len(line), "correctly.")

			addLineBreak(1)

			fmt.Println("It took you", time.Now().Unix()-start, "seconds to finish this quiz")

			addLineBreak(1)
		}
	}
}
