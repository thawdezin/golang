package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	// The document we want to calculate the TF-IDF for
	document := "The quick brown fox jumps over the lazy dog."

	// Split the document into a slice of words
	words := strings.Fields(document)

	// The term we want to calculate the TF-IDF for
	term := "fox"

	// The total number of words in the document
	wordCount := len(words)

	// The number of times the term appears in the document
	termCount := 0
	for _, word := range words {
		if word == term {
			termCount++
		}
	}

	// The term frequency (TF) is the number of times the term appears in the document divided by the total number of words in the document
	tf := float64(termCount) / float64(wordCount)

	// In this example, we will assume that the term appears in all documents in the collection
	// If the term appears in N documents, the inverse document frequency (IDF) is calculated as log(D/N), where D is the total number of documents in the collection
	idf := math.Log(float64(1) / float64(1))

	// The TF-IDF is the product of the TF and IDF
	tfidf := tf * idf

	fmt.Printf("TF-IDF of term '%s' in document '%s': %f\n", term, document, tfidf)
}
