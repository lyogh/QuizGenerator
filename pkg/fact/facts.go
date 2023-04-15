package fact

import (
	"github.com/lyogh/QuizGenerator/internal/slice"
)

type Facts []*Fact

/*
Удаляет факт
*/
func (f *Facts) Delete(i int) {
	copy((*f)[i:], (*f)[i+1:])
	slice.Shrink(f, len(*f)-1)
}
