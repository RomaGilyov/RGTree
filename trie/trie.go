package trie

import (
	"errors"
)

/*
	Trie
 */
type Trie struct {
	Val interface{}
	Children map[interface{}]*Trie
}

func Construct() *Trie {
	return &Trie{Val: nil, Children: make(map[interface{}]*Trie)}
}

func (t *Trie) PrefixSearch(prefix []interface{}) [][]interface{} {
	results := make([][]interface{}, 0, 10)

	suffix, err := t.GetSuffix(prefix)

	if err != nil {
		return results
	}

	results = suffix.Flatten()

	for i, col := range results {
		results[i] = append(prefix, col...)
	}

	return results
}

/*
			1
			/\
		   /  \
		  2    3
		 /\    /\
		/  \  /  \
	   /    \/	  \
	  4    5  6    7

	1

	1 2
	1 3

	1 2 4
	1 2 5
	1 3 6
	1 3 7
 */
func (t *Trie) Flatten() [][]interface{} {
	results := make([][]interface{}, 0, 10)

	t.trieToDictUtil(make([]interface{}, 0, 10), &results)

	return results
}

func (t *Trie) trieToDictUtil(path []interface{}, results *[][]interface{}) {
	if len(t.Children) == 0 {
		clonePath := make([]interface{}, 0)

		for _, v := range path {
			clonePath = append(clonePath, v)
		}

		*results = append(*results, clonePath)
	} else {
		for _, child := range t.Children {
			child.trieToDictUtil(append(path, child.Val), results)
		}
	}
}

func (t *Trie) GetSuffix(values []interface{}) (*Trie, error) {
	if len(values) == 0 {
		return t, nil
	}

	val := values[0]

	if child, ok := t.Children[val]; ok {
		return child.GetSuffix(values[1:])
	}

	return nil, errors.New("suffix does not exists")
}

func (t *Trie) Exists(values []interface{}) bool {
	return t.existsUtil(values, 0)
}

func (t *Trie) existsUtil(values []interface{}, index int) bool {
	if index == len(values) {
		return true
	}

	current := values[index]

	if child, ok := t.Children[current]; ok {
		return child.existsUtil(values, index + 1)
	}

	return false
}

func (t *Trie) Add(values []interface{}) {
	if len(values) == 0 {
		return
	}

	val := values[0]

	if child, ok := t.Children[val]; ok {
		child.Add(values[1:])
	} else {
		t.Children[val] = &Trie{Val: val, Children: make(map[interface{}]*Trie)}

		t.Children[val].Add(values[1:])
	}
}

func (t *Trie) Remove(values []interface{}) {
	if len(values) == 0 {
		return
	}

	suffixParent, pErr := t.GetSuffix(values[:len(values) - 1])

	if pErr != nil {
		return
	}

	suffix, err := t.GetSuffix(values)

	if err != nil {
		return
	}

	if len(suffix.Children) == 0 {
		delete(suffixParent.Children, suffix.Val)

		t.Remove(values[:len(values) - 1])
	}
}
