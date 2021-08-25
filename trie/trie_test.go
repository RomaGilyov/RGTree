package trie

import (
	"strings"
	"testing"
)

func testTrie() *Trie {
	trie := Construct()

	dict := testDict()

	for _, word := range dict {
		trie.Add(word)
	}

	return trie
}

func testDict() [][]interface{} {
	s := "Declare a main package In Go code executed as an application must be in a main package " +
		"Import two packages example com greetings and the fmt package " +
		"This gives your code access to functions in those packages " +
		"Importing example com greetings " +
		"the package contained in the module you created earlier " +
		"gives you access to the Hello function You also import fmt " +
		"with functions for handling input and output text " +
		"such as printing text to the console " +
		"Get a greeting by calling the greetings packages Hello function " +
		"For production use you would publish the example com greetings " +
		"module from its repository with a module path that reflected its published location " +
		"where Go tools could find it to download it For now because you have not published the module yet " +
		"you need to adapt the example com hello module so it can find the example com greetings code on your " +
		"local file system Go is a new language Although it borrows ideas from existing languages " +
		"it has unusual properties that make effective Go programs different in character " +
		"from programs written in its relatives A straightforward translation of a C or Java program into Go " +
		"is unlikely to produce a satisfactory result Java programs are written in Java not Go " +
		"On the other hand thinking about the problem from a Go perspective could produce a " +
		"successful but quite different program In other words to write Go well it is important " +
		"to understand its properties and idioms It is also important to know the established " +
		"conventions for programming in Go such as naming formatting program construction and so on " +
		"so that programs you write will be easy for other Go programmers to understand win window cat"

	words := strings.Split(s, " ")

	dict := make([][]interface{}, 0)

	for _, word := range words {
		w := make([]interface{}, 0)

		for i := 0; i < len(word); i++ {
			w = append(w, string(word[i]))
		}

		dict = append(dict, w)
	}

	return dict
}

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

	if trie.Exists([]interface{}{"w", "i", "n", "d", "o", "w"}) {
		t.Fatal("window word must be removed")
	}
}

func TestFlatten(t *testing.T) {
	trie := testTrie()

	dict := trie.Flatten()

	if len(dict) != 130 {
		t.Fatal("must have 130 words")
	}
}

func BenchmarkExists(b *testing.B) {
	trie := testTrie()

	for n := 0; n < b.N; n++ {
		trie.Exists([]interface{}{"u", "n", "d", "e", "r", "s", "t", "a", "n", "d"})
	}
}

func BenchmarkExistsPlainDict(b *testing.B) {
	dict := testDict()

	needle := []interface{}{"u", "n", "d", "e", "r", "s", "t", "a", "n", "d"}

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
