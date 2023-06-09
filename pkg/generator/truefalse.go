package generator

import (
	"fmt"

	"github.com/lyogh/QuizGenerator/pkg/card"
	"github.com/lyogh/QuizGenerator/pkg/fact"
	"github.com/lyogh/QuizGenerator/pkg/option"
)

type trueFalseGenerator struct {
	generator
}

/*
Создает новый объект генератора карточек вопросов True/False
*/
func NewTrueFalseGenerator(params *Parameters) Generator {
	p := *params

	p.SetTypes([]card.CardType{card.TypeTrueFalse})

	return &trueFalseGenerator{
		generator: *NewGenerator(&p).(*generator),
	}
}

func (g *trueFalseGenerator) CreateCards() error {
	// Верные утверждения
	for i, f := range *g.groups.Data()[fact.RootKey].Facts() {
		if i > int(g.parameters.CardsMax()) {
			break
		}

		for _, stm := range *f.Statements() {
			if err := g.addCard(f, *stm, true); err != nil {
				return err
			}
		}
	}

	// Ошибочные утверждения
	for i, dis := range *g.groups.Data()[fact.RootKey].Lies() {
		if i > int(g.parameters.CardsMax()) {
			break
		}

		for _, stm := range *dis.Statements() {
			if err := g.addCard(dis, *stm, false); err != nil {
				return err
			}
		}
	}

	g.Shuffle()

	return nil
}

func (g *trueFalseGenerator) addCard(f *fact.Fact, stm fact.Statement, answer bool) error {
	c, err := card.NewCard(
		fmt.Sprintf("%s: %s - %s", card.QuestionTextStatementCorrect, f.Object(), stm), card.TypeTrueFalse)
	if err != nil {
		return err
	}

	c.AddOption(option.NewOption(card.AnswerYes, answer))
	c.AddOption(option.NewOption(card.AnswerNo, !answer))

	c.Options().Beautify()

	g.cards = append(g.cards, c)

	return nil
}
