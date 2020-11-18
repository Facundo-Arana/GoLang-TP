package afa

import (
	"net/http"

	"github.com/Facundo-Arana/GoLang-TP/internal/config"
	"github.com/gin-gonic/gin"
)

// HTTPTournament ...
type HTTPTournament interface {
	Run()
}

type httpService struct {
	r *gin.Engine
}

// NewHTTPTournament returns a new instance of HTTPTournament
func NewHTTPTournament(t Tournament, conf *config.Config) HTTPTournament {
	r := gin.Default()
	group := r.Group(conf.Tournament.Version)
	group.GET("/funcionaPorFavor", unaFuncion(t))
	return httpService{r}
}

func unaFuncion(t Tournament) gin.HandlerFunc {
	return func(c *gin.Context) {
		s := c.Query("name")
		c.JSON(http.StatusOK, gin.H{
			"algo": t.unaFuncion(s),
		})
	}
}

func (s httpService) Run() {
	s.r.Run()
}

/*
func main() {

	r := gin.Default()
	r.GET("/team/:name", showTeams)
	r.POST("/team", addTeam)
	r.Run()
}

func showTeams(c *gin.Context) {
	name := c.Param("name")
	lastname := c.Query("lastname")
	c.JSON(200, gin.H{
		"name": name + " " + lastname,
	})

}

func addTeam(c *gin.Context) {
	c.JSON(201, gin.H{
		"name": "name",
	})
}

*/
