package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	stopBytes, _ := os.ReadFile("../stop_words.txt"); stopWords := strings.Split(string(stopBytes), ",")
	bookBytes, _ := os.ReadFile(os.Args[1])
	bookWords := regexp.MustCompile("[^a-zA-Z]+").Split(strings.ToLower(string(bookBytes)), -1)
	word_freq := [][]string{}
	for _, word := range bookWords {
		if (!slices.Contains(stopWords, word) && len(word) > 1) {
			idx := slices.IndexFunc(word_freq, func(x []string) bool {
				if (len(x) > 0) { return x[0] == word } else { return false }})
			if (idx != -1) { y, _ := strconv.Atoi(word_freq[idx][1]); word_freq[idx][1] = strconv.Itoa(y + 1)} else {
				word_freq = append(word_freq, []string{word, "1"})}}}
	for i := 0; i < len(word_freq); i++ {
		for n := len(word_freq) - 1; n > i; n-- {
			current_index, _ := strconv.Atoi(word_freq[n][1]); lower_index, _ := strconv.Atoi(word_freq[n-1][1])
			if current_index > lower_index {
				word_freq[n], word_freq[n-1] = word_freq[n-1], word_freq[n]
			}}} ; for i := 0; i < 25; i++ { fmt.Printf("%s - %s\n", word_freq[i][0], word_freq[i][1])}}