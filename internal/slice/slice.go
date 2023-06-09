package slice

import (
	"math/rand"

	"golang.org/x/exp/slices"
)

/*
Инициализирует срез
*/
func Init[T any, S ~[]*T](s *S, zero bool) {
	if s == nil {
		*s = make(S, 0)
	} else {
		if zero {
			Shrink(s, 0)
		} else {
			*s = (*s)[:0]
		}
	}
}

/*
Перемешивает элементы среза
*/
func Shuffle[T any, S ~[]T](s S) {
	rand.Shuffle(len(s),
		func(i, j int) { s[i], s[j] = s[j], s[i] })
}

/*
Удаляет элемент из среза
*/
func Delete[T comparable, S ~[]T](s S, v T) S {
	i := slices.Index(s, v)
	if i < 0 {
		return s
	}

	return slices.Delete(s, i, i)
}

/*
Сокращает длину среза
*/
func Shrink[T any, S ~[]*T](s *S, l int) {
	if l >= len(*s) {
		return
	}

	for i := l; i < len(*s); i++ {
		(*s)[i] = nil
	}

	*s = (*s)[:l]
}
