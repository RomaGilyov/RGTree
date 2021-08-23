package trie

import (
	"fmt"
	"testing"
)

func TestMakeTreeMap(t *testing.T) {
	word := []interface{}{"c", "a", "r"}

	trie := Construct(word)

	trie.Add([]interface{}{"c", "a", "t"})
	trie.Add([]interface{}{"c", "a", "r", "d"})
	trie.Add([]interface{}{"c", "a", "r", "s"})

	fmt.Println(trie.PrefixSearch([]interface{}{"c", "a"}))
}
