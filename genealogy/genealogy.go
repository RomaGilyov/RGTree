package genealogy

import (
	"errors"
	"strconv"
)

type Genealogy interface {
	GetID() int
	GetMotherID() int
	GetFatherID() int
	SetFather(Genealogy)
	SetMother(Genealogy)
	GetFather() Genealogy
	GetMother() Genealogy
}

func MakeGenealogy(root Genealogy, ancestors []Genealogy) error {
	memo := make(map[int]bool, 1)

	memo[root.GetID()] = true

	mapped := make(map[int]Genealogy, len(ancestors))

	for _, ancestor := range ancestors {
		mapped[ancestor.GetID()] = ancestor
	}

	return makeGenealogyUtil(root, mapped, memo)
}

/*
	Infinite recursion example:

				 5[3, 4]
				/\
			   /  \
			  3    4[8, 6]
			 /\    /\
			/  \  /  \
		   /    \/	  \
		  1    2  8    6[nil, 5] -> Grandfather is a child of the 5th grandchild
		 /\       /\
		/  \     /  \
       7   2    9    2
			   /\
			  /  \
			 x	  2

	Two important things to note:
		1. There are can be the same ancestors for example 3 and 8 has the same ancestor 2
		2. The infinite recursion happens when there is a cyclic relation like leaf 6 that points to child 5
			that would come back to node 5 and go down to 4 and then to 6 and so on...

	5 -> 4 -> 6 -> 5
 */
func makeGenealogyUtil(root Genealogy, mapped map[int]Genealogy, memo map[int]bool) error {
	if memo[root.GetFatherID()] == true {
		return errors.New("infinite ref: " + strconv.Itoa(root.GetID()) + " <-> " + strconv.Itoa(root.GetFatherID()))
	}

	if memo[root.GetMotherID()] == true {
		return errors.New("infinite ref: " + strconv.Itoa(root.GetID()) + " <-> " + strconv.Itoa(root.GetMotherID()))
	}

	if mother, exists := mapped[root.GetMotherID()]; exists {
		motherMemo := appendMemo(memo, mother.GetID())

		root.SetMother(mother)

		err := makeGenealogyUtil(mother, mapped, motherMemo)

		if err != nil {
			return err
		}
	}

	if father, exists := mapped[root.GetFatherID()]; exists {
		fatherMemo := appendMemo(memo, father.GetID())

		root.SetFather(father)

		err := makeGenealogyUtil(father, mapped, fatherMemo)

		if err != nil {
			return err
		}
	}

	return nil
}

func appendMemo(memo map[int]bool, id int) map[int]bool {
	newPath := make(map[int]bool, len(memo) + 1)

	for pid, exists := range memo {
		newPath[pid] = exists
	}

	newPath[id] = true

	return newPath
}

func Traverse(root Genealogy, handler func(Genealogy)) {
	handler(root)

	if root.GetFather() != nil {
		Traverse(root.GetFather(), handler)
	}

	if root.GetMother() != nil {
		Traverse(root.GetMother(), handler)
	}
}
