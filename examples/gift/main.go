package main

import (
	"log"
	"os"

	"github.com/lyogh/QuizGenerator/pkg/encoding/gift"
	"github.com/lyogh/QuizGenerator/pkg/fact"
	"github.com/lyogh/QuizGenerator/pkg/generator"
)

var geoData = `
- name: Москва
  statements:
    - город
    - река
    - столица
    - столица России
    - город в России
- name: Берлин
  statements:
    - город
    - город в Германии
    - столица
    - столица Германии
- name: Россия
  statements:
    - страна
- name: Лена
  statements:
    - река
- name: Волга
  statements:
    - река
- name: Волгоград
  statements:
    - город
    - город в России
`

func main() {
	var facts fact.Facts

	// Разбираем факты
	facts.Parse([]byte(geoData))

	// Создаем генератор карточек вопросов
	g := generator.NewGenerator(generator.DefaultParameters)

	// Добавляем факты в генератор
	g.AddFacts(facts)

	// Создаем карточки вопросов
	if err := g.CreateCards(); err != nil {
		log.Fatalf("%v", err)
	}

	// Выводим результат в формате GIFT
	if err := gift.NewEncoder(os.Stdout).Encode(g.Cards()); err != nil {
		log.Fatal(err)
	}
}
