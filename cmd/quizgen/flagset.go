package main

import (
	"errors"
	"flag"
	"os"

	"github.com/lyogh/QuizGenerator/pkg/card"
	enc "github.com/lyogh/QuizGenerator/pkg/encoding"
	"github.com/lyogh/QuizGenerator/pkg/generator"
)

type flagSet struct {
	*flag.FlagSet

	cardsMax  uint            // Макисмальное количество вопросов в тесте
	encType   enc.EncoderType // Формат файла выгрузки
	dataFile  string          // Файл с описанием фактов
	outFile   string          // Файл выгрузки тестовых вопросов
	cardTypes [2]bool         // Типы карточек
}

func NewFlagSet() (*flagSet, error) {
	fs := &flagSet{
		FlagSet: flag.NewFlagSet("app", flag.ExitOnError)}

	fs.UintVar(&fs.cardsMax, "c", generator.DefaultParameters.CardsMax(), "максимальное количество вопросов в тесте")
	fs.TextVar(&fs.encType, "f", enc.EncoderTypeAiken, "формат выгрузки теста (Aiken, GIFT)")

	// Типы карточек
	fs.BoolVar(&fs.cardTypes[card.TypeTrueFalse], "ttf", true, "разрешить тип карточки Да/Нет")
	fs.BoolVar(&fs.cardTypes[card.TypeMultipleChoice], "tmc", true, "разрешить тип карточки множественного выбора")

	return fs, nil
}

/*
Разбирает параметры командной строки
*/
func (fs *flagSet) parse() error {
	fs.Parse(os.Args[1:])

	fs.dataFile = fs.Arg(0)
	if len(fs.dataFile) == 0 {
		return errors.New("укажите файл с описанием фактов")
	}

	fs.outFile = fs.Arg(1)

	return nil
}
