package option

import (
	"unicode"

	"github.com/lyogh/QuizGenerator/internal/slice"
	"github.com/lyogh/QuizGenerator/pkg/types"
)

type Options []Option

var _ types.Shuffler = (Options)(nil)

func (o Options) Shuffle() {
	slice.Shuffle(o)
	o.Beautify()
}

func (o Options) Beautify() {
	const (
		first = 65
	)

	for i, opt := range o {
		opt.SetId(uint(i + 1))
		opt.SetSymbol(unicode.ToUpper(rune(first + i)))
	}
}
