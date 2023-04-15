package fact

import (
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

const RootKey = ""

// Группы фактов (блоки файла с фактами)
type groups map[string]Group

type Groups interface {
	// Импортирует данные из файла
	Import(string) error
	// Разбирает факты из данных файла формата yaml
	Parse([]byte) error
	// Возвращает группы
	Data() map[string]Group

	AddGroup(string, Group)
}

func NewGroups() Groups {
	g := make(groups)

	return &g
}

func (g *groups) AddGroup(gname string, group Group) {
	(*g)[gname] = group
}

/*
Добавляет факт в группу
*/
func (g *groups) addFact(gname string, fact *Fact) {
	if _, ok := (*g)[gname]; !ok {
		(*g)[gname] = NewGroup()
	}

	(*g)[gname].AddFact(fact)
}

/*
Импортирует факты из файла формата yaml
*/
func (g *groups) Import(fn string) error {
	file, err := os.Open(fn)
	if err != nil {
		return err
	}

	defer file.Close()

	b, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	if err := g.Parse(b); err != nil {
		return err
	}

	return nil
}

/*
Разбирает факты из данных файла формата yaml
*/
func (g *groups) Parse(data []byte) error {
	blocks := make([]map[string][]map[string][]string, 0)

	if err := yaml.Unmarshal(data, &blocks); err != nil {
		return err
	}

	for _, bs := range blocks {
		for bkey, b := range bs {
			for _, os := range b {
				for okey, o := range os {
					stms := make(Statements, len(o))

					for i, s := range o {
						stms[i] = NewStatement(s)
					}

					g.addFact(bkey, NewFact(okey, stms))
				}
			}
		}
	}

	return nil
}

// Возвращает группы
func (g *groups) Data() map[string]Group {
	return *g
}
