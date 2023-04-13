package falsifier

import (
	"errors"

	"github.com/lyogh/QuizGenerator/pkg/fact"
)

var ErrNotChanged = errors.New("утверждение невозможно фальсифицировать")

/*
Интерфейс фальсификатора.
Подменяет данные факта.
*/
type Falsifier interface {
	// Фальсифицирует факты
	Falsify(fact.Facts) (fact.Facts, error)
}

type Falsifiers []Falsifier
