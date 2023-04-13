package card

import (
	"errors"
)

const (
	MinAnswers = 1
	MaxAnswers = 10
	MinOptions = MinAnswers + 1
	MaxOptions = MaxAnswers
)

type Parameters struct {
	// Количество всех вариантов ответов
	options *Interval
	// Количество правильных ответов
	answers *Interval
}

var (
	DefaultParameters = NewParameters()   // Параметры по умолчанию
	AikenParameters   = aikenParameters() // Параметры карточки вопроса для последующего экспорта в формат Aiken
)

var ErrOptionsAnswersMax = errors.New("количество правильных ответов не может быть больше общего количества вариантов ответов")

func NewParameters() *Parameters {
	p := &Parameters{
		options: NewInterval(MinOptions, MaxOptions),
		answers: NewInterval(MinAnswers, MaxAnswers),
	}

	return p
}

func aikenParameters() *Parameters {
	p := NewParameters()

	p.options.max = 10
	p.answers.max = 1

	return p
}

func (p *Parameters) Options() *Interval {
	return p.options
}

func (p *Parameters) SetOptions(v *Interval) error {
	if v.min < 2 {
		return errors.New("количество вариантов ответов не может быть меньше двух")
	}

	if v.max < p.Answers().max {
		return ErrOptionsAnswersMax
	}

	p.options = v

	return nil
}

func (p *Parameters) Answers() *Interval {
	return p.answers
}

func (p *Parameters) SetAnswers(v *Interval) error {
	if v.min < 1 {
		return errors.New("количество правильных ответов не может быть меньше одного")
	}

	if v.max > p.Options().max {
		return ErrOptionsAnswersMax
	}

	p.answers = v

	return nil
}
