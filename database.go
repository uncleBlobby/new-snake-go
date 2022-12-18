package main

import (
	"database/sql"
	"encoding/json"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDB() *sql.DB {
	log.Printf("Connecting to database...")
	db, err := sql.Open("sqlite3", "./games.db")
	if err != nil {
		panic(err)
	}
	return db
}

func CreateGamesTable(db *sql.DB) {
	log.Printf("Creating table...")
	sqlStmt := `
	create table IF NOT EXISTS games (id integer not null primary key, game_id text, turn text, game_data text, move_response text);`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		panic(err)
	}
}

func WriteTurnToDB(db *sql.DB, gs GameState, mr string) {
	log.Printf("Writing turn to database...")
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	stmt, err := tx.Prepare("insert into games(game_id, turn, game_data, move_response) values(?, ?, ?, ?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	gsJSON, err := json.Marshal(gs)
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(gs.Game.ID, gs.Turn, gsJSON, mr)
	if err != nil {
		panic(err)
	}
	tx.Commit()
}

func CloseDB(db *sql.DB) {
	db.Close()
}
