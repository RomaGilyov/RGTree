package trie

import (
	"testing"
)

func testTrie() *Trie {
	trie := Construct()

	trie.Add([]interface{}{"c", "a", "r"})
	trie.Add([]interface{}{"c", "a", "t"})
	trie.Add([]interface{}{"c", "a", "r", "s"})
	trie.Add([]interface{}{"c", "a", "r", "d"})
	trie.Add([]interface{}{"d", "a", "y"})
	trie.Add([]interface{}{"w", "i", "n"})
	trie.Add([]interface{}{"w", "i", "n", "d", "o", "w"})
	trie.Add([]interface{}{"c", "l", "o", "s", "e"})
	trie.Add([]interface{}{"i", "m", "l", "e", "m", "e", "n", "t", "a", "t", "i", "n"})

	return trie
}

//func testDict() [][]interface{} { todo
//	dict := []string{
//		"cat",
//		"dog",
//		"car",
//		"ball",
//		"analogy",
//		"dufference",
//		"dufference",
//	}
//}

func TestAdd(t *testing.T) {
	trie := testTrie()

	trie.Add([]interface{}{"d", "o", "o", "r"})

	if _, ok := trie.Children["d"].Children["o"].Children["o"].Children["r"]; ! ok {
		t.Fatal("must have word door")
	}
}

func TestExists(t *testing.T) {
	trie := testTrie()

	if trie.Exists([]interface{}{"w", "i", "n"}) == false {
		t.Fatal("word win exists in the trie")
	}
}

func TestSuffix(t *testing.T) {
	trie := testTrie()

	s, err := trie.GetSuffix([]interface{}{"w", "i", "n"})

	if err != nil {
		t.Fatal("must have suffix node to win")
	}

	if s.Val != "n" {
		t.Fatal("suffix node must be `n`")
	}
}

func TestPrefixSearch(t *testing.T) {
	trie := testTrie()

	dict := trie.PrefixSearch([]interface{}{"c", "a"})

	if len(dict) != 3 {
		t.Fatal("dict must have 3 words")
	}
}

func TestRemove(t *testing.T) {
	trie := testTrie()

	trie.Remove([]interface{}{"w", "i", "n"})

	if _, ok := trie.Children["w"].Children["i"].Children["n"]; ! ok {
		t.Fatal("trie can not remove win word because it is prefix of word window")
	}

	trie.Remove([]interface{}{"w", "i", "n", "d", "o", "w"})

	if _, ok := trie.Children["w"]; ok {
		t.Fatal("window word must be removed")
	}
}

func TestFlatten(t *testing.T) {
	trie := testTrie()

	dict := trie.Flatten()

	if len(dict) != 7 {
		t.Fatal("must have 7 words")
	}
}

func BenchmarkExists(b *testing.B) {
	trie := testTrie()

	for n := 0; n < b.N; n++ {
		trie.Exists([]interface{}{"w", "i", "n"})
	}
}

func BenchmarkExistsPlainDict(b *testing.B) {
	trie := testTrie()

	dict := trie.Flatten()

	needle := []interface{}{"w", "i", "n"}

	for n := 0; n < b.N; n++ {
		plainSearchBench(needle, dict)
	}
}

func plainSearchBench(needle []interface{}, haystack [][]interface{}) bool {
	for _, word := range haystack {
		found := false

		for i, l := range needle {
			if word[i] != l {
				found = false
				break
			} else {
				found = true
			}
		}

		if found == true {
			return true
		}
	}

	return false
}
