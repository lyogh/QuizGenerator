package generator

import (
	"fmt"
	"math/rand"

	"github.com/lyogh/QuizGenerator/internal/slice"
	"github.com/lyogh/QuizGenerator/pkg/card"
	"github.com/lyogh/QuizGenerator/pkg/fact"
	"github.com/lyogh/QuizGenerator/pkg/option"
)

type multipleChoiceCard struct {
	generator

	stmsMap map[*fact.Statements]*fact.Statements
}

/*
Создает новый объект генератора карточек множественного выбора
*/
func NewMultipleChoiceCard(parameters *Parameters) Generator {
	p := *parameters

	p.SetTypes([]card.CardType{card.TypeMultipleChoice})

	return &multipleChoiceCard{
		generator: *NewGenerator(&p).(*generator),
	}
}

func (g *multipleChoiceCard) CreateCards() error {
	slice.Shuffle(g.facts)
	slice.Shuffle(g.distractors)

	g.stmsMap = make(map[*fact.Statements]*fact.Statements)

	for i := 0; i < int(g.parameters.CardsMax()) && len(g.facts) > 0; i++ {
		if err := g.addCard(); err != nil {
			return err
		}
	}

	g.Shuffle()

	return nil
}

func (g *multipleChoiceCard) addCard() error {

	c, err := card.NewCard(card.QuestionTextChooseStatements, card.TypeMultipleChoice)
	if err != nil {
		return err
	}

	// Количество вариантов ответа
	olen := g.parameters.card.Options().RandomValue()

	// Количество правильных вариантов ответа
	alen := g.parameters.card.Answers().RandomValue()
	if alen > olen {
		alen = olen
	}

	// Добавляем правильные ответы
	for i := 0; i < alen && len(g.facts) > 0; i++ {
		g.addOptions(c, &g.facts, true)
	}

	for i := len(c.Answers()); i < olen && len(g.distractors) > 0; i++ {
		g.addOptions(c, &g.distractors, false)
	}

	if rand.Intn(2) == 0 && len(c.Options()) > card.MinOptions {
		// Группируемы ответы
		c.GroupOptions()
	}

	c.Options().Shuffle()

	g.cards = append(g.cards, c)

	return nil
}

/*
Добавляет варианты ответов в карточку с вопросом
*/
func (g *multipleChoiceCard) addOptions(c card.Card, facts *fact.Facts, correct bool) error {
	// Случайный объект
	fi := rand.Intn(len(*facts))
	f := (*facts)[fi]

	if _, ok := g.stmsMap[f.Statements()]; !ok {
		g.stmsMap[f.Statements()] = new(fact.Statements)
		*g.stmsMap[f.Statements()] = append(*g.stmsMap[f.Statements()], *(f.Statements())...)
	}

	// Случайное утверждение об объекте
	si := rand.Intn(len(*g.stmsMap[f.Statements()]))
	s := (*g.stmsMap[f.Statements()])[si]

	// Удаляем использованное утверждение
	(*g.stmsMap[f.Statements()]).Delete(si)

	if len((*g.stmsMap[f.Statements()])) == 0 {
		// Удаляем использованный объект
		facts.Delete(fi)
		delete(g.stmsMap, f.Statements())
	}

	c.AddOption(option.NewOption(fmt.Sprintf("%s - %s", f.Object(), s), correct))

	return nil
}
