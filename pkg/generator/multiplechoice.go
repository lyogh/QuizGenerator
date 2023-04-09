package generator

import (
	"github.com/lyogh/QuizGenerator/pkg/card"
)

type multipleChoiceCard struct {
	generator
}

func NewMultipleChoiceCard(parameters *Parameters) Generator {
	p := *parameters

	p.SetTypes([]card.CardType{card.TypeMultipleChoice})

	return &multipleChoiceCard{
		generator: *NewGenerator(&p).(*generator),
	}
}

func (g *multipleChoiceCard) CreateCards() error {
	for i := 0; i < int(g.parameters.CardsMax()); i++ {
		c, err := card.NewCard(card.QuestionTextChooseStatements, card.TypeMultipleChoice)
		if err != nil {
			return err
		}

		if err := c.GenerateOptionsByFacts(g.facts, g.distractors, *g.parameters.CardParameters()); err != nil {
			return err
		}

		g.cards = append(g.cards, c)
	}

	return nil
}
