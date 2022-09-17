# gosuggest
Simple Go suggestion engine for films (among other things).

* ```data.csv``` - 5900 film titles with ratings, generated with ```data_extractor.py``` from [dataset](https://www.kaggle.com/datasets/rounakbanik/the-movies-dataset). Only movies with ratings and titles are present.
* ```trie``` is a Golang module which contains a "Fuzzy Trie" - modified Trie data structure which supports finding a "close match" among words, with BFS-like traversal. To rank the matches, it uses length difference, Levenshtein distance and film's score. It has some unit tests.
* ```main.go``` - simple CLI to showcase the functionality of the app.

Example:
```
./gosuggest
Trie built. Size is: 5828 entries
Input search query:
>Десять негритят
0: Десять негритят
1: Marfa Girl
2: Место на земле
3: 12 Storeys
4: Sudba Cheloveka
Input search query:
>Десять нег
0: Десять негритят
1: 숙명
2: Аэлита
3: Платон
4: Шестой
Input search query:
>Butman
0: Bean
1: Batman
2: 숙명
3: Hitch
4: Laura
Input search query:
>Batman
0: Batman
1: Batman Begins
2: Batman Returns
3: Batman Forever
4: Batman & Robin
Input search query:
>Superman
0: Superman
1: Superman vs. The Elite
2: Bean
3: Oblivion
4: Seven
```
