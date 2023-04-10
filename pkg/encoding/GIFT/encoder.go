package gift

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/lyogh/QuizGenerator/pkg/card"
	"github.com/lyogh/QuizGenerator/pkg/encoding"
)

var escapeRegexp = regexp.MustCompile(`([~=#{}:])`)

type encoder struct {
	writer *bufio.Writer
}

func NewEncoder(w io.Writer) encoding.Encoder {
	return &encoder{
		writer: bufio.NewWriter(w),
	}
}

func (e *encoder) Encode(cards card.Cards) error {
	for _, c := range cards {
		if err := e.encodeCard(c); err != nil {
			return err
		}

		e.writer.WriteString("\n\n")
	}

	return e.writer.Flush()
}

func (e *encoder) encodeCard(c card.Card) error {
	switch c.CardType() {
	case card.TypeTrueFalse:
		e.encodeTrueFalse(c)
	case card.TypeMultipleChoice:
		e.encodeMultipleChoice(c)
	}

	return nil
}

/*
Преобразует данные каточки True/False
*/
func (e *encoder) encodeTrueFalse(c card.Card) {
	ao := c.Answers()[0]
	a := "{F}"

	if ao.String() == card.AnswerYes {
		a = "{T}"
	}

	e.writer.WriteString(fmt.Sprintf("%s%s", e.escapeString(c.Question()), a))
}

// Преобразует данные карточки множественного выбора
func (e *encoder) encodeMultipleChoice(c card.Card) {
	const wrong = "~"

	var b strings.Builder

	ac := len(c.Answers())

	for _, o := range c.Options() {
		b.WriteString("\n")

		if o.Correct() {
			if ac > 1 {
				b.WriteString(fmt.Sprintf("%s%%%f%%", wrong, 1/float32(ac)*100))
			} else {
				b.WriteString("=")
			}
		} else {
			b.WriteString(wrong)

			if ac > 1 {
				b.WriteString("%-100%")
			}
		}

		b.WriteRune(o.Symbol())
		b.WriteString(". ")
		b.WriteString(e.escapeString(o.String()))
	}

	e.writer.WriteString(fmt.Sprintf("%s {%s\n}", e.escapeString(c.Question()), b.String()))
}

func (e *encoder) escapeString(s string) string {
	return string(escapeRegexp.ReplaceAll([]byte(s), []byte(`\$1`)))
}
