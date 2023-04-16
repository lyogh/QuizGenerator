package card

import (
	"math/rand"
	"time"

	"github.com/lyogh/QuizGenerator/internal/slice"
	"github.com/lyogh/QuizGenerator/pkg/option"
	"github.com/lyogh/QuizGenerator/pkg/types"
	"golang.org/x/exp/slices"
)

const (
	TypeMultipleChoice = iota
	TypeTrueFalse
)

const (
	QuestionTextStatementCorrect        = "Верно ли утверждение"
	QuestionTextChooseCorrectStatements = "Выберите верные утверждения"
	QuestionTextChooseWrongStatements   = "Выберите неверные утверждения"
)

const (
	AnswerYes      = "Да"
	AnswerNo       = "Нет"
	AnswerAllTrue  = "Все указанные ответы правильные"
	AnswerTrueList = "Правильные ответы"
)

type (
	CardType  byte
	CardTypes []CardType
)

type CardResult struct {
	// Оценка ответа
	Value float32
	// Длительность ответа
	Duration time.Duration
}

type card struct {
	// Ид вопроса
	id uint
	// Вопрос
	question string
	// Варианты ответов
	options option.Options
	// Тип карточки
	ctype  CardType
	result CardResult
}

// Карточка вопроса
type Card interface {
	types.IdSetter
	types.IdGetter
	option.OptionsManager

	Question() string
	CardType() CardType
	Answers() option.Options
	GroupOptions()

	// Оценка ответа пользователя
	SetResult(CardResult)
	Result() CardResult
}

func NewCard(question string, ctype CardType) (Card, error) {
	card := card{
		question: question,
		ctype:    ctype,
	}

	return &card, nil
}

func (c *card) SetId(id uint) {
	c.id = id
}

func (c *card) Id() uint {
	return c.id
}

/*
Добавляет вариант для группы ответов
*/
func (c *card) GroupOptions() {
	var (
		correct, all bool
	)

	c.Options().Shuffle()

	olen := rand.Intn(len(c.options)-MinOptions) + MinOptions

	options := make(option.Options, olen)
	copy(options, c.options)

	answers := c.Answers()

	if slices.CompareFunc(options, answers, func(o1 option.Option, o2 option.Option) int { return int(o1.Id() - o2.Id()) }) == 0 {
		correct = true

		// т.к. добавленный вариант покрывает все правильные ответы, то у других вариантов убираем признак верного ответа
		for _, a := range answers {
			a.SetCorrect(false)
		}
	}

	if len(options) == len(c.options) {
		all = true
	}

	c.AddOption(option.NewMixedOption(options, correct, all))
}

/*
Добавляет вариант ответа
*/
func (c *card) AddOption(option option.Option) {
	c.options = append(c.options, option)
}

func (c *card) DeleteOption(option option.Option) {
	c.options = slice.Delete(c.options, option)
}

func (c *card) Options() option.Options {
	return c.options
}

func (c *card) Question() string {
	return c.question
}

func (c *card) CardType() CardType {
	return c.ctype
}

/*
Возвращает правильные ответы
*/
func (c *card) Answers() option.Options {
	options := make(option.Options, len(c.options))

	i := 0
	for _, o := range c.options {
		if o.Correct() {
			options[i] = o
			i++
		}
	}

	return options[:i]
}

/*
Определяет оценку ответа пользователя
*/
func (c *card) SetResult(result CardResult) {
	c.result = result
}

/*
Возвращает оценку ответа пользователя
*/
func (c *card) Result() CardResult {
	return c.result
}
