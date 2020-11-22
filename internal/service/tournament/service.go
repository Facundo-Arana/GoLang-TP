package tournament

import (
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

// Service interface
type Service interface {
	AddTeam(Team) error
	GetTeam(int64) []*Team
	GetAllTeam() []*Team
	DeleteTeam(int64) error
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
func (s service) GetAllTeam() []*Team {
	var list []*Team
	err := s.db.Select(&list, "SELECT * FROM tournament")
	if err != nil {
		fmt.Println(err.Error())
	}
	return list
}

// Get ...
func (s service) GetTeam(i int64) []*Team {
	var t []*Team
	query := `SELECT * FROM tournament WHERE id = (?)`
	err := s.db.Select(&t, query, i)
	if err != nil {
		fmt.Println(err.Error())
	}
	return t
}

// Add ...
func (s service) AddTeam(t Team) error {
	insert := `INSERT INTO tournament (name) VALUES (?)`
	fmt.Printf("Team N° %s", t.Name)
	s.db.MustExec(insert, t.Name)
	return nil
}

// DeleteTeam ...
func (s service) DeleteTeam(i int64) error {
	query := `DELETE FROM tournament WHERE id = (?)`
	s.db.Exec(query, i)
	return nil
}
