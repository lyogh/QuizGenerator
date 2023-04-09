package gift

import (
	"bufio"
	"fmt"
	"io"

	"github.com/lyogh/QuizGenerator/pkg/card"
	"github.com/lyogh/QuizGenerator/pkg/encoding"
)

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
		e.encodeTrueFalseCard(c)
	case card.TypeMultipleChoice:
	}

	return nil
}

func (e *encoder) encodeTrueFalseCard(c card.Card) {
	const (
		trueAnswer  = "{T}"
		falseAnswer = "{F}"
	)

	ao := c.Answers()[0]
	a := falseAnswer

	if ao.String() == card.AnswerYes {
		a = trueAnswer
	}

	e.writer.WriteString(fmt.Sprintf("%s%s", c.Question(), a))
}
