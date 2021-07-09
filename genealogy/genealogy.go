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

/*
	Complexity O(2^n) Because worst case scenario is:

	Each tree level has the same ancestor for all children makes it
	1 = 1
	2 = 3
	3 = 7
	4 = 15
	5 = 31
	...

	F.e n = 3
				1
				/\
			   /  \
			  2    2
			 /\    /\
			/  \  /  \
		   /    \/	  \
		  3    3  3    3

	In case when all elements are ~unique O(n) (Genealogies usually have ~unique elements)

	F.e n = 7

				 1
				/\
			   /  \
			  2    3
			 /\    /\
			/  \  /  \
		   /    \/	  \
		  4    5  6    7

 */
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
		2. The infinite recursion happens when there is a cyclic reference f.e leaf 6 that points to child 5
			that would come back to node 5 and go down to 4 and then to 6 and so on...

	5 -> 4 -> 6 -> 5
 */
func makeGenealogyUtil(root Genealogy, mapped map[int]Genealogy, memo map[int]bool) error {
	if memo[root.GetFatherID()] == true {
		return errors.New("f self ref: " + strconv.Itoa(root.GetID()) + " <-> " + strconv.Itoa(root.GetFatherID()))
	}

	if memo[root.GetMotherID()] == true {
		return errors.New("m self ref: " + strconv.Itoa(root.GetID()) + " <-> " + strconv.Itoa(root.GetMotherID()))
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

/*
	Complexity O(n)
 */
func Traverse(root Genealogy, handler func(Genealogy, int) bool) {
	traverseUtil(root, 0, handler)
}

func traverseUtil(root Genealogy, level int, handler func(Genealogy, int) bool) {
	if handler(root, level) == false {
		return
	}

	level++

	if root.GetFather() != nil {
		traverseUtil(root.GetFather(), level, handler)
	}

	if root.GetMother() != nil {
		traverseUtil(root.GetMother(), level, handler)
	}
}

/*
	Example of traversing a genealogy up:

	F.e we have some entities and we need to calculate how a certain trait
	has transferred from ancestors to a certain child. F.e we have a trait X and we
	need to calculate what percentage of that trait has passed to a child and so on...

	For the sake of simplicity assume that X has 100% passing further:
	1. If ancestor A has the trait and B does not: 100% + 0%/2 = 50% passed further
	2. If ancestor A has the trait and B has the trait: 100% + 100%/2 = 100% passed further
	3. If ancestor A does not have the trait and B does not have the trait: 0% + 0%/2 = 0% passed further
	4. Assume a missing ancestor has 0% as X
	In this case the genealogy must have this values:

				 5 X transfer is 56.25%
				/\
			   /  \
			  3    4 [3]75% + [4]37.5%/2 ^
			 /\    /\
			/  \  /  \
		   /    \/	  \
		  1    2  8    6 [8]75% + [6]0%/2 ^
		 /\       /\
		/  \     /  \
       7   2    9    2  [9]50% + [2]100%/2 ^
			    \
			     \
			 	  2 [Gen X] -> [?] [2]100%/2

	Complexity O(2n)
*/
func TraverseUp(root Genealogy, handler func(Genealogy) bool) {
	levels := make(map[int][]Genealogy, 20)

	i := 0

	Traverse(root, func (g Genealogy, level int) bool {
		levels[level] = append(levels[level], g)

		if i < level {
			i = level
		}

		return true
	})

	for i >= 0 {
		for _, item := range levels[i] {
			if handler(item) == false {
				return
			}
		}

		i--
	}
}
