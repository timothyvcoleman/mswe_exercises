// Exercise 1: Implement Term Frequency Program
// Timothy Coleman
// 10/05/2026
// MSWE 262P

package main

import (
	"os"
	"fmt"
	"strings"
	"strconv"

)

func readFiles() {

	word_freq := [][]any{}

	// var = make(map[KeyType]ValueType, *initialCapacity)

	// count := map[string]int <-- since maps are reference types, the intial val is nil

	// you must use make to initialize maps before writing, else you get a panic

	count := make(map[string]int)
	
	// os.ReadFile returns file contents as a byte slice and err == nil if successful
	// the _ is used to ignore the error value, thus foregoing error handling

	rawContentStopWords, _ := (os.ReadFile("../stop_words.txt"))

	// the string function converts bytes to a string

	stringContentStopWords := string(rawContentStopWords)

	// the Split() function segments the string into substrings, seperated by commas

	stopWords := strings.Split(stringContentStopWords, ",")

	// create new map for stopWords, by iterating through the word-values, 
	// take each word from the slice, and assign it to the map
	stopWordsMap := make(map[string]bool)
	for _, word := range stopWords {
		stopWordsMap[word] = true
	}

	rawContentBook, _ := (os.ReadFile(os.Args[1]))
	stringContentBook := string(rawContentBook)

	endOfHeader, _ := strconv.Atoi(os.Args[2])
	endOfBook, _ := strconv.Atoi(os.Args[3])

	// remove irrelevant lines
	lineByLine := strings.Split(stringContentBook, "\n")
	skipLines := lineByLine[endOfHeader:endOfBook]
	justBook := strings.Join(skipLines, " ")

	// convert all words to lowercase
	lowercaseBook := strings.ToLower(justBook)

	// remove punctuation
	removeCommas := strings.ReplaceAll(lowercaseBook, ",", "")
	removePeriods := strings.ReplaceAll(removeCommas, ".", "")
	removeQuestions := strings.ReplaceAll(removePeriods, "?", "")
	removeExclamations := strings.ReplaceAll(removeQuestions, "!", "")
	removeQuotes := strings.ReplaceAll(removeExclamations, "\"", "")
	removeApostrophes := strings.ReplaceAll(removeQuotes, "'", "")
	removeparens := strings.ReplaceAll(removeApostrophes, "(", "")
	removeparens2 := strings.ReplaceAll(removeparens, ")", "")
	removeColons := strings.ReplaceAll(removeparens2, ":", "")
	removeSemicolons := strings.ReplaceAll(removeColons, ";", "")
	removeUnderscores := strings.ReplaceAll(removeSemicolons, "_", "")
	removeDoubleHyphens := strings.ReplaceAll(removeUnderscores, "--", "")

	// remove numbers
	removeOne := strings.ReplaceAll(removeDoubleHyphens, "1", "")
	removeTwo := strings.ReplaceAll(removeOne, "2", "")
	removeThree := strings.ReplaceAll(removeTwo, "3", "")
	removeFour := strings.ReplaceAll(removeThree, "4", "")
	removeFive := strings.ReplaceAll(removeFour, "5", "")
	removeSix := strings.ReplaceAll(removeFive, "6", "")
	removeSeven := strings.ReplaceAll(removeSix, "7", "")
	removeEight := strings.ReplaceAll(removeSeven, "8", "")
	removeNine := strings.ReplaceAll(removeEight, "9", "")
	removeZero := strings.ReplaceAll(removeNine, "0", "")
	// seperate individual words
	bookWords := strings.Split(removeZero, " ")



	// for every word in the book, check if it is a stopWord
	// if not, go to the count map and increment the value for that word by 1
	// also checks for single letters
	for _, word := range bookWords {
		if !stopWordsMap[word] && len(word) > 1 {
				count[word]++
			}
		}
	// add map values to slice for sorting
	for key, value := range count {
		word_freq = append(word_freq, []any{key, value})
	}

	//sort
	for i := 0; i < len(word_freq); i++ {
		
		for n := len(word_freq) - 1; n > i; n-- {
			
			if word_freq[n][1].(int) > word_freq[n-1][1].(int) {
				word_freq[n], word_freq[n-1] = word_freq[n-1], word_freq[n]
			}
		}
	}

	// template strings for output
	for _, pair := range word_freq {
		fmt.Printf("%s - %d\n", pair[0], pair[1])
	}
}

func main() {
	readFiles()
}