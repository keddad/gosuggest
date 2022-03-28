package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/keddad/gosuggest/trie"
	"os"
	"strconv"
)

func readFile(file string) [][]string {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	lines, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	return lines
}

func main() {
	lines := readFile("data.csv")
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
