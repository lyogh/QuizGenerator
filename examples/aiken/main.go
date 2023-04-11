package main

import (
	"log"
	"os"

	"github.com/lyogh/QuizGenerator/pkg/card"
	"github.com/lyogh/QuizGenerator/pkg/encoding/aiken"
	"github.com/lyogh/QuizGenerator/pkg/fact"
	"github.com/lyogh/QuizGenerator/pkg/generator"
)

var abap = `
  - name: The Object Navigator
    statements:
  - incorporates a total of 11 browsers
  - can display and edit ABAP programs
  - can display and edit screens
  - can display and edit menus
  - can maintain ABAP Dictionary
  - name: The Repository Browser
    statements:
  - is started by default when you execute Transaction SE80 for the Object Navigator
  - name: The Repository Information System
    statements:
  - is a useful tool to search for customer exits/function exits and BAdIs in the SAP system
  - name: Enhancement Information System
    statements:
  - can display Enhancement definitions and implementations
  - name: Customer repository object
    statements:
  - have to be assigned to a package
  - name: Package
    statements:
  - use interfaces and visibility to make their elements visible to other packages
  - can be nested
`

func main() {
	var facts fact.Facts

	// Парсим факты
	facts.Parse([]byte(abap))

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
