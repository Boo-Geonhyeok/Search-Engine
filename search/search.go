package search

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/Boo-Geonhyeok/playground/models"
)

func Search(input string) {

	addToCsv(input)
	// run searchMorpheme.py
	c := exec.Command("python3", "./python/searchMorpheme.py")
	if err := c.Run(); err != nil {
		log.Fatal("Can't run searchMorpheme.py ", err)
	}

	TFIDFs := getTFIDF()
	TFIDFindex := models.SortAndReturnIdxs(TFIDFs)
	firstZero := 0
	for index, tfidf := range TFIDFs {
		s := fmt.Sprintf("%f", tfidf)
		if s == "0.000000" {
			firstZero = index
			break
		}
	}

	getBooks(firstZero, TFIDFindex)
}

func addToCsv(input string) {
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
}

func getBooks(firstZero int, TFIDFindex []int) {
	bookFile, err := os.Open("./book.csv")
	if err != nil {
		log.Fatal("Can't Open File" + err.Error())
	}
	defer bookFile.Close()

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
