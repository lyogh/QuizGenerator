package generator

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/lyogh/QuizGenerator/internal/slice"
	"github.com/lyogh/QuizGenerator/pkg/card"
	"github.com/lyogh/QuizGenerator/pkg/fact"
	"github.com/lyogh/QuizGenerator/pkg/option"
)

// Генератор карточек вопросов с множеством вариантов ответов
type multipleChoiceCard struct {
	generator

	stmsMap map[*fact.Statements]*fact.Statements
}

var ErrNoMoreLies = errors.New("не определены дистракторы")

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

/*
Генерирует карточки вопросов теста
*/
func (g *multipleChoiceCard) CreateCards() error {
	facts := g.groups.Data()[fact.RootKey].Facts()

	slice.Shuffle(*facts)
	slice.Shuffle(*g.groups.Data()[fact.RootKey].Lies())

	g.stmsMap = make(map[*fact.Statements]*fact.Statements)

	for i := 0; i < int(g.parameters.CardsMax()) && len(*facts) > 0; i++ {
		if err := g.addCard(); err != nil && err != ErrNoMoreLies {
			return err
		}
	}

	g.Shuffle()

	return nil
}

/*
Создает карточку вопроса
*/
func (g *multipleChoiceCard) addCard() error {
	var (
		facts, lies *fact.Facts
		q           string
	)

	// Случайно выбираем вопрос "Выберите верные утверждения" или "Выберите НЕ верные утверждения"
	if rand.Intn(2) == 0 {
		facts, lies, q = g.groups.Data()[fact.RootKey].Facts(), g.groups.Data()[fact.RootKey].Lies(), card.QuestionTextChooseCorrectStatements
	} else {
		facts, lies, q = g.groups.Data()[fact.RootKey].Lies(), g.groups.Data()[fact.RootKey].Facts(), card.QuestionTextChooseWrongStatements
	}

	c, err := card.NewCard(q, card.TypeMultipleChoice)
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
	for i := 0; i < alen && len(*facts) > 0; i++ {
		g.addOptions(c, facts, true)
	}

	// В карточке не может быть только один вариант ответа
	if len(*lies) == 0 && len(c.Answers()) == 1 {
		return ErrNoMoreLies
	}

	for i := len(c.Answers()); i < olen && len(*lies) > 0; i++ {
		g.addOptions(c, lies, false)
	}

	if rand.Intn(2) == 0 && len(c.Options()) > card.MinOptions {
		// Группируем ответы
		//c.GroupOptions()
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
