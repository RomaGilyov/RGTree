package RGTree

import (
	"fmt"
	"testing"
)

type ItemTest struct {
	RGTree
	ID int
	ParentId int
}

func (it ItemTest) GetId() int {
	return it.ID
}

func (it ItemTest) GetParentId() int {
	return it.ParentId
}

func TestMakeTreeMap(t *testing.T) {
	var plain []RGTree

	root := ItemTest{ID: 1}

	plain = append(plain, root)
	plain = append(plain, ItemTest{ID: 2, ParentId: 1})
	plain = append(plain, ItemTest{ID: 3, ParentId: 1})
	plain = append(plain, ItemTest{ID: 4, ParentId: 2})
	plain = append(plain, ItemTest{ID: 5, ParentId: 2})
	plain = append(plain, ItemTest{ID: 6, ParentId: 2})
	plain = append(plain, ItemTest{ID: 7, ParentId: 3})
	plain = append(plain, ItemTest{ID: 8, ParentId: 3})
	plain = append(plain, ItemTest{ID: 9, ParentId: 4})
	plain = append(plain, ItemTest{ID: 10, ParentId: 4})

	fmt.Println(MakeTree(root, plain))
}
