package genealogy

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
	memo := make(map[int]bool, len(ancestors))

	mapped := make(map[int]Genealogy, len(ancestors))

	for _, ancestor := range ancestors {
		mapped[ancestor.GetID()] = ancestor
	}

	return makeGenealogyUtil(root, mapped, memo)
}

func makeGenealogyUtil(root Genealogy, mapped map[int]Genealogy, memo map[int]bool) error {
	if memo[root.GetID()] == true {

	}
}
