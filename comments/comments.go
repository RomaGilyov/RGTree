package comments

import (
	"errors"
	"github.com/RomaGilyov/RGTree/util"
)

type Comment interface {
	SetChild(Comment)
	GetChildren() []Comment
	GetID() interface{}
	GetParentID() interface{}
}

func MakeTree(root Comment, comments []Comment) error {
	memo := make(map[interface{}]bool, len(comments))

	memo[root.GetID()] = true

	return makeTreeUtil(root, comments, memo)
}

func makeTreeUtil(root Comment, comments []Comment, memo map[interface{}]bool) error {
	var children []Comment

	var err error

	for _, child := range comments {
		if child.GetParentID() == root.GetID() {
			if memo[child.GetID()] == true {
				errMes := "Recursive data: " + util.NumericToString(child.GetID()) + " and " + util.NumericToString(root.GetID())

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
	memo := make(map[interface{}]bool, 10)

	return traverseUtil(root, handler, memo)
}

func traverseUtil(root Comment, handler func(Comment) bool, memo map[interface{}]bool) error {
	if memo[root.GetID()] == true {
		return errors.New("there are duplicate primary ID elements of value: " + util.NumericToString(root.GetID()))
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
