package tournament

import (
	"net/http"
	"strconv" // lo comento pa que no se queje

	"github.com/gin-gonic/gin" // comentar porque sino se pone todo amarillo
)

// HTTPService ...
type HTTPService interface {
	Register(*gin.Engine)
}

type endpoint struct {
	method   string
	path     string
	function gin.HandlerFunc
}

type httpService struct {
	endpoints []*endpoint
}

// NewHTTPTransport ...
func NewHTTPTransport(s Service) HTTPService {
	endpoints := makeEndpoints(s)
	return httpService{endpoints}
}

func makeEndpoints(s Service) []*endpoint {
	list := []*endpoint{}

	// añadir un team
	list = append(list, &endpoint{
		method:   "POST",
		path:     "/team",
		function: add(s),
	})

	// obtener todos los teams
	list = append(list, &endpoint{
		method:   "GET",
		path:     "/team",
		function: getAll(s),
	})

	// obtener un unico team
	list = append(list, &endpoint{
		method:   "GET",
		path:     "/team/:ID",
		function: get(s),
	})

	// borrar un team
	list = append(list, &endpoint{
		method:   "DELETE",
		path:     "/team/:ID",
		function: delete(s),
	})

	// editar un team
	list = append(list, &endpoint{
		method:   "PUT",
		path:     "/team/:ID",
		function: edit(s),
	})

	// añadir un player a team
	list = append(list, &endpoint{
		method:   "POST",
		path:     "/team/:ID/player",
		function: addPlayer(s),
	})

	// obtener los players de team
	list = append(list, &endpoint{
		method:   "GET",
		path:     "/team/:ID/player",
		function: getPlayersByTeam(s),
	})

	// obtener todos los players.... endpoint solo para desarrollo
	list = append(list, &endpoint{
		method:   "GET",
		path:     "/player",
		function: getAllPlayers(s),
	})

	return list
}

func getAllPlayers(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		players, err := s.GetAllPlayers()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{":( ": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"players": players,
		})
	}
}

func getPlayersByTeam(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		i, _ := strconv.ParseInt(c.Param("ID"), 10, 64)

		players, err := s.GetPlayersByTeam(i)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{":( ": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"players": players,
		})
	}
}

func addPlayer(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		i, _ := strconv.ParseInt(c.Param("ID"), 10, 64)
		playerName := c.Query("name")
		num, _ := strconv.ParseInt(c.Query("numero"), 10, 64)

		player, err := s.AddPlayer(NewPlayer(playerName, num, i))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{":( ": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"player": player,
		})
	}
}

func getAll(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		teams, err := s.GetAllTeams()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{":( ": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"teams": teams,
		})
	}
}

func get(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.ParseInt(c.Param("ID"), 10, 64)
		t, err := s.GetTeam(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{":( ": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"team": t,
		})
	}
}

func delete(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		i, _ := strconv.ParseInt(c.Param("ID"), 10, 64)

		err := s.DeleteTeam(i)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{":( ": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	}
}

func add(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := NewTeam(c.Query("name"))
		t, err := s.AddTeam(t)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{":( ": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"team": t,
		})
	}
}

func edit(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		i, _ := strconv.ParseInt(c.Param("ID"), 10, 64)
		n := c.Query("newName")

		err := s.EditTeam(n, i)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{":( ": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	}
}

// Register ...
func (s httpService) Register(r *gin.Engine) {
	for _, e := range s.endpoints {
		r.Handle(e.method, e.path, e.function)
	}
}
