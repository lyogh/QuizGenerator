package encoding

import "github.com/lyogh/QuizGenerator/pkg/card"

type Encoder interface {
	Encode(card.Cards) error
}
