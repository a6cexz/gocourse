package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("data.csv")
	if err != nil {
		log.Fatalln("Can not read the csv file", err)
	}

	reader := csv.NewReader(file)
	for {
		data, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		for _, record := range data {
			fmt.Printf("%s ", record)
		}
		fmt.Println()
	}
}
