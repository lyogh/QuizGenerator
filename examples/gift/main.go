package main

import (
	"log"
	"os"

	"github.com/lyogh/QuizGenerator/pkg/encoding/gift"
	"github.com/lyogh/QuizGenerator/pkg/fact"
	"github.com/lyogh/QuizGenerator/pkg/generator"
)

func main() {
	g := generator.NewGenerator(nil)

	g.AddFact(fact.NewFact("Москва", fact.Statements{
		fact.NewStatement("Город"),
		fact.NewStatement("Столица"),
	}))

	g.AddFact(fact.NewFact("Берлин", fact.Statements{
		fact.NewStatement("Город"),
		fact.NewStatement("Столица"),
	}))

	g.AddFact(fact.NewFact("Волгоград", fact.Statements{
		fact.NewStatement("Город"),
	}))

	g.AddFact(fact.NewFact("Россия", fact.Statements{
		fact.NewStatement("Страна"),
	}))

	g.AddFact(fact.NewFact("Германия", fact.Statements{
		fact.NewStatement("Страна"),
	}))

	if err := g.CreateCards(); err != nil {
		log.Fatal(err)
	}

	if err := gift.NewEncoder(os.Stdout).Encode(g.Cards()); err != nil {
		log.Fatal(err)
	}
}
