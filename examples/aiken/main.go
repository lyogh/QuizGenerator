package main

import (
	"log"
	"os"

	"github.com/lyogh/QuizGenerator/pkg/card"
	"github.com/lyogh/QuizGenerator/pkg/encoding/aiken"
	"github.com/lyogh/QuizGenerator/pkg/fact"
	"github.com/lyogh/QuizGenerator/pkg/generator"
)

var geoData = `
# Географические факты
---
- Все факты:
  - Москва:
    - город
    - река
    - столица
    - столица России
    - город в России
    - город в России c 10 млн. жителей
  - Берлин:
    - город
    - город в Германии
    - столица
    - столица Германии
  - Россия:
    - страна
  - Лена:
    - река
  - Волга:
    - река
  - Волгоград:
    - город
    - город в России`

func main() {
	facts := fact.NewGroups()

	// Парсим факты
	if err := facts.Parse([]byte(geoData)); err != nil {
		log.Fatal(err)
	}

	p := generator.DefaultParameters
	p.SetCardParameters(card.AikenParameters)

	// Создаем генератор карточек вопросов с параметрами для формата Aiken (один правильный ответ на вопрос)
	g := generator.NewGenerator(p)

	// Добавляем факты в генератор
	g.SetFacts(facts)

	// Создаем карточки вопросов
	if err := g.CreateCards(); err != nil {
		log.Fatalf("%v", err)
	}

	// Выводим результат в формате Aiken
	if err := aiken.NewEncoder(os.Stdout).Encode(g.Cards()); err != nil {
		log.Fatal(err)
	}
}
