package main

type Trie struct {
	words [26]*Trie
	end   bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	var cur = this
	for _, b := range word {
		if cur.words[b-'a'] == nil {
			cur.words[b-'a'] = &Trie{}
		}
		cur = cur.words[b-'a']
	}
	cur.end = true
}

func (this *Trie) Search(word string) bool {
	var cur = this
	for _, b := range word {
		if cur.words[b-'a'] == nil {
			return false
		}
		cur = cur.words[b-'a']
	}
	return cur.end
}

func (this *Trie) StartsWith(word string) bool {
	var cur = this
	for _, b := range word {
		if cur.words[b-'a'] == nil {
			return false
		}
		cur = cur.words[b-'a']
	}
	return true
}
