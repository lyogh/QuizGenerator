package falsifier

import "github.com/lyogh/QuizGenerator/pkg/fact"

/*
Фальсификатор: подменяет утверждения на случайно подобранные из других объектов
*/
type statementsShuffler struct {
}

func NewStatementsShuffler() Falsifier {
	return &statementsShuffler{}
}

func (f *statementsShuffler) Falsify(facts fact.Facts) (fact.Facts, error) {
	var lies fact.Facts

	// Все возможные утверждения
	smap := make(map[fact.Statement]struct{})

	for _, oldFact := range facts {
		// Собираем утверждения
		for _, stm := range *oldFact.Statements() {
			smap[*stm] = struct{}{}
		}
	}

	// Перемешиваем утверждения различных фактов
	for _, oldFact := range facts {
		stms := make(fact.Statements, 0, len(smap)-len(*oldFact.Statements()))

		// Добавляем ложные утверждения для обрабатываемого объекта
		for stm := range smap {
			if !oldFact.HasStatement(stm) {
				stms = append(stms, fact.NewStatement(string(stm)))
			}
		}

		if len(stms) > 0 {
			lies = append(lies, fact.NewFact(oldFact.Object(), stms))
		}
	}

	return lies, nil
}
