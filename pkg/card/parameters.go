package card

import "errors"

const (
	MinAnswers = 1
	MinOptions = 2
)

type Parameters struct {
	// Количество всех вариантов ответов
	options *Interval
	// Количество правильных ответов
	answers *Interval
}

// Параметры карточки вопроса для последующего экспорта в формат Aiken
var AikenParameters = aikenParameters()

func NewParameters() *Parameters {
	p := &Parameters{
		options: NewInterval(MinOptions, ^uint(0)),
		answers: NewInterval(MinAnswers, ^uint(0)),
	}

	return p
}

func (p *Parameters) Options() *Interval {
	return p.options
}

func (p *Parameters) SetOptions(v *Interval) error {
	if v.min < 2 {
		return errors.New("количество вариантов ответов не может быть меньше двух")
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

	p.answers = v

	return nil
}
