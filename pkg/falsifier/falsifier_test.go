package falsifier

import (
	"testing"

	"github.com/lyogh/QuizGenerator/pkg/fact"
)

func TestNumeric(t *testing.T) {
	type testData struct {
		fact   fact.Fact
		result bool
	}
	var test []testData = []testData{
		{fact: *fact.NewFact("Moscow", fact.Statements{fact.NewStatement("is a city with a population estimated at 13.0 million residents")}), result: true},
		{fact: *fact.NewFact("Moscow", fact.Statements{fact.NewStatement("is the capital and largest city of Russia")}), result: false},
	}

	f := NewNumericFalsifier()

	for _, td := range test {
		lies, err := f.Falsify(fact.Facts{&td.fact})
		if td.result && (len(lies) == 0 || lies[0].HasStatement(*(*td.fact.Statements())[0]) || err != nil) {
			t.Errorf("expected %v got %v", td.result, false)
		}

		if !td.result && len(lies) > 0 {
			t.Errorf("expected %v got %v", td.result, true)
		}
	}

}
