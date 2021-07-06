package RGTree

type RGTree interface {
	GetParent () RGTree
	SetParent (RGTree)
	GetID () int
	GetParentID () int
}

func MakeTree(root RGTree, items []RGTree) RGTree {
	mapped := make(map[int]RGTree, len(items))

	for _, tree := range items {
		mapped[tree.GetID()] = tree
	}

	return MakeTreeUtil(root, mapped)
}

func MakeTreeUtil(root RGTree, mapped map[int]RGTree) RGTree {
	parent, exists := mapped[root.GetParentID()]

	if exists {
		child := MakeTreeUtil(parent, mapped)

		root.SetParent(child)
	}

	return root
}
