package comments

import (
	"errors"
	"strconv"
)

type Comment interface {
	SetChild(Comment)
	GetChildren() []Comment
	GetID() int
	GetParentID() int
}

func MakeTree(root Comment, comments []Comment) error {
	memo := make(map[int]bool, len(comments))

	memo[root.GetID()] = true

	return makeTreeUtil(root, comments, memo)
}

func makeTreeUtil(root Comment, comments []Comment, memo map[int]bool) error {
	var children []Comment

	var err error

	for _, child := range comments {
		if child.GetParentID() == root.GetID() {
			if memo[child.GetID()] == true {
				errMes := "Recursive data: " + strconv.Itoa(child.GetID()) + " and " + strconv.Itoa(root.GetID())

				return errors.New(errMes)
			}

			memo[child.GetID()] = true

			children = append(children, child)
		}
	}

	for _, child := range children {
		err = makeTreeUtil(child, comments, memo)

		if err != nil {
			return err
		}

		root.SetChild(child)
	}

	return nil
}

func Traverse(root Comment, handler func(Comment) bool) error {
	memo := make(map[int]bool, 10)

	return traverseUtil(root, handler, memo)
}

func traverseUtil(root Comment, handler func(Comment) bool, memo map[int]bool) error {
	if memo[root.GetID()] == true {
		return errors.New("there are duplicate primary ID elements of value: " + strconv.Itoa(root.GetID()))
	}

	if handler(root) == false {
		return nil
	}

	memo[root.GetID()] = true

	for _, child := range root.GetChildren() {
		err := traverseUtil(child, handler, memo)

		if err != nil {
			return err
		}
	}

	return nil
}
