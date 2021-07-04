package RGTree

import "fmt"

type RGTree interface {
	GetParent () RGTree
	SetParent (RGTree) RGTree
	GetID () int
	GetParentID () int
}

func MakeTree(root RGTree, items []RGTree) RGTree {
	mapped := make(map[int]RGTree, len(items))

	for _, tree := range items {
		fmt.Println(tree)
		mapped[tree.GetID()] = tree
	}

	return MakeTreeUtil(root, mapped)
}

func MakeTreeUtil(root RGTree, mapped map[int]RGTree) RGTree {
	parent, exists := mapped[root.GetParentID()]

	if exists {
		root.SetParent(MakeTreeUtil(parent, mapped))
	}

	return root
}
