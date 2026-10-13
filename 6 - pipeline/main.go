// Exercise 6: Implement Term Frequency Program
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


func read_file(bookPath string, startBook string, endBook string) (string, string, string) {
	// opens a file referenced in the command line from the current directory
   rawBook, _ := (os.ReadFile(bookPath))
   stringBook := string(rawBook)
	return stringBook, startBook, endBook
}

func reduce_file(stringBook string, startBook string, endBook string) string {
	// only select the lines of actual content
   contentStart, _ := strconv.Atoi(startBook)
   contentEnd, _ := strconv.Atoi(endBook)
   contentOnly := strings.Split(stringBook, "\n")[contentStart:contentEnd]
	mergedContent := strings.Join(contentOnly, " ")
   return mergedContent
}

func filter_chars_and_normalize(mergedContent string) string {
	lowerCase := strings.ToLower(mergedContent)
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
	cleanBook := replacer.Replace(lowerCase)
	return cleanBook
}

func scan(cleanBook string) []string {
	splitWords := strings.Split(cleanBook, " ")
	return splitWords
}

func remove_stop_words(splitWords []string) []string {
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

func frequencies(validWords []string) [][]any {
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

	return word_freq
}


func sort(word_freq [][]any) [][]any {
	for i := 0; i < len(word_freq); i++ {
		
		for n := len(word_freq) - 1; n > i; n-- {
			
			if word_freq[n][1].(int) > word_freq[n-1][1].(int) {
				word_freq[n], word_freq[n-1] = word_freq[n-1], word_freq[n]
			}
		}
	}
	return word_freq
}

func printall(word_freq[][]any) {
for i := 0; i < 25; i++ {
		fmt.Printf("%s - %d\n", word_freq[i][0], word_freq[i][1])
	}
}


func main() {
	printall(sort(frequencies(remove_stop_words(scan(filter_chars_and_normalize(reduce_file(read_file(os.Args[1], os.Args[2], os.Args[3]))))))))
}
