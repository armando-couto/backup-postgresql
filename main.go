package main

import "github.com/armando-couto/barkup"

func main() {

	postgres := &barkup.Postgres{
		Host: "127.0.0.1",
		Port: "5432",
		DB:   "banco_de_dados",

		// Not necessary if the program runs as an authorized pg user/role
		Username:  "postgres",

		// Caminho de onde está instalado o PostgreSQL
		PGDumpCmd: "/Library/PostgreSQL/11/bin/pg_dump",

		// Any extra pg_dump options
		Options: []string{"--no-owner"},
	}

	// Writes a file `./bu_DBNAME_TIMESTAMP.sql.tar.gz`
	result := postgres.Export()

	if result.Error != nil {
		panic(result.Error)
	}
}
