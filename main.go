package main

import (
	"html/template"
	"math/rand"
	"net/http"
)

type Game struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Players int    `json:"players"`
}

var games = []Game{
	{ID: 1, Name: "Monopoly", Players: 2},
	{ID: 2, Name: "Chess", Players: 2},
	{ID: 3, Name: "Scrabble", Players: 2},
	{ID: 4, Name: "Risk", Players: 2},
}

var tmpl = template.Must(template.ParseFiles("templates/games.html", "templates/game-info.html"))

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/games", gamesHandler)
	http.HandleFunc("/change-game", changeGameHandler)

	http.ListenAndServe(":8000", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Hello, World!"))
}

func gamesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	game := games[rand.Intn(len(games))]

	tmpl.ExecuteTemplate(w, "game.html", game)
}

func changeGameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	game := games[rand.Intn(len(games))]

	w.Header().Set("Content-Type", "text/html")
	tmpl.ExecuteTemplate(w, "game-info.html", game)
}
