package encoding

import (
	"strings"

	"github.com/lyogh/QuizGenerator/pkg/card"
)

// Текстовые обозначения форматов выгрузки
const (
	EncoderTypeTextAiken = "aiken"
	EncoderTypeTextGIFT  = "gift"
)

// Числовые обозначения форматов выгрузки
const (
	EncoderTypeAiken = EncoderType(iota)
	EncoderTypeGIFT
)

type EncoderType int // Код формата выгрузки

type Encoder interface {
	Encode(card.Cards) error
}

/*
Преобразует текстовое обозначение формата в код
*/
func (e *EncoderType) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {
	case EncoderTypeTextAiken:
		*e = EncoderTypeAiken
	case EncoderTypeTextGIFT:
		*e = EncoderTypeGIFT
	}

	return nil
}

/*
Преобразует код формата в текстовое обозначение
*/
func (e EncoderType) MarshalText() (text []byte, err error) {
	switch e {
	case EncoderTypeAiken:
		text = []byte(EncoderTypeTextAiken)
	case EncoderTypeGIFT:
		text = []byte(EncoderTypeTextGIFT)
	}

	return
}
