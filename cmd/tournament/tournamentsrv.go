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
	//    go run cmd/tournament/tournamentsrv.go -config ./config/config.yaml

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
		idteamFK integer NOT NULL );`)
	//  FOREIGN KEY (idteamFK) REFERENCES teams(id) ON DELETE CASCADE);`

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
