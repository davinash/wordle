package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"os"
	"strings"
	"unicode"
)

func IsAlphaOnly(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

var reader = bufio.NewReader(os.Stdin)

func scanf(s string, v ...interface{}) {
	fmt.Fscanf(reader, s, v...)
}

type Colors int

const (
	GREEN Colors = iota
	RED
	YELLOW
)

type Pair struct {
	ch    string
	color Colors
}

func main() {
	dictionaryFile := flag.String("dict-file", "/usr/share/dict/words", "Dictionary file with words seperated by newline")
	wordLen := flag.Int("word-len", 5, "Minimum length of the word")
	attempt := flag.Int("attempt", 6, "Number of attempts")

	flag.Parse()

	dictFile, err := os.Open(*dictionaryFile)
	if err != nil {
		fmt.Println("error opening file ", err)
		os.Exit(1)
	}

	defer func(dictFile *os.File) {
		err := dictFile.Close()
		if err != nil {
			fmt.Println("error closing file ", err)
			os.Exit(1)
		}
	}(dictFile)

	wordsMap := make(map[string]bool)
	scanner := bufio.NewScanner(dictFile)
	for scanner.Scan() {
		word := strings.ToUpper(scanner.Text())
		if len(word) == *wordLen && IsAlphaOnly(word) {
			wordsMap[word] = true
		}
	}
	var pWord string
	for pWord = range wordsMap {
	}
	fmt.Println("Word Chosen --> ", pWord)
	matchCount := 0
	for tryIdx := 0; tryIdx < *attempt; tryIdx++ {
		var w string
		scanf("%s\n", &w)
		w = strings.ToUpper(w)
		if len(w) != *wordLen {
			fmt.Printf("Length should be %d !!\n", *wordLen)
			continue
		}
		result := make([]Pair, *wordLen)
		for idx := 0; idx < *wordLen; idx++ {
			if pWord[idx] == w[idx] {
				result[idx] = Pair{ch: string(w[idx]), color: GREEN}
				matchCount++
				continue
			}
			if strings.Contains(pWord, string(w[idx])) {
				result[idx] = Pair{ch: string(w[idx]), color: YELLOW}
			} else {
				result[idx] = Pair{ch: string(w[idx]), color: RED}
			}
		}
		success := true
		for _, p := range result {
			switch p.color {
			case RED:
				c := color.New(color.FgRed, color.Bold)
				c.Printf("%s ", p.ch)
				success = false
			case GREEN:
				c := color.New(color.FgGreen, color.Bold)
				c.Printf("%s ", p.ch)
			case YELLOW:
				c := color.New(color.FgYellow, color.Bold)
				c.Printf("%s ", p.ch)
				success = false
			}
		}
		fmt.Println()
		if success {
			os.Exit(0)
		}
	}
	fmt.Println("Expected word => ", pWord)
}
