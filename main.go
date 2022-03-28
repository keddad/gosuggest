package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	trie "github.com/keddad/gosuggest/trie"
	"log"
	"os"
	"strconv"
)

func getRealSizeOf(v interface{}) int {
	b := new(bytes.Buffer)
	return b.Len()
}

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
	trie := trie.InitTrie()

	for _, line := range lines {
		score, _ := strconv.Atoi(line[2])
		trie.Insert([]rune(line[1]), score)
	}

	fmt.Printf("Trie built. Size is: %d \n", trie.Words)
	reader := bufio.NewReader(os.Stdin)

	for true {
		fmt.Print("Input search query:\n>")
		text, _ := reader.ReadString('\n')

		matches := trie.FindClosest([]rune(text[:len(text)-1]))

		for i, match := range matches {
			fmt.Printf("%d: %s\n", i, string(match))
		}

	}
}
