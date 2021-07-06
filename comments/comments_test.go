package comments

import (
	"fmt"
	"testing"
)

type CommentTest struct {
	ID int
	ParentId int
	Children []Comment
}

func (ct *CommentTest) SetChild(c Comment) {
	ct.Children = append(ct.Children, c)
}

func (ct CommentTest) GetChildren() []Comment {
	return ct.Children
}

func (ct CommentTest) GetID() int {
	return ct.ID
}

func (ct CommentTest) GetParentID() int {
	return ct.ParentId
}

func TestMakeTreeMap(t *testing.T) {
	var plain []Comment

	root := &CommentTest{ID: 1}

	plain = append(plain, root)
	plain = append(plain, &CommentTest{ID: 2, ParentId: 1})
	plain = append(plain, &CommentTest{ID: 3, ParentId: 1})
	plain = append(plain, &CommentTest{ID: 4, ParentId: 2})
	plain = append(plain, &CommentTest{ID: 5, ParentId: 2})
	plain = append(plain, &CommentTest{ID: 6, ParentId: 2})
	plain = append(plain, &CommentTest{ID: 7, ParentId: 3})
	plain = append(plain, &CommentTest{ID: 8, ParentId: 3})
	plain = append(plain, &CommentTest{ID: 9, ParentId: 4})
	plain = append(plain, &CommentTest{ID: 10, ParentId: 4})

	fmt.Println(MakeTree(root, plain))
}

