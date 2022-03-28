package trie

import (
	"reflect"
	"testing"
)

func TestInsertTrie(t *testing.T) {
	trie := InitTrie()

	trie.Insert([]rune("OK"), 42)

	if trie.root.children[0].currentChar != 'O' {
		t.Error("Insert failed")
	}

	if trie.root.children[0].children[0].currentChar != 'K' {
		t.Error("Insert failed")
	}

	if !trie.root.children[0].children[0].isEnd {
		t.Error("Insert failed")
	}

	if trie.root.children[0].children[0].score != 42 {
		t.Error("Insert failed")
	}

	if trie.root.children[0].children[0].parent != trie.root.children[0] {
		t.Error("Insert failed")
	}
}

func TestNodeGetWord(t *testing.T) {
	trie := InitTrie()

	trie.Insert([]rune("OK"), 42)

	if !reflect.DeepEqual(trie.root.children[0].children[0].getWord(), []rune("OK")) {
		t.Error("GetWord failed")
	}
}

func TestSearchExactMatch(t *testing.T) {
	trie := InitTrie()

	trie.Insert([]rune("Elephant"), 3)
	trie.Insert([]rune("Elephunt"), 3)
	trie.Insert([]rune("Dino"), 5)
	trie.Insert([]rune("Owl"), 1)
	trie.Insert([]rune("Lion"), 4)
	trie.Insert([]rune("Tiger"), 2)
	trie.Insert([]rune("Human"), 4)
	trie.Insert([]rune("Cow"), 2)

	answ := trie.FindClosest([]rune("Elephant"))

	if string(answ[0]) != "Elephant" {
		t.Error("Search failed")
	}

}

func TestSearchSimilar(t *testing.T) {
	trie := InitTrie()

	trie.Insert([]rune("Elephant"), 3)
	trie.Insert([]rune("Elephunt"), 3)
	trie.Insert([]rune("Dino"), 5)
	trie.Insert([]rune("Owl"), 1)
	trie.Insert([]rune("Lion"), 4)
	trie.Insert([]rune("Tiger"), 2)
	trie.Insert([]rune("Human"), 4)
	trie.Insert([]rune("Cow"), 2)

	answ := trie.FindClosest([]rune("Elephunt"))

	for _, v := range answ {
		if string(v) == "Elephant" {
			return
		}
	}

	t.Error("Search failed")
}

func TestSearchContinuous(t *testing.T) {
	trie := InitTrie()

	trie.Insert([]rune("Elephant"), 3)
	trie.Insert([]rune("Elephunt"), 3)
	trie.Insert([]rune("Dino"), 5)
	trie.Insert([]rune("Owl"), 1)
	trie.Insert([]rune("Lion"), 4)
	trie.Insert([]rune("Tiger"), 2)
	trie.Insert([]rune("Human"), 4)
	trie.Insert([]rune("Cow"), 2)

	answ := trie.FindClosest([]rune("Elep"))

	for _, v := range answ {
		if string(v) == "Elephant" {
			return
		}
	}

	t.Error("Search failed")
}
