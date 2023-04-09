package aiken

import (
	"bufio"
	"fmt"
	"io"
	"strings"

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
	var (
		options strings.Builder
		answer  string
	)

	for _, o := range c.Options() {
		options.WriteString("\n")
		options.WriteString(string(o.Symbol()))
		options.WriteString(". ")
		options.WriteString(o.String())

		if o.Correct() {
			if len(answer) > 0 {
				return fmt.Errorf("в формате Aiken может быть только один вариант правильного ответа")
			}

			answer = string(o.Symbol())
		}
	}

	e.writer.WriteString(fmt.Sprintf("%s%s\nANSWER: %s", c.Question(), options.String(), answer))

	return nil
}
