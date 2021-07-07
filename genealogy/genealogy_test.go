package genealogy

import (
	"strconv"
	"testing"
)

////////////////////////////////////////////////////// Testing data //////////////////////////////////////

type GenealogyTest struct {
	Name string
	ID int
	FatherID int
	MotherID int
	Mother Genealogy
	Father Genealogy
}

func (ct *GenealogyTest) SetFather(f Genealogy) {
	ct.Father = f
}

func (ct *GenealogyTest) SetMother(m Genealogy) {
	ct.Mother = m
}

func (ct GenealogyTest) GetFather() Genealogy {
	return ct.Father
}

func (ct GenealogyTest) GetMother() Genealogy {
	return ct.Mother
}

func (ct GenealogyTest) GetID() int {
	return ct.ID
}

func (ct GenealogyTest) GetFatherID() int {
	return ct.FatherID
}

func (ct GenealogyTest) GetMotherID() int {
	return ct.MotherID
}

func (ct GenealogyTest) String() string {
	return "{" +
		"ID: " + strconv.Itoa(ct.ID) + ", " +
		"FatherID: " + strconv.Itoa(ct.FatherID) + ", " +
		"MotherID: " + strconv.Itoa(ct.MotherID) + ", " +
		"Name: " + ct.Name + "}"
}

/*
				5[3, 4]
				/\
			   /  \
			  3    4[8, 6]
			 /\    /\
			/  \  /  \
		   /    \/	  \
		  1    2  8    6[nil, 5] -> inf ref
		 /\       /\
		/  \     /  \
       7   2    9    2
			   /\
			  /  \
			 x	  2
*/
func testData(infRef bool) (Genealogy, []Genealogy) {
	var plain []Genealogy

	var infRefAncestor Genealogy

	root := &GenealogyTest{ID: 5, Name: "Test 5", MotherID: 3, FatherID: 4}

	if infRef {
		infRefAncestor = &GenealogyTest{ID: 6, Name: "Test 6", FatherID: 5}
	} else {
		infRefAncestor = &GenealogyTest{ID: 6, Name: "Test 6"}
	}

	plain = append(plain, root)
	plain = append(plain, &GenealogyTest{ID: 3, Name: "Test 3", MotherID: 1, FatherID: 2})
	plain = append(plain, &GenealogyTest{ID: 1, Name: "Test 1", MotherID: 7, FatherID: 2})
	plain = append(plain, &GenealogyTest{ID: 7, Name: "Test 7"})
	plain = append(plain, &GenealogyTest{ID: 2, Name: "Test 2"})
	plain = append(plain, &GenealogyTest{ID: 4, Name: "Test 4", MotherID: 8, FatherID: 6})
	plain = append(plain, &GenealogyTest{ID: 8, Name: "Test 8", MotherID: 9, FatherID: 2})
	plain = append(plain, &GenealogyTest{ID: 9, Name: "Test 9", FatherID: 2})
	plain = append(plain, infRefAncestor)

	return root, plain
}

////////////////////////////////////////////////////// Testing /////////////////////////////////////////

func TestMakeTreeMap(t *testing.T) {
	root, plain := testData(false)

	err := MakeGenealogy(root, plain)

	if err != nil || root.GetFather().GetID() != 4 || root.GetFather().GetMother().GetID() != 8 {
		t.Fatal("Genealogy maker failed")
	}
}

func TestMakeTreeMapRecursiveError(t *testing.T) {
	root, plain := testData(true)

	err := MakeGenealogy(root, plain)

	if err != nil {
		t.Fatal("Infinite reference must cause an error!")
	}
}

func TestTraverse(t *testing.T) {
	root, plain := testData(false)

	MakeGenealogy(root, plain)

	var iterations int

	Traverse(root, func (item Genealogy) {
		iterations++
	})

	if iterations != 12 {
		t.Fatal("There are 12 nodes in the tree, only " + strconv.Itoa(iterations) + " made")
	}
}