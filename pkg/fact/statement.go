package fact

type Statement string

/*
Создает новый объект - утверждение
*/
func NewStatement(stm string) *Statement {
	s := new(Statement)

	*s = Statement(stm)

	return s
}

func (s Statement) String() string {
	return string(s)
}
