package generator

import (
	"log"
	"sync"

	"github.com/lyogh/QuizGenerator/pkg/card"
	"github.com/lyogh/QuizGenerator/pkg/fact"
	"github.com/lyogh/QuizGenerator/pkg/falsifier"
	"golang.org/x/exp/slices"
)

// Генератор карточек вопросов
type Generator interface {
	CreateCards() error

	AddFact(fact *fact.Fact)
	AddFacts(fact.Facts)

	AddDistractor(distractor *fact.Fact)
	AddDistractors(fact.Facts)

	Cards() card.Cards
}

// Генератор карточек вопросов
type generator struct {
	// Карточки теста
	cards card.Cards
	// Факты
	facts fact.Facts
	// Дистракторы
	distractors fact.Facts
	// Параметры генератора
	parameters *Parameters
	// Фальсификаторы фактов
	falsifiers falsifier.Falsifiers
}

/*
Создает новый объект генератора карточек вопросов
*/
func NewGenerator(params *Parameters) Generator {
	g := &generator{
		parameters: params,
		falsifiers: falsifier.Falsifiers{
			falsifier.NewNumericFalsifier(),
			falsifier.NewStatementsShuffler(),
		},
	}

	if g.parameters == nil {
		g.parameters = NewParameters()
	}

	return g
}

/*
Добавляет факт в генератор
*/
func (g *generator) AddFact(fact *fact.Fact) {
	g.facts = append(g.facts, fact)
}

/*
Добавляет факты в генератор
*/
func (g *generator) AddFacts(facts fact.Facts) {
	for _, f := range facts {
		g.AddFact(f)
	}
}

/*
Добавляет дистрактор в генератор
*/
func (g *generator) AddDistractor(distractor *fact.Fact) {
	g.distractors = append(g.distractors, distractor)
}

/*
Добавляет дистракторы в генератор
*/
func (g *generator) AddDistractors(distractors fact.Facts) {
	for _, d := range distractors {
		g.AddDistractor(d)
	}
}

/*
Возвращает созданные карточки вопросов
*/
func (g *generator) Cards() card.Cards {
	return g.cards
}

/*
Создает карточки вопросов
*/
func (g *generator) CreateCards() error {
	var (
		wg sync.WaitGroup
		mu sync.Mutex
	)

	distractors, err := g.Falsify(g.facts)
	if err != nil {
		return err
	}

	maxCardsOfType := g.parameters.CardsMax()/uint(len(g.parameters.Types())) + g.parameters.CardsMax()%uint(len(g.parameters.Types()))

	create := func(gen Generator) {
		defer wg.Done()

		gen.AddFacts(g.facts)
		gen.AddDistractors(distractors)

		if err := gen.CreateCards(); err != nil {
			log.Println(err)
		}

		mu.Lock()
		g.cards = append(g.cards, gen.Cards()...)
		mu.Unlock()
	}

	getParams := func() *Parameters {
		p := *g.parameters
		p.SetCardsMax(maxCardsOfType)
		return &p
	}

	// Карточки вида True/False
	if g.parameters.CardParameters().Answers().GetMin() == 1 && // В карточках True/False может быть только два выбора с одним правильным ответом
		slices.Contains(g.parameters.Types(), card.TypeTrueFalse) {
		wg.Add(1)
		go create(NewTrueFalseGenerator(getParams()))
	}

	// Карточки со множественным выбором
	if slices.Contains(g.parameters.Types(), card.TypeMultipleChoice) {
		wg.Add(1)
		go create(NewMultipleChoiceCard(getParams()))
	}

	wg.Wait()

	g.Shuffle()

	return nil
}

/*
Перемешивает карточки вопросов
*/
func (g *generator) Shuffle() {
	g.cards.Shuffle()

	if len(g.cards) > int(g.parameters.CardsMax()) {
		g.cards = g.cards[:g.parameters.CardsMax()]
	}
}

/*
Фальсифицирует факты
*/
func (g *generator) Falsify(facts fact.Facts) (fact.Facts, error) {
	var lies fact.Facts

	for _, f := range g.falsifiers {
		dis, err := f.Falsify(facts)
		if err != nil {
			return nil, err
		}

		lies = append(lies, dis...)
	}

	return lies, nil
}
