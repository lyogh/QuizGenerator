package main

import (
	"errors"
	"flag"
	"os"

	enc "github.com/lyogh/QuizGenerator/pkg/encoding"
	"github.com/lyogh/QuizGenerator/pkg/generator"
)

type flagSet struct {
	*flag.FlagSet

	cardsMax uint            // Макисмальное количество вопросов в тесте
	encType  enc.EncoderType // Формат файла выгрузки
	dataFile string          // Файл с описанием фактов
	outFile  string          // Файл выгрузки тестовых вопросов
}

func NewFlagSet() (*flagSet, error) {
	fs := &flagSet{
		FlagSet: flag.NewFlagSet("app", flag.ExitOnError)}

	fs.UintVar(&fs.cardsMax, "c", generator.DefaultParameters.CardsMax(), "максимальное количество вопросов в тесте")
	fs.TextVar(&fs.encType, "f", enc.EncoderTypeAiken, "формат выгрузки теста (Aiken, GIFT)")

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
