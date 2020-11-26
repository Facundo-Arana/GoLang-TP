package tournament

import (
	"errors"

	"github.com/Facundo-Arana/GoLang-TP/internal/config"
	"github.com/jmoiron/sqlx"
)

/*Service interface

Ya se que no deberia recibir parametros de tipo string pero no encontre la forma de
poder convertir los parametros de la request en un estructura
*/
type Service interface {
	AddTeam(*Team) (*Team, error)
	GetAllTeams() ([]*Team, error)
	GetTeam(int64) (*Team, error)
	DeleteTeam(int64) error
	EditTeam(string, int64) error

	AddPlayer(*Player) (*Player, error)
	GetPlayersByTeam(int64) ([]*Player, error)
	GetAllPlayers() ([]*Player, error)
}

// Team ...
type Team struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`

	//Players map[int]*Player
}

// Player ...
type Player struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
	Num  int64  `db:"num"`
	Team int64  `db:"idteamFK"`
}

// NewTeam ...
func NewTeam(s string) *Team {
	return &Team{
		0,
		s,
	}
}

// NewPlayer ...
func NewPlayer(n string, i int64, t int64) *Player {
	return &Player{
		0,
		n,
		i,
		t,
	}
}

type service struct {
	db   *sqlx.DB
	conf *config.Config
}

// New ...
func New(db *sqlx.DB, c *config.Config) (Service, error) {
	return service{db, c}, nil
}

func (s service) GetAllPlayers() ([]*Player, error) {
	query := `SELECT * FROM players`
	var list []*Player
	err := s.db.Select(&list, query)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (s service) GetPlayersByTeam(i int64) ([]*Player, error) {

	var list []*Player

	//`SELECT players.* FROM players JOIN teams ON player.teamFk = tournament.name WHERE tournament.name = (?)`
	query := `SELECT * FROM players WHERE idteamFK = (?)`
	err := s.db.Select(&list, query, i)
	if err != nil {
		return nil, err
	}

	return list, nil
}

// AddPlayer ...
func (s service) AddPlayer(p *Player) (*Player, error) {
	query := `INSERT INTO players (name, num, idteamFK) VALUES (?, ?, ?)`

	res, err := s.db.Exec(query, p.Name, p.Num, p.Team)
	if err != nil {
		return nil, err
	}

	LastID, _ := res.LastInsertId()
	p.ID = LastID

	return p, nil
}

// GetAllTeams ...
func (s service) GetAllTeams() ([]*Team, error) {
	var list []*Team

	query := "SELECT * FROM teams"
	err := s.db.Select(&list, query)
	if err != nil {
		return nil, err
	}
	return list, nil
}

// Get ...
func (s service) GetTeam(i int64) (*Team, error) {
	query := `SELECT * FROM teams WHERE id = (?)`

	var t Team
	err := s.db.Get(&t, query, i)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// Add ...
func (s service) AddTeam(t *Team) (*Team, error) {

	query := `INSERT INTO teams (name) VALUES (?)`

	res, err := s.db.Exec(query, t.Name)
	if err != nil {
		return nil, errors.New("DATABASE ERROR - " + err.Error())
	}
	LastID, _ := res.LastInsertId()
	t.ID = LastID
	return t, nil

}

// DeleteTeam ...
func (s service) DeleteTeam(i int64) error {
	query := `DELETE FROM teams WHERE id = (?)`
	res, err := s.db.Exec(query, i)
	if err != nil {
		return errors.New("DATABASE ERROR - " + err.Error())
	}
	col, _ := res.RowsAffected()
	if col == 0 {
		return errors.New("non-existent team")
	}
	return nil
}

// EditTeam ...
func (s service) EditTeam(n string, i int64) error {
	query := `UPDATE teams SET name = ? WHERE id = ?`
	_, err := s.db.Exec(query, n, i)

	if err != nil {
		return errors.New("DATABASE ERROR - " + err.Error())
	}

	return nil
}
