package fact

import "github.com/lyogh/QuizGenerator/internal/slice"

type Statements []*Statement

/*
Удаляет утверждение
*/
func (s *Statements) Delete(i int) {
	copy((*s)[i:], (*s)[i+1:])
	slice.Shrink(s, len(*s)-1)
}
