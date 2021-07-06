package RGTree

import (
	"fmt"
	"testing"
)

type ItemTest struct {
	ID int
	ParentId int
	Parent RGTree
}

func (it *ItemTest) SetParent(t RGTree) {
	it.Parent = t
}

func (it ItemTest) GetParent() RGTree {
	return it.Parent
}

func (it ItemTest) GetID() int {
	return it.ID
}

func (it ItemTest) GetParentID() int {
	return it.ParentId
}

func TestMakeTreeMap(t *testing.T) {
	var plain []RGTree

	root := &ItemTest{ID: 1}

	plain = append(plain, root)
	plain = append(plain, &ItemTest{ID: 2, ParentId: 1})
	plain = append(plain, &ItemTest{ID: 3, ParentId: 1})
	plain = append(plain, &ItemTest{ID: 4, ParentId: 2})
	plain = append(plain, &ItemTest{ID: 5, ParentId: 2})
	plain = append(plain, &ItemTest{ID: 6, ParentId: 2})
	plain = append(plain, &ItemTest{ID: 7, ParentId: 3})
	plain = append(plain, &ItemTest{ID: 8, ParentId: 3})
	plain = append(plain, &ItemTest{ID: 9, ParentId: 4})
	plain = append(plain, &ItemTest{ID: 10, ParentId: 4})

	root.SetParent(&ItemTest{ID: 10, ParentId: 4})

	fmt.Println(root.Parent.GetID())

	fmt.Println(MakeTree(root, plain))
}
