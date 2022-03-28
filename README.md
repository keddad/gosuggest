# gosuggest
Simple suggestion engine for films (among other things) written in Go as part of VK internship tests.

* ```data.csv``` - 5900 film titles with ratings, generated with ```data_extractor.py``` from [Dataset](https://www.kaggle.com/datasets/rounakbanik/the-movies-dataset). Only movies with ratings and titles are present.
* ```trie``` is a Golang module which contains a "Fuzzy Trie" - modified Trie data structure which supports finding a "close match" among words, with BFS-like traversal. To rank the matches, it uses length difference, Levenshtein distance and film's score. It has some unit tests.
* ```main.go``` - simple CLI to showcase the functionality of the app.