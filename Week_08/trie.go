package main

import (
	"fmt"
	"runtime/debug"
)

func init() {
	debug.FreeOSMemory()
	debug.SetGCPercent(-1)
}

type Trie struct {
	c  []*Trie
	ed bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	if len(word) == 0 {
		return
	}
	if this == nil {
		*this = Trie{}
	}
	for i := 0; i < len(word); i++ {
		c := word[i] - 'a'
		if this.c == nil {
			this.c = make([]*Trie, 26)
		}
		if this.c[c] == nil {
			this.c[c] = &Trie{}
		}
		this = this.c[c]
	}
	this.ed = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	if len(word) == 0 {
		return true
	}
	for i := 0; i < len(word); i++ {
		c := word[i] - 'a'
		if this == nil || this.c == nil || this.c[c] == nil {
			return false
		}
		this = this.c[c]
	}
	return this.ed
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return true
	}
	for i := 0; i < len(prefix); i++ {
		c := prefix[i] - 'a'
		if this == nil || this.c == nil || this.c[c] == nil {
			return false
		}
		this = this.c[c]
	}
	return true
}

func main() {
	trie := Constructor()
	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))
	fmt.Println(trie.Search("app"))
	fmt.Println(trie.StartsWith("app"))
	trie.Insert("app")
	fmt.Println(trie.Search("app"))
}
