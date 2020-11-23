package tournament

import (
	"database/sql"
	"fmt"

	"github.com/Facundo-Arana/GoLang-TP/internal/config"
	"github.com/jmoiron/sqlx"
)

// Team ...
type Team struct {
	ID   int64
	Name string
}

// NewTeam ...
func NewTeam(s string, i int64) Team {
	return Team{
		i,
		s,
	}
}

/*
	// Player ...
	type Player struct {
		ID     int64
		Name   string
		Num    string
		AttrFK string
	}
	// NewPlayer ...
	func NewPlayer(s string, i int64, n string, attr string) Player {
		return Player{
			i,
			s,
			n,
			attr,
		}
	}
*/

// Service interface
type Service interface {
	GetAllTeams() []*Team
	GetTeam(string) []*Team
	AddTeam(Team) sql.Result
	DeleteTeam(string) sql.Result
	EditTeam(string, string) sql.Result
}

type service struct {
	db   *sqlx.DB
	conf *config.Config
}

// New ...
func New(db *sqlx.DB, c *config.Config) (Service, error) {
	return service{db, c}, nil
}

// GetAllTeams ...
func (s service) GetAllTeams() []*Team {
	var list []*Team
	query := "SELECT * FROM tournament"
	err := s.db.Select(&list, query)
	if err != nil {
		fmt.Println(err.Error())
	}
	return list
}

// Get ...
func (s service) GetTeam(i string) []*Team {
	var t []*Team
	query := `SELECT * FROM tournament WHERE id = (?)`
	err := s.db.Select(&t, query, i)
	if err != nil {
		fmt.Println(err.Error())
	}
	return t
}

// Add ...
func (s service) AddTeam(t Team) sql.Result {
	query := `INSERT INTO tournament (name) VALUES (?)`
	return s.db.MustExec(query, t.Name)
}

// DeleteTeam ...
func (s service) DeleteTeam(i string) sql.Result {
	query := `DELETE FROM tournament WHERE id = (?)`
	return s.db.MustExec(query, i)
}

// EditTeam ...
func (s service) EditTeam(n string, i string) sql.Result {
	query := `UPDATE tournament SET name = ? WHERE id = ?`
	return s.db.MustExec(query, n, i)
}

/*
// AddPlayer ...
func (s service) AddPlayer(p Player) error {
	query := `INSERT INTO player (name, num, attributeFK) VALUES (?, ?, ?)`
	s.db.MustExec(query, p.Name, p.Num, p.AttrFK)
	return nil
}

// GetAllPlayers ...
func (s service) GetAllPlayers(n string) []*Player {
	var list []*Player
	query := "SELECT * FROM player "
	err := s.db.Select(&list, query, n)
	if err != nil {
		fmt.Println(err.Error())
	}
	return list
}
*/
