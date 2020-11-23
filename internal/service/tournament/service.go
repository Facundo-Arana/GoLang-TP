package tournament

import (
	"errors"
	"fmt"

	"github.com/Facundo-Arana/GoLang-TP/internal/config"
	"github.com/jmoiron/sqlx"
)

// Team ...
type Team struct {
	ID   int64
	Name string
}

// Player ...
type Player struct {
	ID   int64
	Name string
	Num  string
	IDFk int64
}

/*
// NewTeam ...
func NewTeam(s string, i int64) Team {
	return Team{
		i,
		s,
	}
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
	AddTeam(string) string
	GetAllTeams() []*Team
	GetTeam(string) *Team
	DeleteTeam(string) string
	EditTeam(string, string) string
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
func (s service) GetTeam(i string) *Team {
	query := `SELECT * FROM tournament WHERE id = (?)`

	var t Team
	err := s.db.Get(&t, query, i)
	if err != nil {
		panic(err.Error())
	}
	return &t
}

// Add ...
func (s service) AddTeam(t string) string {
	query := `INSERT INTO tournament (name) VALUES (?)`

	res, err := s.db.Exec(query, t)
	if err != nil {
		return fmt.Sprintf("%v", errors.New("DATABASE ERROR - "+err.Error()))
	}
	LastID, _ := res.LastInsertId()

	return fmt.Sprintf("New team ID: %d", LastID)
}

// DeleteTeam ...
func (s service) DeleteTeam(i string) string {
	query := `DELETE FROM tournament WHERE id = (?)`
	res, err := s.db.Exec(query, i)

	if err != nil {
		return fmt.Sprintf("%v", errors.New("DATABASE ERROR - "+err.Error()))
	}
	RowsAffected, _ := res.RowsAffected()

	return fmt.Sprintf("Columns affected: %d", RowsAffected)
}

// EditTeam ...
func (s service) EditTeam(n string, i string) string {
	query := `UPDATE tournament SET name = ? WHERE id = ?`
	res, err := s.db.Exec(query, n, i)

	if err != nil {
		return fmt.Sprintf("%v", errors.New("DATABASE ERROR - "+err.Error()))
	}
	RowsAffected, _ := res.RowsAffected()

	return fmt.Sprintf("Columns affected: %d", RowsAffected)
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
