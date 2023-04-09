package generator

import (
	"log"
	"sync"

	"github.com/lyogh/QuizGenerator/pkg/card"
	"github.com/lyogh/QuizGenerator/pkg/fact"
	"golang.org/x/exp/slices"
)

type Generator interface {
	CreateCards() error

	AddFact(fact *fact.Fact)
	AddFacts(fact.Facts)

	AddDistractor(distractor *fact.Fact)
	AddDistractors(fact.Facts)

	Cards() card.Cards
}

type generator struct {
	// Карточки теста
	cards card.Cards
	// Факты
	facts fact.Facts
	// Дистракторы
	distractors fact.Facts
	// Параметры генератора
	parameters *Parameters
}

func NewGenerator(params *Parameters) Generator {
	g := &generator{
		parameters: params,
	}

	if g.parameters == nil {
		g.parameters = NewParameters()
	}

	return g
}

func (g *generator) AddFact(fact *fact.Fact) {
	g.facts = append(g.facts, fact)
}

func (g *generator) AddFacts(facts fact.Facts) {
	for _, f := range facts {
		g.AddFact(f)
	}
}

func (g *generator) AddDistractor(distractor *fact.Fact) {
	g.distractors = append(g.distractors, distractor)
}

func (g *generator) AddDistractors(distractors fact.Facts) {
	for _, d := range distractors {
		g.AddDistractor(d)
	}
}

func (g *generator) Cards() card.Cards {
	return g.cards
}

/*
Автоматически создает дистракторы на основе фактов
*/
func (g *generator) createDistractors() {
	// Все возможные утверждения
	smap := make(map[fact.Statement]struct{})

	for _, f := range g.facts {
		// Собираем утверждения
		for _, stm := range f.Statements() {
			smap[*stm] = struct{}{}
		}
	}

	// Создаем дистракторы на основе фактов
	for _, f := range g.facts {
		stms := make(fact.Statements, 0, len(smap)-len(f.Statements()))

		// Добавляем ложные утверждения для обрабатываемого объекта
		for stm := range smap {
			if !f.HasStatement(stm) {
				stms = append(stms, fact.NewStatement(string(stm)))
			}
		}

		g.AddDistractors(fact.Facts{fact.NewFact(f.Object(), stms)})
	}
}

func (g *generator) CreateCards() error {
	var (
		wg sync.WaitGroup
		mu sync.Mutex
	)

	g.createDistractors()

	maxCardsOfType := g.parameters.CardsMax()/uint(len(g.parameters.Types())) + g.parameters.CardsMax()%uint(len(g.parameters.Types()))

	create := func(gen Generator) {
		defer wg.Done()

		gen.AddFacts(g.facts)
		gen.AddDistractors(g.distractors)

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

func (g *generator) Shuffle() {
	g.cards.Shuffle()

	if len(g.cards) > int(g.parameters.CardsMax()) {
		g.cards = g.cards[:g.parameters.CardsMax()]
	}
}
