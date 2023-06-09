package main

import (
	"fmt"
	"log"
	"os"

	"github.com/lyogh/QuizGenerator/pkg/card"
	enc "github.com/lyogh/QuizGenerator/pkg/encoding"
	"github.com/lyogh/QuizGenerator/pkg/encoding/aiken"
	"github.com/lyogh/QuizGenerator/pkg/encoding/gift"
	"github.com/lyogh/QuizGenerator/pkg/fact"
	"github.com/lyogh/QuizGenerator/pkg/generator"
)

type app struct {
	fs *flagSet
	g  generator.Generator
}

func NewApp() *app {
	return new(app)
}

/*
Запускает приложение
*/
func (a *app) start() {
	var (
		err error
	)

	a.fs, err = NewFlagSet()
	if err != nil {
		log.Fatal(err)
	}

	a.fs.parse()

	p := *generator.DefaultParameters

	// Максимальное количество карточек в тесте
	p.SetCardsMax(a.fs.cardsMax)

	// Типы карточек
	ct := make(card.CardTypes, 0)
	for t, tb := range a.fs.cardTypes {
		if tb {
			ct = append(ct, card.CardType(t))
		}
	}

	p.SetTypes(ct)

	// Новый генератор
	a.g = generator.NewGenerator(&p)

	facts := fact.NewGroups()

	// Импортируем факты из файла
	if err := facts.Import(a.fs.dataFile); err != nil {
		log.Fatal(err)
	}

	// Добавляем факты в генератор
	a.g.SetFacts(facts)

	// Генерируем карточки вопросов
	if err := a.g.CreateCards(); err != nil {
		log.Fatal(err)
	}

	// Если передали название файла для экспорта, тогда выгружаем вопросы в указанный файл
	if len(a.fs.outFile) > 0 {
		if err := a.export(); err != nil {
			log.Fatal(err)
		}
	} else {
		a.test()
	}
}

/*
Выгружает тестовые вопросы в файл
*/
func (a *app) export() error {
	var encoder enc.Encoder

	f, err := os.Create(a.fs.outFile)
	if err != nil {
		return err
	}

	defer f.Close()

	switch a.fs.encType {
	case enc.EncoderTypeAiken:
		encoder = aiken.NewEncoder(f)
	case enc.EncoderTypeGIFT:
		encoder = gift.NewEncoder(f)
	default:
		log.Fatal("неизвестный формат выгрузки теста")
	}

	if err := encoder.Encode(a.g.Cards()); err != nil {
		return err
	}

	return nil
}

/*
Запускает тестирование
*/
func (a *app) test() error {
	var result card.CardResult

	// Начинаем тестирование
	t := NewTerminal()

	for _, c := range a.g.Cards() {
		err := t.renderCard(c)
		if err != nil {
			log.Fatal(err)
		}

		result.Value += c.Result().Value
		result.Duration += c.Result().Duration
	}

	result.Value = result.Value / float32(len(a.g.Cards()))

	color := ColorIncomplete

	switch result.Value {
	case 0:
		color = ColorNegative
	case 1:
		color = ColorPositive
	}

	fmt.Print("\n")
	t.print(color, fmt.Sprintf("Общий результат: %.2f%%", result.Value*100))
	fmt.Print("\n")
	t.print(color, fmt.Sprintf("Общее время: %.2fс", result.Duration.Seconds()))
	fmt.Print("\n")

	return nil
}
