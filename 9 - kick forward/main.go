package main


import (
   "os"
   "fmt"
   "strings"
   "strconv"
   "slices"
   "regexp"
)

// 
func read_file(data string, function func(string, func(func(string, func(func(string, func(func(string, func(func(string, func(func(string, func())))))))))))){
	rawBook, _ := os.ReadFile(os.Args[1])
	stringBook := string(rawBook)

	function(stringBook, filter_chars)
}

func filter_chars(data string, function func(func(string, func(func(string, func(func(string, func(func(string, func(func(string, func()))))))))))) {
	

	function(bookWords, scan)
}

func normalize(data string, function func(string, func(func(string, func(func(string, func(func(string, func())))))))) {
	lowerCaseWords := strings.ToLower(data)

	function(lowerCaseWords, remove_stop_words())
}

func scan(data string, function func(string, func(func(string, func(func(string, func())))))) {
	splitWords := strings.Split(cleanBook, " ")
	return splitWords
}

func remove_stop_words(data string, function func(string, func(func(string, func())))) []string {
   rawStopWords, _ := (os.ReadFile("../stop_words.txt"))
   stringStopWords := string(rawStopWords)
   stopWords := strings.Split(stringStopWords, ",")

   stopWordsMap := make(map[string]bool)
   for _, word := range stopWords {
       stopWordsMap[word] = true
   }

   for i := 0; i < len(splitWords); i++ {
		if stopWordsMap[splitWords[i]] {
			splitWords = slices.Delete(splitWords, i, i+1)
			i--
		} else if len(splitWords[i]) <= 1 {
			splitWords = slices.Delete(splitWords, i, i+1)
			i--
		}
   }
   validWords := splitWords
   return validWords
}

func frequencies(data string, function func(string, func(func()))) [][]any {
	word_freq := [][]any{}

	for _, word := range validWords {
		found := false
		if len(word_freq) > 0 {
			for i, pairs := range word_freq {
				if pairs[0] == word {
					word_freq[i][1] = word_freq[i][1].(int) + 1
					found = true
					break
				}
			}
			if !found {
				word_freq = append(word_freq, []any{word, 1})
			}
		} else if len(word_freq) == 0 {
			word_freq = append(word_freq, []any{word, 1})
		}
	}

	function(data, print)
}


func sort(data string, function func(string, func())) [][]any {
	for i := 0; i < len(word_freq); i++ {
		
		for n := len(word_freq) - 1; n > i; n-- {
			
			if word_freq[n][1].(int) > word_freq[n-1][1].(int) {
				word_freq[n], word_freq[n-1] = word_freq[n-1], word_freq[n]
			}
		}
	}
	function(data, no_op)
}

func print(data string, function func()) {
	for i := 0; i < 25; i++ {
			fmt.Printf("%s - %d\n", word_freq[i][0], word_freq[i][1])
		}
	function()
}

func no_op() {
	return
}


func main() {
	read_file(os.Args[1], filter_chars())
}
