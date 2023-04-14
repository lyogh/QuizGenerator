package falsifier

import (
	"math/rand"
	"regexp"
	"strconv"
	"strings"

	"github.com/lyogh/QuizGenerator/pkg/fact"
)

/*
Фальсификатор: подменяет числовые значения в утверждении на случайно подобранные
*/
type numericFalsifier struct {
	rx *regexp.Regexp
}

func NewNumericFalsifier() Falsifier {
	return &numericFalsifier{
		rx: regexp.MustCompile(`\d+`),
	}
}

func (f *numericFalsifier) Falsify(facts fact.Facts) (fact.Facts, error) {
	var lies fact.Facts

	// Перебираем факты
	for _, oldFact := range facts {
		stms := make(fact.Statements, 0)

		// Подменяем утверждения
		for _, oldStm := range *oldFact.Statements() {
			newStm, err := f.falsifyStatement(*oldStm)
			if err != nil {
				if err != ErrNotChanged {
					return nil, err
				}

				continue
			}

			stms = append(stms, newStm)
		}

		if len(stms) > 0 {
			lies = append(lies, fact.NewFact(oldFact.Object(), stms))
		}
	}

	return lies, nil
}

func (f *numericFalsifier) falsifyStatement(stm fact.Statement) (*fact.Statement, error) {
	s := string(stm)

	idx := f.rx.FindAllStringIndex(s, -1)

	if len(idx) == 0 {
		// Не нашли числовых значений в утверждении
		return nil, ErrNotChanged
	}

	repl := make([][]int, len(idx))

	for i, p := range idx {
		old, err := strconv.Atoi(s[p[0]:p[1]])
		if err != nil {
			return nil, err
		}

		new := old

		// Увеличиваем или уменьшаем значение
		if rand.Intn(2) == 1 {
			new -= rand.Intn(old + 1)
		} else {
			new += rand.Intn(old + 1)
		}

		repl[i] = append(repl[i], []int{old, new}...)
	}

	for _, r := range repl {
		s = strings.Replace(s, strconv.Itoa(r[0]), strconv.Itoa(r[1]), -1)
	}

	return fact.NewStatement(s), nil
}
