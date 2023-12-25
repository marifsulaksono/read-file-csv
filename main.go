package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func readRecords(file io.ReadSeeker) ([][]string, error) {
	firstRow, err := bufio.NewReader(file).ReadSlice('\n')
	if err != nil {
		return nil, err
	}

	_, err = file.Seek(int64(len(firstRow)), io.SeekStart)
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func main() {
	file, err := os.Open("employee.csv")
	if err != nil {
		log.Fatalln("Error when open file csv :", err)
	}

	defer file.Close()

	records, err := readRecords(file)
	if err != nil {
		log.Fatalln("Error when reading records :", err)
	}

	for _, eachRecord := range records {
		fmt.Printf("name : %s \t role : %s\n", eachRecord[1], eachRecord[2])
	}
}
