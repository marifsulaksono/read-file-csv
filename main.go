package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"read-csv/config"
	"read-csv/helper"
)

func main() {
	ctx := context.Background()
	file, err := os.Open("organizers.csv")
	if err != nil {
		log.Fatalln("Error when open file csv :", err)
	}
	defer file.Close()

	db, err := config.ConnectToDatabase()
	if err != nil {
		log.Fatalf("Connection failed : %v", err)
	}

	records, err := helper.ReadAllRecordsWithoutHeader(file)
	if err != nil {
		log.Fatalln("Error when reading records :", err)
	}

	for i, eachRecord := range records {
		query := "INSERT INTO organizers VALUES(?,?,?,?,?,?)"
		stmt, err := db.PrepareContext(ctx, query)
		if err != nil {
			log.Printf("Error prepare statement : %v", err)
		}

		_, err = stmt.ExecContext(ctx, nil, eachRecord[1], eachRecord[2], eachRecord[3], eachRecord[4], eachRecord[5])
		if err != nil {
			log.Printf("Error execute statement : %v", err)
		}

		fmt.Printf("Success insert data - %d\n", i)
	}
}
