package main

/*
A golang implementation of the Rapid Automatic Keyword Extraction (RAKE)
algorithm as described in: Rose, S., Engel, D., Cramer, N., & Cowley, W. (2010).
Automatic Keyword Extraction from Individual Documents.
In M. W. Berry & J. Kogan (Eds.), Text Mining: Theory and Applications: John Wiley & Sons.

This is a port of the python implementation available at: https://github.com/aneesha/RAKE

The MIT License (MIT)

Copyright (c) 2015 Wolfgang Meyers

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

var wordSep = regexp.MustCompile("\\s+")
var wordSplitter = regexp.MustCompile("[^\\p{L}\\p{N}\\+\\-/]")

func sentenceSplitter(r rune) bool {
	// Ignore hyphens, usually they are used to join words (e.g., "ice-cream-flavored candy")
	return unicode.IsPunct(r) && !unicode.Is(unicode.Properties["Hyphen"], r)
}

func IsAcceptable(phrase string, minCharLength int, maxWordsLength int) bool {
	// a phrase must have a min length in characters
	if len(phrase) < minCharLength {
		return false
	}

	// a phrase must have a max number of words
	words := wordSep.Split(phrase, -1)
	if len(words) > maxWordsLength {
		return false
	}
	digits := 0
	alpha := 0
	for _, c := range phrase {
		if unicode.IsDigit(c) {
			digits++
		} else if unicode.IsLetter(c) {
			alpha++
		}
	}
	// a phrase must have at least one alpha character
	if alpha == 0 {
		return false
	}
	// a phrase must have more alpha than digits characters
	if digits > alpha {
		return false
	}
	return true
}

func findStopwordIndices(words []string, stopwords map[string]bool) []int {
	result := []int{}
	if stopwords != nil {
		for i, word := range words {
			if stopwords[strings.ToLower(word)] {
				result = append(result, i)
			}
		}
	}
	return result
}

func GenerateCandidateKeywords(sentenceList []string, stopwords map[string]bool, minCharLength int, maxWordsLength int) []string {
	phraseList := []string{}
	for _, s := range sentenceList {
		words := strings.Fields(s)
		stopwordIndices := findStopwordIndices(words, stopwords)

		if len(stopwordIndices) > 0 {
			if stopwordIndices[0] != 0 {
				phraseWords := words[0:stopwordIndices[0]]
				phrase := strings.Join(phraseWords, " ")
				if phrase != "" && IsAcceptable(phrase, minCharLength, maxWordsLength) {
					phraseList = append(phraseList, phrase)
				}
			}
			for i, index := range stopwordIndices {
				j := i + 1
				if j < len(stopwordIndices) {
					index2 := stopwordIndices[j]
					if index2-index == 1 {
						continue
					}
					phraseWords := words[index+1 : index2]
					phrase := strings.Join(phraseWords, " ")
					if phrase != "" && IsAcceptable(phrase, minCharLength, maxWordsLength) {
						phraseList = append(phraseList, phrase)
					}
				}
			}
			if stopwordIndices[len(stopwordIndices)-1] != len(words)-1 {
				index := stopwordIndices[len(stopwordIndices)-1]
				phraseWords := words[index+1:]
				phrase := strings.Join(phraseWords, " ")
				if phrase != "" && IsAcceptable(phrase, minCharLength, maxWordsLength) {
					phraseList = append(phraseList, phrase)
				}
			}
		} else {
			phrase := strings.Join(words, " ")
			if phrase != "" && IsAcceptable(phrase, minCharLength, maxWordsLength) {
				phraseList = append(phraseList, phrase)
			}
		}
	}
	return phraseList
}

func IsNumber(s string) bool {
	if strings.Index(s, ".") != -1 {
		_, err := strconv.ParseFloat(s, 32)
		return err == nil
	} else {
		_, err := strconv.ParseInt(s, 10, 32)
		return err == nil
	}
}

func SeparateWords(text string, minWordReturnSize int) []string {
	//    Utility function to return a list of all words that are have a length
	//    greater than a specified number of characters.
	//    @param text The text that must be split in to words.
	//    @param min_word_return_size The minimum no of characters
	//           a word must have to be included.
	words := []string{}
	for _, singleWord := range wordSplitter.Split(text, -1) {
		currentWord := strings.ToLower(strings.TrimSpace(singleWord))
		// leave numbers in phrase, but don't count as words, since they tend to invalidate scores of their phrases
		if len(currentWord) > minWordReturnSize && currentWord != "" && !IsNumber(currentWord) {
			words = append(words, currentWord)
		}
	}
	return words
}

func CalculateWordScores(phraseList []string) map[string]float64 {
	wordFrequency := map[string]int{}
	wordDegree := map[string]int{}
	for _, phrase := range phraseList {
		wordList := SeparateWords(phrase, 0)
		wordListLength := len(wordList)
		wordListDegree := wordListLength - 1
		// if word_list_degree > 3: word_list_degree = 3 #exp.
		for _, word := range wordList {
			wordFrequency[word]++
			wordDegree[word] += wordListDegree
		}
	}
	for item := range wordFrequency {
		wordDegree[item] = wordDegree[item] + wordFrequency[item]
	}
	// Calculate Word scores = deg(w)/frew(w)
	wordScore := map[string]float64{}
	for item := range wordFrequency {
		wordScore[item] = float64(wordDegree[item]) / float64(wordFrequency[item])
	}
	return wordScore
}

func StringCount(stringList []string) map[string]int {
	stringCount := map[string]int{}
	for _, item := range stringList {
		stringCount[item]++
	}
	return stringCount
}

func GenerateCandidateKeywordScores(phraseList []string, wordScore map[string]float64, minKeywordFrequency int) (map[string]float64, map[string]int) {
	keywordCandidates := map[string]float64{}
	phraseCounts := StringCount(phraseList)
	for _, phrase := range phraseList {
		if minKeywordFrequency > 1 {
			if phraseCounts[phrase] < minKeywordFrequency {
				continue
			}
		}
		wordList := SeparateWords(phrase, 0)
		var candiateScore float64
		for _, word := range wordList {
			candiateScore += wordScore[word]
		}
		keywordCandidates[phrase] = candiateScore
	}

	return keywordCandidates, phraseCounts
}

func MapToKeywordScores(keywordScoreMap map[string]float64, keywordCountMap map[string]int) []KeywordScore {
	keywordScores := []KeywordScore{}
	for key, value := range keywordScoreMap {
		keywordScore := KeywordScore{
			Keyword: key,
			Score:   value,
		}
		count, _ := keywordCountMap[key]
		keywordScore.Count = count
		keywordScores = append(keywordScores, keywordScore)
	}
	return keywordScores
}

func loadStopWords(stopWordFile string) []string {
	stopWords := []string{}
	contents, err := ioutil.ReadFile(stopWordFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lines := strings.Split(string(contents), "\n")
	for _, line := range lines {
		line = strings.ToLower(strings.TrimSpace(line))
		if strings.Index(line, "#") != 0 {
			// in case more than one per line
			for _, word := range strings.Fields(line) {
				stopWords = append(stopWords, word)
			}

		}
	}
	return stopWords
}

func BuildStopWordMap(stopWordList []string) map[string]bool {
	stopWordMap := make(map[string]bool, len(stopWordList))
	for _, word := range stopWordList {
		stopWordMap[strings.ToLower(word)] = true
	}
	return stopWordMap
}

// ByScore implements sort.Interface for []KeywordScore based on
// the Score field. Sorts in reverse order.
type ByScore []KeywordScore

func (a ByScore) Len() int           { return len(a) }
func (a ByScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByScore) Less(i, j int) bool { return a[i].Score > a[j].Score }

type Rake struct {
	stopWordsPath       string
	stopWords           map[string]bool
	minCharLength       int
	maxWordsLength      int
	minKeywordFrequency int
}

type KeywordScore struct {
	Keyword string
	Score   float64
	Count   int
}

func NewRake(stopWordsPath string, minCharLength int, maxWordsLength int, minKeywordFrequency int) *Rake {
	rake := &Rake{
		minCharLength:       minCharLength,
		maxWordsLength:      maxWordsLength,
		minKeywordFrequency: minKeywordFrequency,
	}

	if stopWordsPath != "" {
		rake.stopWords = BuildStopWordMap(loadStopWords(stopWordsPath))
	}

	return rake
}

func (rake *Rake) SetStopWords(stopWordsList []string) {
	rake.stopWords = BuildStopWordMap(stopWordsList)
}

func (rake *Rake) Run(text string) []KeywordScore {
	sentenceList := strings.FieldsFunc(text, sentenceSplitter)
	phraseList := GenerateCandidateKeywords(sentenceList, rake.stopWords, rake.minCharLength, rake.maxWordsLength)
	wordScores := CalculateWordScores(phraseList)
	keywordCandidates, keywordCounts := GenerateCandidateKeywordScores(phraseList, wordScores, rake.minKeywordFrequency)
	keywordScores := MapToKeywordScores(keywordCandidates, keywordCounts)
	sort.Sort(ByScore(keywordScores))
	return keywordScores
}
