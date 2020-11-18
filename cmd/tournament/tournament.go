package main

import (
	"flag"

	"github.com/Facundo-Arana/GoLang-TP/internal/config"

	"github.com/Facundo-Arana/GoLang-TP/internal/service/tournament"
)

func main() {
	configFile := flag.String("config", "./config.yaml", "this is the service config")
	flag.Parse()

	conf := config.LoadConfig(*configFile)

	service := tournament.NewTournament()
	http := tournament.NewHTTPTournament(service, conf)

	http.Run()
}
