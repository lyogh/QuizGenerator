package main

import (
	"log"
	"os"

	"github.com/lyogh/QuizGenerator/pkg/encoding/gift"
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

	// Создаем генератор карточек вопросов
	g := generator.NewGenerator(generator.DefaultParameters)

	// Добавляем факты в генератор
	g.SetFacts(facts)

	// Создаем карточки вопросов
	if err := g.CreateCards(); err != nil {
		log.Fatalf("%v", err)
	}

	// Выводим результат в формате GIFT
	if err := gift.NewEncoder(os.Stdout).Encode(g.Cards()); err != nil {
		log.Fatal(err)
	}
}
