package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Facundo-Arana/GoLang-TP/internal/config"
	"github.com/Facundo-Arana/GoLang-TP/internal/database"
	"github.com/Facundo-Arana/GoLang-TP/internal/service/tournament"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func main() {

	configFile := flag.String("config", "./config.yaml", "this is the service config")
	flag.Parse()

	conf := config.LoadConfig(*configFile)

	db, err := database.NewDatabase(conf)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	service, _ := tournament.New(db, conf)

	/*
		if err := createSchema(db); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		service.AddTeam(tournament.NewTeam("BOKITA", 1))
	*/

	httpService := tournament.NewHTTPTransport(service)

	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	httpService.Register(r)

	r.Run()
}

func createSchema(db *sqlx.DB) error {

	// aca se guardan nombres de team y su ID
	schema1 := `CREATE TABLE IF NOT EXISTS tournament (
		id integer primary key autoincrement,
		name text);`

	/*
		INTENTE HACER DOS TABLAS Y QUE SE RELACIONEN

		POR HOY LO MENOS HOY NO LO LOGRE

		schema2 := `CREATE TABLE IF NOT EXISTS player (
			id integer primary key autoincrement,
			name  varchar(56),
			num   integer,
			attributeFK FOREIGN KEY (attributeFK) REFERENCES tournament(name));`

	*/
	_, err := db.Exec(schema1 /*schema2*/)
	if err != nil {
		panic(err.Error())
	}

	return nil
}
