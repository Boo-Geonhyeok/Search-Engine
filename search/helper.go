package search

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
)

func getBookRecords() [][]string {
	file, err := os.Open("./bookMorpheme.csv")
	if err != nil {
		log.Fatal("Can't Open File" + err.Error())
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Can't Read File", err.Error())
	}

	return records
}

func getSearchMorpheme() []string {
	file, err := os.Open("./searchMorpheme.csv")
	if err != nil {
		log.Fatal("Can't Open File" + err.Error())
	}
	defer file.Close()

	reader := csv.NewReader(file)
	searchMorpheme, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Can't Read File", err.Error())
	}

	rawString := strings.ReplaceAll(searchMorpheme[0][1][1:len(searchMorpheme[0][1])-1], "'", "")
	nouns := strings.Split(rawString, ",")

	return nouns
}
