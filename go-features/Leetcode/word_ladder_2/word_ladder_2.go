package main

import "fmt"

func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot","dot","dog","lot","log","cog"}
	fmt.Println(findLadders(beginWord, endWord, wordList))
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	var output [][]string
	var hasMatch bool
	currentWord := beginWord

	// check if endWord is in the list
	var hasEndWord bool
	for _, word := range wordList {
		if endWord == word {
			hasEndWord = true
			break
		}
	}
	if !hasEndWord {
		return nil
	}

	fmt.Println(hasMatch, currentWord)

	return output
}


