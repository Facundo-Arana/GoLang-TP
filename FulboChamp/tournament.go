package main

import "fmt"

// Tournament tiene diferentes Teams
type Tournament struct {
	teams map[int]*Team
}

// NewTournament ... CONSTRUCTOR
func NewTournament() Tournament {
	teams := make(map[int]*Team)
	return Tournament{
		teams,
	}
}

// Team tiene un nombre y diferentes players
type Team struct {
	ID      int
	name    string
	dt      string
	players map[int]string
}

// NewTeam ... CONSTRUCTOR
// @param {n string} es el nombre del Team
// @param {i int} es el ID del Team
func NewTeam(n string, d string, ID int) Team {
	players := make(map[int]string)
	return Team{
		ID,
		n,
		d,
		players,
	}
}

/*
  Añade players a un Team
  @param { s string }  es el nombre del jugador
  @param { v int } es el numero de camiseta

*/
func (e *Team) add(c int, n string) {
	e.players[c] = n
}

/*
	Añade una Team al Tournament
*/
func (t *Tournament) add(e Team) {
	t.teams[e.ID] = &e
}

// GetTeam retorna un Team dado un ID
func (t Tournament) GetTeam(ID int) *Team {
	return t.teams[ID]
}

// Delete borra un Team
func (t *Tournament) Delete(ID int) {
	delete(t.teams, ID)
}

// Update actualiza un Team del Tournament
func (t *Tournament) Update(e Team) {
	t.teams[e.ID] = &e
}

// UpdateName actualiza el nombre de un Team
func (e *Team) UpdateName(s string) {
	e.name = s
}

// UpdateDt actualiza el DT de un Team
func (e *Team) UpdateDt(s string) {
	e.dt = s
}

// UpdatePlayer actualiza el DT de un Team
func (e *Team) UpdatePlayer(s string) {
	e.dt = s
}

// Print printea lo que quiero printear
func (t Tournament) Print() {
	for _, e := range t.teams {
		fmt.Printf("\t\t %s \nDT: %s \n", e.name, e.dt)
		e.Print()
	}
}

// Print printea lo que quiero printear
func (e Team) Print() {
	fmt.Printf("players: \n")
	for n, v := range e.players {
		fmt.Println(n, v)
	}
}
