package trie

import (
	"fmt"
	"testing"
)

func TestPrefixSearch(t *testing.T) {
	trie := Construct()

	trie.Add([]interface{}{"c", "a", "r"})
	trie.Add([]interface{}{"c", "a", "t"})
	trie.Add([]interface{}{"c", "a", "r", "s"})
	trie.Add([]interface{}{"c", "a", "r", "d"})

	fmt.Println(trie.PrefixSearch([]interface{}{}))
	fmt.Println(trie.Exists([]interface{}{"c", "r"}))
	//fmt.Println(trie.Children["c"].Children["a"].Children["r"])
}
