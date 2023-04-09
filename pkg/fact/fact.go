package fact

type Fact struct {
	object     string
	statements Statements
	stmMap     map[Statement]struct{}
}

type Facts []*Fact

func NewFact(object string, stms Statements) *Fact {
	fact := &Fact{
		object:     object,
		statements: stms,
	}

	fact.stmMap = make(map[Statement]struct{})

	for _, stm := range fact.statements {
		fact.stmMap[*stm] = struct{}{}
	}

	return fact
}

func (f *Fact) Object() string {
	return f.object
}

func (f *Fact) Statements() Statements {
	return f.statements
}

func (f *Fact) HasStatement(stm Statement) bool {
	_, ok := f.stmMap[stm]
	return ok
}
