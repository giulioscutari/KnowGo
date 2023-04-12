package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const jsonContentType = "application/json"

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() []Player
}
type PlayerServer struct {
	store PlayerStore
	http.Handler
}
type Player struct {
	Name string
	Wins int
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)
	p.store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.LeagueHandler))

	router.Handle("/players/", http.HandlerFunc(p.PlayersHandler))
	p.Handler = router
	return p
}

func (p *PlayerServer) PlayersHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	switch r.Method {
	case http.MethodPost:
		p.ProcessWin(w, player)
	case http.MethodGet:
		p.ShowScore(w, player)

	}
}
func (p *PlayerServer) LeagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonContentType)
	json.NewEncoder(w).Encode(p.store.GetLeague())
}

func (p *PlayerServer) ShowScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, p.store.GetPlayerScore(player))
}

func (p *PlayerServer) ProcessWin(w http.ResponseWriter, player string) {

	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}
