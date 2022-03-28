package trie

import (
	"github.com/agnivade/levenshtein"
	"github.com/gammazero/deque"
	"sort"
)

const (
	SearchRange  = 15
	QueryAnswers = 5
)

type node struct {
	parent      *node
	children    []*node
	currentChar rune
	isEnd       bool
	score       int
}

func (N *node) getWord() []rune {
	var word []rune
	for N.parent != nil {
		word = append(word, N.currentChar)
		N = N.parent
	}

	for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
		word[i], word[j] = word[j], word[i]
	}

	return word
}

type Trie struct {
	root *node
}

func InitTrie() *Trie {
	return &Trie{
		root: &node{},
	}
}

func binarySearch(arr []*node, target rune) (*node, bool) {
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid].currentChar == target {
			return arr[mid], true
		} else if arr[mid].currentChar < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return nil, false
}

func (T *Trie) Insert(word []rune, word_score int) {
	currentNode := T.root

	for currentChar := 0; currentChar < len(word); currentChar++ {
		next, found := binarySearch(currentNode.children, word[currentChar])

		if !found {
			next = &node{
				currentChar: word[currentChar],
				parent:      currentNode,
			}
			currentNode.children = append(currentNode.children, next)
			sort.Slice(currentNode.children, func(i, j int) bool {
				return currentNode.children[i].currentChar < currentNode.children[j].currentChar
			})

		}

		currentNode = next
	}

	currentNode.isEnd = true
	currentNode.score = word_score
}

func (T *Trie) FindClosest(word []rune) [][]rune {
	currentNode := T.root

	for i := 0; i < len(word); i++ {
		next, found := binarySearch(currentNode.children, word[i])

		if !found {
			break
		}

		currentNode = next
	}

	type queueItem struct {
		node *node
		hops int
	}

	type foundItem struct {
		word  []rune
		score int
	}

	var q deque.Deque
	foundItems := make([]foundItem, 0)
	metNodes := make(map[*node]bool)

	q.PushBack(&queueItem{node: currentNode, hops: 0})

	for q.Len() != 0 {
		item := q.PopFront().(*queueItem)

		if item.node == nil {
			continue
		}

		metNodes[item.node] = true

		if item.node.isEnd {
			currentWord := item.node.getWord()
			foundItems = append(foundItems, foundItem{word: currentWord, score: 10 + item.node.score - (levenshtein.ComputeDistance(string(word), string(currentWord)) - len(word) - len(currentWord))}) // Closest matches are more important than just good options. If the word is bigger then query, forgive some levenshtein penalty.
		}

		if item.hops < SearchRange {
			for _, child := range item.node.children {
				if !metNodes[child] {
					q.PushBack(&queueItem{node: child, hops: item.hops + 1})
				}
			}

			if !metNodes[item.node.parent] {
				q.PushBack(&queueItem{node: item.node.parent, hops: item.hops + 1})
			}
		}
	}

	sort.Slice(foundItems, func(i, j int) bool {
		return foundItems[i].score > foundItems[j].score
	})

	answ := make([][]rune, 0, QueryAnswers)

	for i := 0; i < QueryAnswers && i < len(foundItems); i++ {
		answ = append(answ, foundItems[i].word)
	}

	return answ
}
