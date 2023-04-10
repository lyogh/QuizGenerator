package fact

import (
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

type Facts []*Fact

type object struct {
	Name       string
	Statements []string
}

/*
Импортирует факты из файла формата yaml
*/
func (f *Facts) Import(fn string) error {
	file, err := os.Open(fn)
	if err != nil {
		return err
	}

	defer file.Close()

	b, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	if err := f.Parse(b); err != nil {
		return err
	}

	return nil
}

/*
Разбирает факты из данных файла формата yaml
*/
func (f *Facts) Parse(data []byte) error {
	objects := make([]object, 0)

	if err := yaml.Unmarshal(data, &objects); err != nil {
		return err
	}

	for _, o := range objects {
		stms := make(Statements, len(o.Statements))

		for i, s := range o.Statements {
			stms[i] = NewStatement(s)
		}

		*f = append(*f, NewFact(o.Name, stms))
	}

	return nil
}
