package search

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/Boo-Geonhyeok/playground/models"
)

func Search(input string) {
	file, err := os.OpenFile("search.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	defer file.Close()

	if err != nil {
		log.Fatalln("failed to open file", err)
	}

	w := csv.NewWriter(file)
	defer w.Flush()

	row := []string{input}
	if err := w.Write(row); err != nil {
		log.Fatalln("error writing record to file", err)
	}

	TFIDFs := getTFIDF()
	TFIDFindex := models.SortAndReturnIdxs(TFIDFs)
	fmt.Println(TFIDFs)
	firstZero := 0
	for index, tfidf := range TFIDFs {
		s := fmt.Sprintf("%f", tfidf)
		if s == "0.000000" {
			firstZero = index
			break
		}
	}

	bookFile, err := os.Open("./book.csv")
	if err != nil {
		log.Fatal("Can't Open File" + err.Error())
	}
	defer file.Close()

	reader := csv.NewReader(bookFile)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Can't Read File", err.Error())
	}

	for i := 0; i < firstZero; i++ {
		index := TFIDFindex[i]
		fmt.Println(records[index][0], records[index][1])
	}
}
