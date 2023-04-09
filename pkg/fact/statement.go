package fact

type (
	Statement  string
	Statements []*Statement
)

func NewStatement(stm string) *Statement {
	s := new(Statement)

	*s = Statement(stm)

	return s
}

func (s Statement) String() string {
	return string(s)
}
