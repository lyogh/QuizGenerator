package types

type IdSetter interface {
	SetId(uint)
}

type IdGetter interface {
	Id() uint
}

type Shuffler interface {
	Shuffle()
}

/*
Интерфейс удаления элемента из коллеции
*/
type Deleter interface {
	Delete(i int)
}
