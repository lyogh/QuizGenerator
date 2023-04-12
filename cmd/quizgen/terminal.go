package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/lyogh/QuizGenerator/pkg/card"
	"github.com/lyogh/QuizGenerator/pkg/option"
	"golang.org/x/exp/slices"
)

const (
	ColorQuestion   = "0;30;47"
	ColorInput      = "0;36"
	ColorPositive   = "1;32"
	ColorNegative   = "1;31"
	ColorIncomplete = "1;33"
)

type terminal struct {
}

func NewTerminal() *terminal {
	return new(terminal)
}

/*
Выводит карточку вопроса и ожидает ответ от пользователя
*/
func (t *terminal) renderCard(c card.Card) (float32, error) {
	var (
		correct int
	)

	answers := c.Answers()

	t.print(ColorQuestion, fmt.Sprintf("%d. %s", c.Id(), c.Question()))

	for _, o := range c.Options() {
		fmt.Printf("\n%c. %s", unicode.ToLower(o.Symbol()), o)
	}

	fmt.Print("\n")

	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanWords(data, atEOF)
		if err == nil {
			if token == nil {
				return 0, nil, bufio.ErrFinalToken
			}

			if !slices.ContainsFunc(c.Options(), func(o option.Option) bool {
				return string(unicode.ToLower(o.Symbol())) == strings.ToLower(string(token))
			}) {
				return 0, nil, fmt.Errorf("вариант ответа %s не найден", token)
			}
		}

		return
	}

	// Обрабатывает ввод ответа пользователя
	uinp := func() (int, error) {
		var (
			correct int
			r       rune
		)

		cmp := func(o option.Option) bool { return unicode.ToLower(o.Symbol()) == unicode.ToLower(r) }

		switch c.CardType() {
		case card.TypeTrueFalse:
			t.print(ColorInput, "Укажите один правильный ответ: ")
		case card.TypeMultipleChoice:
			t.print(ColorInput, "Укажите один или несколько правильных ответов (через пробел): ")
		}

		sc := bufio.NewScanner(os.Stdin)

		sc.Split(split)

		a := make([]rune, 0)

		for sc.Scan() && len(sc.Text()) > 0 {
			a = append(a, []rune(sc.Text())[0])
		}

		if sc.Err() != nil {
			return 0, sc.Err()
		}

		for _, r = range a {
			if slices.ContainsFunc(answers, cmp) {
				correct++
			}
		}

		if len(a) > len(c.Options()) {
			return 0, fmt.Errorf("количество %d ответов превышает количество вариантов %d", len(a), len(c.Options()))
		}

		if c.CardType() == card.TypeTrueFalse && len(a) > 1 {
			return 0, errors.New("превышено допустимое количество ответов")
		}

		return correct, nil
	}

	for {
		var err error

		correct, err = uinp()

		if err != nil {
			fmt.Println(err)
			continue
		}

		break
	}

	result := float32(correct) / float32(len(answers))

	color := ColorIncomplete

	switch result {
	case 0:
		color = ColorNegative
	case 1:
		color = ColorPositive
	}

	t.print(color, fmt.Sprintf("Результат: %.2f (%d/%d)", result*100, correct, len(answers)))
	fmt.Print("\n")

	return result, nil
}

/*
Выводит цветной текст
*/
func (t *terminal) print(color string, text string) {
	fmt.Print("\x1b[" + color + "m" + text + "\x1b[0m")
}
