package model

type gopher struct {
	name string
	age int
	isAdult bool
}

type horse struct {
	name string
	weight int
}

func (g gopher) Jump() string{
	if g.age < 65 {
		return g.name + " can jump high!!"
	} else {
		return g.name + " can still jump"
	}
}

func (h horse) Jump() string {
	if h.weight > 2000 {
		return "Too fat...can't jump"
	} else {
		return "Neigh neigh...lol i can jump"
	}
}

type jumper interface {
	Jump() string
}

func GetList() []jumper {
	phil := &gopher{name: "Phil", age: 30}
	noodles := &gopher{name: "Noodles", age: 90}
	barbaro := &horse{name: "Barbaro", weight: 2000}

	list := []jumper{phil, noodles, barbaro}
	return list
}
