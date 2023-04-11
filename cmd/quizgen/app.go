package main

import (
	"fmt"
	"log"
	"os"

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
		facts fact.Facts
		err   error
	)

	a.fs, err = NewFlagSet()
	if err != nil {
		log.Fatal(err)
	}

	a.fs.parse()

	p := generator.DefaultParameters

	p.SetCardsMax(a.fs.cardsMax)

	// Новый генератор
	a.g = generator.NewGenerator(generator.DefaultParameters)

	// Импортируем факты из файла
	if err := facts.Import(a.fs.dataFile); err != nil {
		log.Fatal(err)
	}

	// Добавляем факты в генератор
	a.g.AddFacts(facts)

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
	var result float32

	// Начинаем тестирование
	t := NewTerminal()

	for _, c := range a.g.Cards() {
		res, err := t.renderCard(c)
		if err != nil {
			log.Fatal(err)
		}

		result += res
	}

	result = result / float32(len(a.g.Cards()))

	fmt.Printf("Общий результат: %.2f%%", result*100)
	return nil
}
