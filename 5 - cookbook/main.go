// Exercise 5: Implement Term Frequency Program
// Timothy Coleman
// 10/12/2026
// MSWE 262P


package main


import (
   "os"
   "fmt"
   "strings"
   "strconv"
   "slices"
)

var word_freq = [][]any{}
var words = []any{}
var data = []string{}

func read_file(){
	// opens a file referenced in the command line from the current directory
   rawBook, _ := (os.ReadFile(os.Args[1]))
   stringBook := string(rawBook)
	data = append(data, stringBook)
}

func reduce_file(){
	// only select the lines of actual content
   contentStart, _ := strconv.Atoi(os.Args[2])
   contentEnd, _ := strconv.Atoi(os.Args[3])
   contentOnly := strings.Split(data[0], "\n")[contentStart:contentEnd]
	mergedContent := strings.Join(contentOnly, " ")
   data[0] = mergedContent
}

func filter_chars_and_normalize() {
	lowerCase := strings.ToLower(data[0])
	// Create the replacer once with all your old/new pairs
	replacer := strings.NewReplacer(
		",", "",
		".", "",
		"?", "",
		"!", "",
		"\"", "",
		"/", "",
		"'", "",
		"(", "",
		")", "",
		"[", "",
		"]", "",
		":", "",
		";", "",
		"_", "",
		"--", "",
		"*", "",
		"|", "",
		"1", "",
		"2","",
		"3", "",
		"4", "",
		"5","",
		"6", "",
		"7", "",
		"8","",
		"9", "",
		"0", "",
	)
	// Execute all replacements in a single pass over the string
	cleanBook := replacer.Replace(lowerCase)
	data[0] = cleanBook
}

func scan() {
	splitWords := strings.Split(data[0], " ")
	for _, w := range splitWords {
		words = append(words, w)
	}
}

func remove_stop_words(){
   rawStopWords, _ := (os.ReadFile("../stop_words.txt"))
   stringStopWords := string(rawStopWords)
   stopWords := strings.Split(stringStopWords, ",")

   stopWordsMap := make(map[string]bool)
   for _, word := range stopWords {
       stopWordsMap[word] = true
   }

   for i := 0; i < len(words); i++ {
		if stopWordsMap[words[i].(string)] {
			words = slices.Delete(words, i, i+1)
			i--
		} else if len(words[i].(string)) <= 1 {
			words = slices.Delete(words, i, i+1)
			i--
		}
   }
}

func frequencies() {
	for _, word := range words {
		found := false
		if len(word_freq) > 0 {
			for i, pairs := range word_freq {
				if pairs[0].(string) == word.(string) {
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
}


func sort() {
	// The outer loop dictates how many sweeps to make
	for i := 0; i < len(word_freq); i++ {
		
		// The inner loop sweeps backward, stopping at 'i' so it 
		// doesn't re-check items that are already perfectly sorted
		for n := len(word_freq) - 1; n > i; n-- {
			
			// Compare the current item with the one before it
			if word_freq[n][1].(int) > word_freq[n-1][1].(int) {
				// Swap them if the current item is larger
				word_freq[n], word_freq[n-1] = word_freq[n-1], word_freq[n]
			}
		}
	}
}


func main() {
	read_file()
	reduce_file()
	filter_chars_and_normalize()
	scan()
	remove_stop_words()
	frequencies()
	sort()

	for i := 0; i < 25; i++ {
		fmt.Printf("%s - %d\n", word_freq[i][0], word_freq[i][1])
	}
}
