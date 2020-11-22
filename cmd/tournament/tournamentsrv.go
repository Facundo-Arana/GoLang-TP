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

			//service.AddTeam(tournament.NewTeam("BOKITA", 1))

			m := service.GetAllTeam()
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			for _, h := range m {
				fmt.Println(h)
			}
	*/
	httpService := tournament.NewHTTPTransport(service)

	r := gin.Default()
	httpService.Register(r)
	r.Run()
}

//estas serian LAS tablas que quiero armar
func createSchema(db *sqlx.DB) error {
	schema1 := `CREATE TABLE IF NOT EXISTS tournament (
		id integer primary key autoincrement,
		name text);`
	/*
		BAJANDO LAS ESPECTATIVAS


		attrFK integer NOT NULL,
		FOREIGN KEY (attrFK) REFERENCES team(id));`

		schema2 := `CREATE TABLE IF NOT EXISTS team (
			id integer primary key autoincrement,
			name  varchar(56),
			attributeFK FOREIGN KEY (attributeFK) REFERENCES player(id));`

		schema3 := `CREATE TABLE IF NOT EXISTS player (
			id integer primary key,
			name  varchar(56));`
	*/
	_, err := db.Exec(schema1 /*,schema2,schema3*/)
	if err != nil {
		panic(err.Error())
	}

	return nil
}
