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

	if err := createSchema(db); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	/*
	 */
	service.AddTeam("BOKITA")

	s := service.AddPlayer("andrada", "1", "BOKITA")
	fmt.Println(s)

	httpService := tournament.NewHTTPTransport(service)

	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	httpService.Register(r)

	r.Run()
}

func createSchema(db *sqlx.DB) error {

	schema1 := (`CREATE TABLE IF NOT EXISTS teams (
		id integer primary key autoincrement,
		name varchar(56) NOT NULL UNIQUE);`)

	schema2 := (`CREATE TABLE IF NOT EXISTS players (
		id integer primary key autoincrement,
		name  varchar(56),
		num   integer,
		team varchar(56));`)

	//FOREIGN KEY (team) REFERENCES teams(name));`

	_, err := db.Exec(schema1)
	if err != nil {
		return err
	}
	_, err = db.Exec(schema2)
	if err != nil {
		return err
	}

	return nil
}
