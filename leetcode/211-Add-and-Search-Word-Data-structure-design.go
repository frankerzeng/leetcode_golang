/*
211. Add and Search Word - Data structure design

Design a data structure that supports the following two operations:

void addWord(word)
bool search(word)
search(word) can search a literal word or a regular expression string containing only letters a-z or .
,A . means it can represent any one letter.

Example:
	addWord("bad")
	addWord("dad")
	addWord("mad")
	search("pad") -> false
	search("bad") -> true
	search(".ad") -> true
	search("b..") -> true

Note:
	You may assume that all words are consist of lowercase letters a-z.
*/
package main

import "fmt"

func main() {
	obj := Constructor()
	obj.AddWord("a")
	obj.AddWord("a")

	fmt.Println(obj.Search("."))
	fmt.Println(obj.Search("a"))
	fmt.Println(obj.Search("aa"))
	fmt.Println(obj.Search("a"))
	fmt.Println(obj.Search(".a"))
	fmt.Println(obj.Search("a."))
}

type WordDictionary struct {
	dictMap map[int][]string
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{dictMap: make(map[int][]string)}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	lenWord := len(word)
	if this.dictMap[lenWord] != nil {
		this.dictMap[lenWord] = append(this.dictMap[lenWord], word)
	} else {
		this.dictMap[lenWord] = []string{word}
	}
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	if this.dictMap[len(word)] == nil {
		return false
	}
	for i := 0; i < len(this.dictMap[len(word)]); i++ {
		tmpA := true
		for j := 0; j < len(word); j++ {
			if word[j] != '.' && this.dictMap[len(word)][i][j] != word[j] {
				tmpA = false
				break
			}
		}
		if tmpA {
			return true
		}
	}
	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
