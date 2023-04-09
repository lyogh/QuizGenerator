package card

import (
	"github.com/lyogh/QuizGenerator/internal/slice"
	"github.com/lyogh/QuizGenerator/pkg/types"
)

type Cards []Card

var _ types.Shuffler = (Cards)(nil)

func (c Cards) Shuffle() {
	slice.Shuffle(c)

	for i, card := range c {
		card.SetId(uint(i + 1))
	}
}
