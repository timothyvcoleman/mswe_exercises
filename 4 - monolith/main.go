package main

import (
	"os"
	"fmt"
	"strings"
	"unicode"
	"strconv"
)

func main() {
	// initializes an empty array to store arrays that consist of valid words and their frequency
	var word_freq = [][]any{}
	// opens "stop_words.txt" from parent directory, converts from bytes to string, create list of words
	rawStopWords, _ := (os.ReadFile("../stop_words.txt"))
	stringStopWords := string(rawStopWords)
	stopWords := strings.Split(stringStopWords, ",")

	stopWordsMap := make(map[string]bool)
	for _, word := range stopWords {
		stopWordsMap[word] = true
	}

	// opens a file referenced in the command line from the current directory
	rawBook, _ := (os.ReadFile(os.Args[1]))
	stringBook := string(rawBook)
	lowerCase := strings.ToLower(stringBook)


	// only select the lines of actual content
	contentStart, _ := strconv.Atoi(os.Args[2])
	contentEnd, _ := strconv.Atoi(os.Args[3])
	contentOnly := strings.Split(lowerCase, "\n")[contentStart:contentEnd]


	// iterate over every line of file
	for _, line := range contentOnly {
		startIndex := -1
		currentIndex := 0

		for _, char := range line {
			if startIndex == -1 {
				if unicode.IsLetter(char){
					startIndex = currentIndex
				}
			} else if startIndex != -1 && !unicode.IsLetter(char) {
				// process the word
				found := false
				word := line[startIndex:currentIndex]
				// check if the word is not a stop word
				if !stopWordsMap[word] && len(word) > 1 {
					pair_index := 0
					for _, pair := range word_freq {
						if word == pair[0] {
							pair[1] = pair[1].(int) + 1
							found = true
							break
						}
						pair_index += 1
					}
					if !found {
						word_freq = append(word_freq, []any{word, 1})
					} else if len(word_freq) > 1 {
						for n := pair_index - 1; n >= 0; n-- {
							if word_freq[pair_index][1].(int) > word_freq[n][1].(int) {
									word_freq[n], word_freq[pair_index] = word_freq[pair_index], word_freq[n]
									pair_index = n
							}
						}
					}	
				}
				startIndex = -1
			}
			currentIndex += 1
		}
	}

	for i := 0; i < 25; i++ {
		fmt.Printf("%s - %d\n", word_freq[i][0], word_freq[i][1])
	}
}
