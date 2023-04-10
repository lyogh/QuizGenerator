package card

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/lyogh/QuizGenerator/internal/slice"
	"github.com/lyogh/QuizGenerator/pkg/fact"
	"github.com/lyogh/QuizGenerator/pkg/option"
	"github.com/lyogh/QuizGenerator/pkg/types"
	"golang.org/x/exp/slices"
)

const (
	TypeMultipleChoice = iota
	TypeTrueFalse
)

const (
	QuestionTextStatementCorrect = "Верно ли утверждение"
	QuestionTextChooseStatements = "Выберите верные утверждения"
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

type card struct {
	// Ид вопроса
	id uint
	// Вопрос
	question string
	// Варианты ответов
	options option.Options
	// Тип карточки
	ctype CardType
}

// Карточка вопроса
type Card interface {
	types.IdSetter
	types.IdGetter
	option.OptionsManager

	Question() string
	CardType() CardType
	Answers() option.Options
	GenerateOptionsByFacts(facts fact.Facts, distractors fact.Facts, params Parameters) error
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
Создает варианты ответов в карточке вопроса на основе фактов и дистракторов
*/
func (c *card) GenerateOptionsByFacts(facts, distractors fact.Facts, params Parameters) (err error) {
	// Добавляем правильные варианты выбора
	c.options, err = c.randomOptions(facts, *params.Options(), true)
	if err != nil {
		return err
	}

	/*	if uint(len(c.options)) < params.options.GetMax() && rand.Intn(2) == 0 {
		// Объединим несколько вариантов в один
		c.addOptionsMix()
	} else {*/
	// Ограничиваем список правильных ответов по максимально допустимому количеству
	if uint(len(c.options)) > params.answers.max {
		c.options = c.options[:params.answers.max]
	}

	// Добавим дистракторы
	c.addDistractors(distractors, params)
	//}

	c.options.Shuffle()

	return nil
}

func (c *card) addDistractors(distractors fact.Facts, params Parameters) error {
	disInt := NewInterval(0, params.options.max)

	if params.answers.max == 1 {
		// Если максимальное количество правильных ответов = 1, тогда должен быть минимум 1 дистрактор
		disInt.SetMin(1)
	}

	// Дистракторы
	dis, err := c.randomOptions(distractors, *disInt, false)
	if err != nil {
		return err
	}

	if len(dis) == 0 {
		return errors.New("не определены дистракторы")
	}

	c.options = append(c.options, dis...)

	// Если количество вариантов ответов превышает допустимое, тогда удаляем лишние
	if uint(len(c.options)) > params.options.GetMax() {
		c.options = c.options[:params.options.GetMax()]
	}

	if uint(len(c.options)) < params.options.GetMax() {
		/* Добавляем вариант ответа для группировки других вариантов.
		   - есть такие:
				1. Ответ A
				2. Ответ B
		   - а добавим новый:
				3. Варианты 1 и 2 правильные
		*/
		if rand.Intn(2) == 0 {
			c.addOptionsMix()
		}
	}

	return nil
}

/*
Добавляет вариант для группы ответов
*/
func (c *card) addOptionsMix() {
	var (
		correct, all bool
	)

	c.options.Shuffle()

	olen := len(c.options)

	if len(c.options) > MinOptions {
		olen = rand.Intn(olen-MinOptions) + MinOptions
	}

	options := make(option.Options, olen)

	for i, o := range c.options {
		if i < len(options) {
			options[i] = o
		}
	}

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

	c.options = append(c.options, option.NewMixedOption(options, correct, all))
}

func (c *card) randomOptions(facts fact.Facts, interval Interval, correctAnswer bool) (option.Options, error) {
	var ol int

	if uint(len(facts)) < interval.GetMin() {
		return nil, errors.New("недостаточно данных для генерации вариантов ответов")
	}

	facts = slices.Clone(facts)

	if interval.GetMax() > uint(len(facts)) {
		interval.SetMax(uint(len(facts)))
	}

	d := interval.GetMax() - interval.GetMin()
	if d > 0 {
		ol = rand.Intn(int(d))
	}

	ol += int(interval.GetMin())

	options := make(option.Options, ol)

	for i := 0; i < len(options); i++ {
		mr := ol
		if mr > len(facts) {
			mr = len(facts)
		}

		if mr == 0 {
			break
		}

		fi := rand.Intn(mr)
		fact := facts[fi]
		facts[fi] = nil

		facts = append(facts[:fi], facts[fi+1:]...)

		stms := fact.Statements()
		options[i] = option.NewOption(fmt.Sprintf("%s - %s", fact.Object(), stms[rand.Intn(len(stms))]), correctAnswer)
	}

	return options, nil
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
