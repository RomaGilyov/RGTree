package trie

/*
	Trie
 */
type Trie struct {
	Val interface{}
	Children map[interface{}]*Trie
}

func Construct(values []interface{}) *Trie {
	t := new(Trie)

	t.Add(values)

	return t
}

func (t *Trie) PrefixSearch(prefix []interface{}) [][]interface{} {
	results := make([][]interface{}, 0, 10)

	if t.Exists(prefix) == false {
		return results
	}

	suffix := t.getSuffixUtil(prefix)

	results = TrieToDict(suffix)

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
func TrieToDict(node *Trie) [][]interface{} {
	return trieToDictUtil(node, make([]interface{}, 0, 10), make([][]interface{}, 0, 10))
}

func trieToDictUtil(node *Trie, path []interface{}, results [][]interface{}) [][]interface{} {
	if len(node.Children) == 0 {
		results = append(results, append(path, node.Val))
	} else {
		for _, child := range node.Children {
			trieToDictUtil(child, append(path, child.Val), results)
		}
	}

	return results
}

func (t *Trie) getSuffixUtil(values []interface{}) *Trie {
	if len(values) == 0 {
		return t
	}

	val := values[0]

	if child, ok := t.Children[val]; ok {
		return child.getSuffixUtil(values[1:])
	}

	panic("Exists but was not found")
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

	if t.Exists(values) == true {
		return
	}

	val := values[0]

	child := &Trie{Val: val}

	t.Children[val] = child

	child.Add(values[1:])
}
