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
	searchMorphemes, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Can't Read File", err.Error())
	}

	morpheme := searchMorphemes[len(searchMorphemes)-1][1]

	rawString := strings.ReplaceAll(morpheme[1:len(morpheme)-1], "'", "")
	nouns := strings.Split(rawString, ",")

	for index, word := range nouns {
		if !strings.HasPrefix(word, " ") {
			nouns[index] = " " + word
		}
		if strings.HasSuffix(word, " ") {
			nouns[index] = word[:len(word)-1]
		}
	}
	return nouns
}
