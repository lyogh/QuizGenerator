package generator

import (
	"errors"

	"github.com/lyogh/QuizGenerator/pkg/card"
)

const defaultCardsMax = 40

type Parameters struct {
	// Типы карточек с вопросами
	types card.CardTypes
	// Максимальное количество вопросов
	maxCards uint
	// Параметры генерации карточки вопроса
	card *card.Parameters
}

var DefaultParameters = NewParameters() // Параметры по умолчанию

func NewParameters() *Parameters {
	p := &Parameters{
		maxCards: defaultCardsMax,
		types: []card.CardType{
			card.TypeMultipleChoice,
			card.TypeTrueFalse,
		},
		card: card.NewParameters(),
	}

	return p
}

func (p *Parameters) Types() card.CardTypes {
	return p.types
}

func (p *Parameters) SetTypes(ct card.CardTypes) {
	p.types = ct
}

func (p *Parameters) CardsMax() uint {
	return p.maxCards
}

func (p *Parameters) SetCardsMax(v uint) error {
	if v < 1 {
		return errors.New("количество вопросов должно быть > 0")
	}

	p.maxCards = v

	return nil
}

func (p *Parameters) CardParameters() *card.Parameters {
	return p.card
}

func (p *Parameters) SetCardParameters(cp *card.Parameters) {
	p.card = cp
}
