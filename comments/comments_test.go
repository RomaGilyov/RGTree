package comments

import (
	"strconv"
	"testing"
)

////////////////////////////////////////////////////// Testing data //////////////////////////////////////

type CommentTest struct {
	Title string
	ID int
	ParentID int
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
	return ct.ParentID
}

func (ct CommentTest) String() string {
	return "{ID: " + strconv.Itoa(ct.ID) + ", ParentID: " + strconv.Itoa(ct.ParentID) + ", Title: " + ct.Title + "}"
}

/*
	There are test comments with this structure:

		1 ->
			2 ->
				4
			3 ->
				5
 */
func testData() (Comment, []Comment) {
	var plain []Comment

	root := &CommentTest{ID: 1, Title: "Test 1"}

	plain = append(plain, root)
	plain = append(plain, &CommentTest{ID: 2, ParentID: 1, Title: "Test 2"})
	plain = append(plain, &CommentTest{ID: 3, ParentID: 1, Title: "Test 3"})
	plain = append(plain, &CommentTest{ID: 4, ParentID: 2, Title: "Test 4"})
	plain = append(plain, &CommentTest{ID: 5, ParentID: 3, Title: "Test 5"})

	return root, plain
}

////////////////////////////////////////////////////// Testing /////////////////////////////////////////

func TestMakeTreeMap(t *testing.T) {
	root, plain := testData()

	MakeTree(root, plain)

	if len(root.GetChildren()) != 2 {
		t.Fatalf("%s must have 2 children", root)
	}

	if root.GetChildren()[0].GetChildren()[0].GetID() != 4 {
		t.Fatalf("%s ->children -> 0 -> children -> 0 -> ID must be 4", root)
	}

	if root.GetChildren()[1].GetChildren()[0].GetID() != 5 {
		t.Fatalf("%s ->children -> 1 -> children -> 0 -> ID must be 5", root)
	}
}

func TestMakeTreeMapRecursiveError(t *testing.T) {
	root, plain := testData()

	ct := root.(*CommentTest)

	ct.ParentID = 3 // ID reference to child

	err := MakeTree(root, plain)

	if err == nil {
		t.Fatalf("Infinite recurtion must return an error")
	}
}

func TestTraverse(t *testing.T) {
	root, plain := testData()

	var TestTraverseIterations int

	MakeTree(root, plain)

	err := Traverse(root, func (comment Comment) {
		if ct, ok := comment.(*CommentTest); ok {
			ct.Title = "Test " + strconv.Itoa(ct.GetID()) + " Modified"
		}

		TestTraverseIterations++
	})

	if TestTraverseIterations != 5 {
		t.Fatal("There are 5 elements", TestTraverseIterations, "given")
	}

	if err != nil {
		t.Fatal("There are no duplicates but there duplicate error")
	}

	root.GetChildren()[0].SetChild(&CommentTest{ID: 2}) // Check for doubles

	err = Traverse(root, func (comment Comment) {})

	if err == nil {
		t.Fatal("Traverse function must give error when a duplicate encountered")
	}
}
