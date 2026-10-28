package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"
)

func count(wordList []string, stopWords []string, wordFreqs map[string]int) {
	if len(wordList) == 0 {
	} else {
			word := wordList[0];
			if (!slices.Contains(stopWords, word) && len(word) != 1) {
				if (wordFreqs[word] == 0) {
					wordFreqs[word] += 1
				} else {
					wordFreqs[word] += 1
				}
			}
			count(wordList[1:], stopWords, wordFreqs)
	}

}

func wf_print(wordFreq [][]any) {
	if len(wordFreq) == 0 {
		return
	} else {
		w, c := wordFreq[0][0], wordFreq[0][1]
		fmt.Printf("%s - %d\n", w, c)
		wf_print(wordFreq[1:])
	}
}

func main() {
	recursionLimit := 5000
	stopBytes, _ := os.ReadFile("../stop_words.txt"); stopWords := strings.Split(string(stopBytes), ",")
	bookBytes, _ := os.ReadFile(os.Args[1])
	wordList := regexp.MustCompile("[^a-zA-Z]+").Split(strings.ToLower(string(bookBytes)), -1)
	wordFreqs := make(map[string]int)
	pairedFreqs := [][]any{}

	for i := 0; i < len(wordList); i += recursionLimit {
		end := i + recursionLimit

		if end > len(wordList) {
			end = len(wordList)
		}
	count(wordList[i:end], stopWords, wordFreqs)
	}

	for key, value := range wordFreqs {
		pairedFreqs = append(pairedFreqs, []any{key, value})
	}

	for i := 0; i < len(pairedFreqs); i++ {
		for n := len(pairedFreqs) - 1; n > i; n-- {
			if (pairedFreqs[n][1].(int) > pairedFreqs[n-1][1].(int)) {
				pairedFreqs[n], pairedFreqs[n-1] = pairedFreqs[n-1], pairedFreqs[n]
			}
		}
	}

	wf_print(pairedFreqs[:25])
}