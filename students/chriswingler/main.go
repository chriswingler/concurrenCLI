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

	reader := csv.NewReader(csvFile)

	stuff, err := reader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(stuff)
}
