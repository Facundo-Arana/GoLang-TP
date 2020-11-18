package afa

import "fmt"

// Tournament service interface
type Tournament interface {
	unafuncion(s string)
}

// Team service interface
type Team interface {
	//TODO encontrar el lugar donde van la funciones
}

// Player service interface
type Player interface {
	//TODO encontrar el lugar donde van la funciones
}

type tournament struct {
	name  string
	teams map[int]*Team
}

type team struct {
	name    string
	players map[int]*Player
}
type player struct {
	name string
	num  int
}

// NewTournament ...
func NewTournament() Tournament {
	teams := make(map[int]*Team)
	return tournament{
		teams,
	}
}

// NewTeam ...
func NewTeam(n string, d string) Team {
	players := make(map[int]*Player)
	return team{
		n,
		d,
		players,
	}
}

// NewPlayer ...
func NewPlayer(n string, num int) Player {
	return player{
		n,
		num,
	}
}

func (t *tournament) unafuncion(s string) {
	fmt.Println(s)
}

//---------me pregunto en este momento si estan bien estas funciones que estan abajo

/*
  Añade players a un Team
  @param { s string }  es el nombre del jugador
  @param { v int } es el numero de camiseta

*/
func (e *team) add(c int, n string) {
	e.players[c] = n
}

/*
	Añade un Team al tournament
*/
func (t *tournament) add(e team) {
	t.teams[e.ID] = &e
}

// GetTeam retorna un Team dado un ID
func (t tournament) GetTeam(ID int) *team {
	return t.teams[ID]
}

// Delete borra un Team
func (t *tournament) Delete(ID int) {
	delete(t.teams, ID)
}

// Update actualiza un Team del Tournament
func (t *tournament) Update(e team) {
	t.teams[e.ID] = &e
}

// UpdateName actualiza el nombre de un Team
func (e *team) UpdateName(s string) {
	e.name = s
}

// UpdateDt actualiza el DT de un Team
func (e *team) UpdateDt(s string) {
	e.dt = s
}

// UpdatePlayer actualiza un player del Team
func (e *team) UpdatePlayer(s string, n int) {
	e.players[n] = s
}

// Print printea lo que quiero printear
func (t tournament) Print() {
	for _, e := range t.teams {
		fmt.Printf("\t\t %s \nDT: %s \n", e.name)
		e.Print()
	}
}

// Print printea lo que quiero printear
func (e team) Print() {
	fmt.Printf("players: \n")
	for n, v := range e.players {
		fmt.Println(n, v)
	}
}
