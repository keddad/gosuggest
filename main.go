package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/keddad/gosuggest/trie"
	"log"
	"os"
	"strconv"
)

func readCsvFile(filePath string) [][]string {
	// https://stackoverflow.com/questions/24999079/reading-csv-file-in-go
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func main() {
	lines := readCsvFile("data.csv")
	localTrie := trie.InitTrie()

	for _, line := range lines {
		score, _ := strconv.Atoi(line[2])
		localTrie.Insert([]rune(line[1]), score)
	}

	fmt.Printf("Trie built. Size is: %d entries\n", localTrie.Words)
	reader := bufio.NewReader(os.Stdin)

	for true {
		fmt.Print("Input search query:\n>")
		text, _ := reader.ReadString('\n')

		matches := localTrie.FindClosest([]rune(text[:len(text)-1]))

		for i, match := range matches {
			fmt.Printf("%d: %s\n", i, string(match))
		}

	}
}
