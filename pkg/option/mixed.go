package option

import (
	"fmt"
	"strings"

	"github.com/lyogh/QuizGenerator/internal/slice"
)

const (
	AnswerTrueList = "Правильные ответы"
	AnswerAllTrue  = "Все ответы правильные"
)

type mixedOption struct {
	option

	// Объединяет все возможные варианты ответов на вопрос ?
	all bool
	// Вложенные варианты ответов
	options Options
}

type MixedOption interface {
	Option
	OptionsManager
}

func NewMixedOption(options Options, correct bool, all bool) MixedOption {
	mo := &mixedOption{
		all:     all,
		options: options,
	}

	mo.option = *NewOption(mo.String(), correct).(*option)

	return mo
}

func (o *mixedOption) AddOption(option Option) {
	o.options = append(o.options, option)
}

func (o *mixedOption) DeleteOption(option Option) {
	o.options = slice.Delete(o.options, option)
}

func (o *mixedOption) Options() Options {
	return o.options
}

func (o *mixedOption) String() string {
	var (
		s     strings.Builder
		descr string
	)

	if o.all {
		descr = AnswerAllTrue
	} else {
		for i, o := range o.options {
			if i > 0 {
				s.WriteString(", ")
			}

			s.WriteString(string(o.Symbol()))
		}

		descr = fmt.Sprintf("%s: %s", AnswerTrueList, s.String())
	}

	return descr
}
