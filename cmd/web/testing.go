package main

import (
	"log"

	"github.com/Lukasloetscher/go_todo_and_events/pkg/postgresql"
)

func test() {

	m, err := postgresql.Initialise_postgresql()
	if err != nil {
		log.Fatal(err)
	}
	col := []string{"ind", "name", "kuerzel"}
	m.Read_table("test", "src", col, nil)

}
