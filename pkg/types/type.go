package types

/*
Интерфейс определения Ид
*/
type IdSetter interface {
	SetId(uint)
}

/*
Интерфейс получения Ид
*/
type IdGetter interface {
	Id() uint
}

/*
Интерфейс перетасовщика
*/
type Shuffler interface {
	Shuffle()
}

/*
Интерфейс удаления элемента из коллеции
*/
type Deleter interface {
	Delete(i int)
}
