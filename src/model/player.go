package player

type Player struct {
	Id int
	Name string
}

func New (id int, name string) Player {
	return Player{id, name}
}

func (p Player) Validate() bool {
	return true
}
