package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	csvFile, err := os.Open("../../problems.csv")

	if err != nil {
		log.Fatal(err)
	}

	csvReader := csv.NewReader(csvFile)

	line, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(line); i++ {
		fmt.Println(line[i][0])
		fmt.Println(line[i][1])
	}

	var val string

	for {
		fmt.Println("enter a value")

		fmt.Scanf("%s", &val)

		fmt.Println("you entered: ", val)
	}
}
