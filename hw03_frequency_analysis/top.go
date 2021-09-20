package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type WordFrequency struct {
	Word  string
	Count int
}

func Top10(initStr string) (result []string) {
	stringSlice := strings.Fields(initStr)

	resultMap := make(map[string]int)
	for _, word := range stringSlice {
		resultMap[word]++
	}

	words := []WordFrequency{}
	for word, count := range resultMap {
		words = append(words, WordFrequency{word, count})
	}

	sort.Slice(words, func(i, j int) bool {
		return words[i].Count > words[j].Count || words[i].Count == words[j].Count && words[i].Word < words[j].Word
	})

	for key, word := range words {
		if key == 10 {
			break
		}
		result = append(result, word.Word)
	}

	return
}
