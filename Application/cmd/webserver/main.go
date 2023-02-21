// main.go
package main

import (
	"log"
	"net/http"
	"os"
	"poker"
)

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("Problem opening file %s %v", dbFileName, err)
	}

	store, err := poker.NewFileSystemPlayerStore(db)

	if err != nil {
		log.Fatalf("Problem creating stopre %s %v", store, err)
	}

	game := poker.NewTexasHoldem(poker.BlindAlerterFunc(poker.Alerter), store)

	server, err := poker.NewPlayerServer(store, game)

	if err != nil {

		log.Fatalf("Problem creating PlayerServer %v", err)
	}

	log.Fatal(http.ListenAndServe(":5000", server))

}
