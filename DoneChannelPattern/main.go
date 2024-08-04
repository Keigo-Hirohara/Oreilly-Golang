package main

import (
	"fmt"
	"strings"
)

func main() {
	var searchers []func(string) []string = []func(string) []string{
		searchForGoogle,
		searchForTDL,
		searchForDisneyCharDB,
	}

	searchString := "asdf"

	results := searchData(searchString, searchers)

	for _, result := range results {
		fmt.Println(result)
	}
}

func searchData(s string, searchers []func(string) []string) []string {
	done := make(chan struct{})
	resultChan := make(chan []string)

	for _, searcher := range searchers {
		go func(f func(string) []string) {
			select {
			case resultChan <- f(s):
				fmt.Println("結果が返ってきた")
			case <-done:
				fmt.Println("doneを選択")
			}
		}(searcher)
	}

	result := <-resultChan
	close(done)
	return result
}

func searchForGoogle(input string) []string {
	words := []string{
		"asdf",
		"asdfasdf",
		"fhfhfhfhfhfhfh",
	}
	var result []string
	for _, word := range words {
		if strings.Contains(word, input) {
			return append(result, word)
		}
	}
	return result
}

func searchForDisneyCharDB(input string) []string {
	words := []string{
		"mickey",
		"minnie",
		"goofy",
	}
	var result []string
	for _, word := range words {
		if strings.Contains(word, input) {
			return append(result, word)
		}
	}
	return result
}

func searchForTDL(input string) []string {
	words := []string{
		"かがみん",
		"舞浜",
		"浦安",
	}
	var result []string
	for _, word := range words {
		if strings.Contains(word, input) {
			return append(result, word)
		}
	}
	return result
}
