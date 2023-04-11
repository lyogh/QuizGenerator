package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lyogh/QuizGenerator/pkg/card"
	"github.com/lyogh/QuizGenerator/pkg/option"
	"golang.org/x/exp/slices"
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
		a                 string
		i, correct, wrong int
	)

	answers := c.Answers()

	fmt.Printf("%d. %s\n", c.Id(), c.Question())

	for _, o := range c.Options() {
		fmt.Printf("%d. %s\n", o.Id(), o)
	}

	cmp := func(o option.Option) bool { return o.Id() == uint(i) }

	for {
		correct, wrong = 0, 0

		switch c.CardType() {
		case card.TypeTrueFalse:
			fmt.Print("Укажите один правильный ответ: ")
		case card.TypeMultipleChoice:
			fmt.Print("Укажите один или несколько правильных ответов через запятую: ")
		}

		_, err := fmt.Scanln(&a)
		if err != nil {
			return 0, err
		}

		for _, a := range strings.Split(a, ",") {
			i, err = strconv.Atoi(a)
			if err != nil {
				fmt.Println(err)
				break
			}

			if !slices.ContainsFunc(c.Options(), cmp) {
				err = fmt.Errorf("вариант ответа %d не найден", i)
				fmt.Println(err)
			}

			if slices.ContainsFunc(answers, cmp) {
				correct++
			} else {
				wrong++
			}
		}

		total := correct + wrong

		if total > len(c.Options()) {
			fmt.Println("Количество ответов превышает количество вариантов")
			continue
		}

		if c.CardType() == card.TypeTrueFalse && total > 1 {
			fmt.Println("Превышено допустимое количество ответов")
			continue
		}

		if err == nil {
			break
		}
	}

	result := float32(correct) / float32(len(answers))

	fmt.Printf("Результат: %.2f (%d/%d)\n", result*100, correct, len(answers))

	return result, nil
}
