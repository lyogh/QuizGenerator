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
