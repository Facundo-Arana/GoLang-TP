package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Facundo-Arana/GoLang-TP/internal/config"
	"github.com/Facundo-Arana/GoLang-TP/internal/database"
	"github.com/Facundo-Arana/GoLang-TP/internal/service/afa"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func main() {
	configFile := flag.String("config", "./config.yaml", "this is the service config")
	flag.Parse()

	cfg := config.LoadConfig(*configFile)

	db, err := database.NewDatabase(cfg)
	defer db.Close()

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	service, _ := afa.New(db, cfg)
	httpService := afa.NewHTTPTransport(service)

	r := gin.Default()
	httpService.Register(r)
	r.Run()
}

//esta seria la tabla que quiero armar
func createSchema(db *sqlx.DB) {
	schema1 := `CREATE TABLE IF NOT EXISTS tournament (
		id integer primary key autoincrement,
		name  varchar(56),
		attrFK integer NOT NULL,
		FOREIGN KEY (attrFK) REFERENCES team(id));`
	schema2 := `CREATE TABLE IF NOT EXISTS team (
		id integer primary key autoincrement,
		name  varchar(56),
		attributeFK FOREIGN KEY (attributeFK) REFERENCES player(id));`
	schema3 := `CREATE TABLE IF NOT EXISTS player (
		id integer primary key,
		name  varchar(56));`

	_, err := db.Exec(schema1, schema2, schema3)
	if err != nil {
		panic(err.Error())
	}
}

/*
func createSchema(db *sqlx.DB) error {
	schema := `CREATE TABLE IF NOT EXISTS messages (
		id integer primary key autoincrement,
		text varchar);`

	// execute a query on the server
	_, err := db.Exec(schema)
	if err != nil {
		return err
	}

	// or, you can use MustExec, which panics on error
	insertMessage := `INSERT INTO messages (text) VALUES (?)`
	s := fmt.Sprintf("Message number %v", time.Now().Nanosecond())
	db.MustExec(insertMessage, s)
	return nil
}
*/
