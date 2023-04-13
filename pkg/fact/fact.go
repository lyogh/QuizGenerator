package fact

// Факт
type Fact struct {
	object     string     // Название объекта
	statements Statements // Утверждения (правдивые факты об объекте)

	stmMap map[Statement]struct{}
}

/*
Создает новый объект - факт
*/
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

/*
Возвращает название объекта
*/
func (f *Fact) Object() string {
	return f.object
}

/*
Возвращает правдивые факты об объекте
*/
func (f *Fact) Statements() *Statements {
	return &f.statements
}

/*
Возвращает признак существования утверждения об объекте
*/
func (f *Fact) HasStatement(stm Statement) bool {
	_, ok := f.stmMap[stm]
	return ok
}
