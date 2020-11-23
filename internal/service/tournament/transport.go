package tournament

import (
	"net/http"

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

	return list
}

func getAll(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"teams": s.GetAllTeams(),
		})
	}
}

func get(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		i := c.Param("ID")
		c.JSON(http.StatusOK, gin.H{
			"team": s.GetTeam(i),
		})
	}
}

func delete(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		i := c.Param("ID")
		c.JSON(http.StatusOK, gin.H{
			"team": s.DeleteTeam(i),
		})
	}
}

func add(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		n := c.Query("name")
		c.JSON(http.StatusOK, gin.H{
			"team": s.AddTeam(n),
		})
	}
}

func edit(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		i := c.Param("ID")
		n := c.Query("newName")
		c.JSON(http.StatusOK, gin.H{
			"team": s.EditTeam(n, i),
		})
	}
}

/*
func addPlayer(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		n := c.Param("name")
		p := c.Param("playerName")
		num := c.Param("numeroCamiseta")
		c.JSON(http.StatusOK, gin.H{
			"players": s.AddPlayer(NewPlayer(p, 0, num, n)),
		})
	}
}

func getAllPlayers(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		n := c.Param("ID")
		c.JSON(http.StatusOK, gin.H{
			"players": s.GetAllPlayers(n),
		})
	}
}
*/

// Register ...
func (s httpService) Register(r *gin.Engine) {
	for _, e := range s.endpoints {
		r.Handle(e.method, e.path, e.function)
	}
}
