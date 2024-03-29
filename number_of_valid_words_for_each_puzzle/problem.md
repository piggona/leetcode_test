# 1178. Number of Valid Words for Each Puzzle


## 题目
With respect to a given puzzle string, a word is valid if both the following conditions are satisfied:
* word contains the first letter of puzzle.
* For each letter in word, that letter is in puzzle.
For example, if the puzzle is "abcdefg", then valid words are "faced", "cabbage", and "baggage"; while invalid words are "beefed" (doesn't include "a") and "based" (includes "s" which isn't in the puzzle).

Return an array answer, where answer[i] is the number of words in the given word list words that are valid with respect to the puzzle puzzles[i].

## 例子

### Input:

words = ["aaaa","asas","able","ability","actt","actor","access"],

puzzles = ["aboveyz","abrodyz","abslute","absoryz","actresz","gaswxyz"]

Output: [1,1,3,2,4,0]

### Explanation:

1 valid word for "aboveyz" : "aaaa" 

1 valid word for "abrodyz" : "aaaa"

3 valid words for "abslute" : "aaaa", "asas", "able"

2 valid words for "absoryz" : "aaaa", "asas"

4 valid words for "actresz" : "aaaa", "asas", "actt", "access"

There're no valid words for "gaswxyz" cause none of the words in the list contains letter 'g'.

### Constraints:

1 <= words.length <= 10^5

4 <= words[i].length <= 50

1 <= puzzles.length <= 10^4

puzzles[i].length == 7

words[i][j], puzzles[i][j] are English lowercase letters.

Each puzzles[i] doesn't contain repeated characters.

## 题解

题目主要考察的有两点：
1. 表示穷举容量确定且与顺序无关的元素组合：bitmap
2. 如何**穷举一个bitmap所表示的元素集的所有<font color="red">子集</font>**：(n-1) & bitmap
> 前提是一个字符集存储在bitmap上，然后最初设置n=bitmap，这样(n-1)&bitmap就能得到元素集的所有子集
