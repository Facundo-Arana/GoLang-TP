package tournament

import (
	"errors"
	"fmt"

	"github.com/Facundo-Arana/GoLang-TP/internal/config"
	"github.com/jmoiron/sqlx"
)

// Service interface
type Service interface {
	AddTeam(string) string
	GetAllTeams() []*Team
	GetTeam(string) *Team
	DeleteTeam(string) string
	EditTeam(string, string) string

	AddPlayer(string, string, string) string
	GetPlayersByTeam(string) []*Player
	GetAllPlayers() []*Player
}

// Team ...
type Team struct {
	ID   int64
	Name string
}

// Player ...
type Player struct {
	ID   int64
	Name string
	Num  int64
	Team string
}

type service struct {
	db   *sqlx.DB
	conf *config.Config
}

// New ...
func New(db *sqlx.DB, c *config.Config) (Service, error) {
	return service{db, c}, nil
}

func (s service) GetAllPlayers() []*Player {
	query := `SELECT * FROM players`
	var list []*Player

	err := s.db.Select(&list, query)
	if err != nil {
		panic(err.Error())
	}
	return list
}

func (s service) GetPlayersByTeam(n string) []*Player {

	var list []*Player
	/*
		query := `SELECT players.* FROM players JOIN teams ON player.teamFk = tournament.name WHERE tournament.name = (?)`
			err := s.db.Select(&list, query, n)
			if err != nil {
				panic(err.Error())
			}
	*/
	return list
}

// AddPlayer ...
func (s service) AddPlayer(n string, num string, t string) string {
	query := `INSERT INTO players (name, num, team) VALUES (?, ?, ?)`

	res := s.db.MustExec(query, n, num, t)
	LastID, _ := res.LastInsertId()

	return fmt.Sprintf("New Player ID: %d", LastID)
}

// GetAllTeams ...
func (s service) GetAllTeams() []*Team {
	var list []*Team

	query := "SELECT * FROM teams"
	err := s.db.Select(&list, query)
	if err != nil {
		panic(err.Error())
	}
	return list
}

// Get ...
func (s service) GetTeam(i string) *Team {
	query := `SELECT * FROM teams WHERE id = (?)`

	var t Team
	err := s.db.Get(&t, query, i)
	if err != nil {
		panic(err.Error())
	}
	return &t
}

// Add ...
func (s service) AddTeam(t string) string {
	query := `INSERT INTO teams (name) VALUES (?)`

	res, err := s.db.Exec(query, t)
	if err != nil {
		return fmt.Sprintf("%v", errors.New("DATABASE ERROR - "+err.Error()))
	}
	LastID, _ := res.LastInsertId()

	return fmt.Sprintf("New team ID: %d", LastID)
}

// DeleteTeam ...
func (s service) DeleteTeam(i string) string {
	query := `DELETE FROM teams WHERE id = (?)`
	res, err := s.db.Exec(query, i)

	if err != nil {
		return fmt.Sprintf("%v", errors.New("DATABASE ERROR - "+err.Error()))
	}
	RowsAffected, _ := res.RowsAffected()

	return fmt.Sprintf("Columns affected: %d", RowsAffected)
}

// EditTeam ...
func (s service) EditTeam(n string, i string) string {
	query := `UPDATE teams SET name = ? WHERE id = ?`
	res, err := s.db.Exec(query, n, i)

	if err != nil {
		return fmt.Sprintf("%v", errors.New("DATABASE ERROR - "+err.Error()))
	}
	RowsAffected, _ := res.RowsAffected()

	return fmt.Sprintf("Columns affected: %d", RowsAffected)
}
