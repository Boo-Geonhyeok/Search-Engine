package search

import (
	"math"
	"strings"
)

func getTFIDF() []float64 {
	records := getBookRecords()
	frequencies := []map[string]int{}
	searchMorpheme := getSearchMorpheme()

	for _, record := range records {
		rawString := strings.ReplaceAll(record[3][1:len(record[3])-1], "'", "")
		nouns := strings.Split(rawString, ",")
		frequency := getWordFrequency(nouns)
		frequencies = append(frequencies, frequency)
	}
	TFs := calculateWordFrequency(frequencies, searchMorpheme)
	IDFs := getIDF(frequencies, searchMorpheme)

	TFIDFs := []float64{}
	wordNum := len(IDFs)
	for _, tf := range TFs {
		tfidf := float64(0)
		for i := 0; i < wordNum; i++ {
			tfidf += float64(tf[i]) * IDFs[i]
		}
		TFIDFs = append(TFIDFs, tfidf)
	}

	return TFIDFs
}

func getWordFrequency(record []string) map[string]int {
	frequency := make(map[string]int, 0)
	for _, word := range record {
		frequency[word] = frequency[word] + 1
	}
	return frequency
}

func calculateWordFrequency(wordFrequency []map[string]int, searchMorpheme []string) [][]float32 {
	AllTFs := [][]float32{}
	for _, wf := range wordFrequency {
		totalNum := 0
		for _, val := range wf {
			totalNum += val
		}
		a := wf
		TFs := []float32{}
		for _, word := range searchMorpheme {
			val, ok := a[word]
			if ok {
				TFs = append(TFs, (float32(val) / float32(totalNum)))
			} else {
				TFs = append(TFs, 0)
			}
		}
		AllTFs = append(AllTFs, TFs)
	}
	return AllTFs
}

func getIDF(frequencies []map[string]int, searchMorpheme []string) []float64 {
	linkNum := len(frequencies)
	IDFs := []float64{}
	for _, word := range searchMorpheme {
		wordExistPage := 0
		for _, frequency := range frequencies {
			if _, ok := frequency[word]; ok {
				wordExistPage += 1
			}
		}
		if wordExistPage == 0 {
			wordExistPage = linkNum
		}
		idf := math.Log2(float64(linkNum) / float64(wordExistPage))
		IDFs = append(IDFs, idf)
	}
	return IDFs
}
