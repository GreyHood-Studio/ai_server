package monster

type Monster struct {
	health			int
	moveSpeed		int

	// fsm status
	status			int
}

func (monster Monster) update() {

}

func (monster Monster) Run() {

}

func NewMonster() *Monster{

	monster := &Monster{

	}

	return monster
}