package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("employee.csv")
	if err != nil {
		log.Fatalln("Error when open file csv :", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalln("Error when reading records :", err)
	}

	for _, eachRecord := range records {
		fmt.Printf("name : %s \t role : %s\n", eachRecord[1], eachRecord[2])
	}
}
