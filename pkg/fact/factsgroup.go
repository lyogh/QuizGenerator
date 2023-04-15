package fact

type group struct {
	facts *Facts
	lies  *Facts
}

// Группа объединяющая факты и дистракторы
type Group interface {
	AddFact(*Fact)
	AddFacts(*Facts)
	Facts() *Facts

	AddLie(*Fact)
	AddLies(*Facts)
	Lies() *Facts

	Clone() Group
}

func NewGroup() Group {
	return &group{
		facts: new(Facts),
		lies:  new(Facts),
	}
}

func (g *group) AddFact(fact *Fact) {
	*g.facts = append(*g.facts, fact)
}

func (g *group) AddFacts(facts *Facts) {
	for _, f := range *facts {
		g.AddFact(f)
	}
}

func (g *group) Facts() *Facts {
	return g.facts
}

func (g *group) AddLie(lie *Fact) {
	*g.lies = append(*g.lies, lie)
}

func (g *group) AddLies(lies *Facts) {
	for _, l := range *lies {
		g.AddLie(l)
	}
}

func (g *group) Lies() *Facts {
	return g.lies
}

func (g *group) Clone() Group {
	c := NewGroup()

	c.AddFacts(g.Facts())
	c.AddLies(g.Lies())

	return c
}
