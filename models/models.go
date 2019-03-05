package models

//######################################################################

import (
	"appleforum/models/db"
)

// Migration setup database
func Migration() {
	db.Exec(`DROP TABLE IF EXISTS users;`)
	db.Exec(`CREATE TABLE users (
	id integer PRIMARY KEY AUTOINCREMENT,
    name varchar(50) NOT NULL,
    passwd varchar(50) NOT NULL,
    email varchar(30) DEFAULT '');`)
	db.Exec(`CREATE TABLE sessions (
	id integer PRIMARY KEY AUTOINCREMENT,
	session_id varchar(50) NOT NULL,
	user_id integer NOT NULL
	);`)
}
