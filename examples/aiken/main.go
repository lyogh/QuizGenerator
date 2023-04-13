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
- name: Москва
  statements:
    - город
    - река
    - столица
    - столица России
    - город в России
    - город в России c 10 млн. жителей
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

	// Парсим факты
	facts.Parse([]byte(geoData))

	p := generator.DefaultParameters
	p.SetCardParameters(card.AikenParameters)

	// Создаем генератор карточек вопросов с параметрами для формата Aiken (один правильный ответ на вопрос)
	g := generator.NewGenerator(p)

	// Добавляем факты в генератор
	g.AddFacts(facts)

	// Создаем карточки вопросов
	if err := g.CreateCards(); err != nil {
		log.Fatalf("%v", err)
	}

	// Выводим результат в формате Aiken
	if err := aiken.NewEncoder(os.Stdout).Encode(g.Cards()); err != nil {
		log.Fatal(err)
	}
}
