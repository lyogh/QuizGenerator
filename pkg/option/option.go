package option

import (
	"fmt"

	"github.com/lyogh/QuizGenerator/pkg/types"
)

type Option interface {
	fmt.Stringer
	types.IdSetter
	types.IdGetter

	SetCorrect(bool)
	Correct() bool

	SetSymbol(rune)
	Symbol() rune
}

type OptionsManager interface {
	Options() Options
	AddOption(Option)
	DeleteOption(Option)
}

type option struct {
	// Ид выбора
	id uint
	// Обозначение выбора
	symbol rune
	// Описание выбора
	description string
	// Признак правильного ответа
	correct bool
}

func NewOption(descr string, correct bool) Option {
	option := &option{
		description: descr,
		correct:     correct,
	}

	return option
}

func (o *option) String() string {
	return o.description
}

func (o *option) SetId(id uint) {
	o.id = id
}

func (o *option) Id() uint {
	return o.id
}

func (o *option) SetSymbol(s rune) {
	o.symbol = s
}

func (o *option) Symbol() rune {
	return o.symbol
}

func (o *option) SetCorrect(ca bool) {
	o.correct = ca
}

func (o *option) Correct() bool {
	return o.correct
}
