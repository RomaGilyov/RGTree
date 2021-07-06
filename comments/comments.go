package comments

type Comment interface {
	SetChild(Comment)
	GetChildren() []Comment
	GetID() int
	GetParentID() int
}

func MakeTree(root Comment, comments []Comment) Comment {
	memo := make(map[int]bool, len(comments))

	memo[root.GetID()] = true

	return makeTreeUtil(root, comments, memo)
}

func makeTreeUtil(root Comment, comments []Comment, memo map[int]bool) Comment {
	var children []Comment

	for _, child := range comments {
		if child.GetParentID() == root.GetID() && memo[child.GetID()] != false {
			memo[child.GetID()] = true

			children = append(children, child)
		}
	}

	for _, child := range children {
		root.SetChild(makeTreeUtil(child, comments, memo))
	}

	return root
}

func Traverse(root *Comment, handler func(*Comment)) {
	memo := make(map[int]bool, 10)

	traverseUtil(root, handler, memo)
}

func traverseUtil(root *Comment, handler func(*Comment), memo map[int]bool) {
	if memo[(*root).GetID()] == true {
		return
	}

	handler(root)

	memo[(*root).GetID()] = true

	for _, child := range (*root).GetChildren() {
		Traverse(&child, handler)
	}
}
