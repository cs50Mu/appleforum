package models

//######################################################################

import (
	"appleforum/models/db"
)

// Migration setup database
func Migration() {
	//	db.Exec(`DROP TABLE IF EXISTS users;`)
	// user
	db.Exec(`CREATE TABLE users (
		id integer PRIMARY KEY AUTOINCREMENT,
	    name varchar(50) NOT NULL,
	    passwd varchar(50) NOT NULL,
	    email varchar(30) DEFAULT '');`)

	// session
	db.Exec(`CREATE TABLE sessions (
		id integer PRIMARY KEY AUTOINCREMENT,
		session_id varchar(50) NOT NULL,
		user_id integer NOT NULL
		);`)

	// group
	db.Exec(`CREATE TABLE groups (
	    id integer PRIMARY KEY AUTOINCREMENT,
	    uid integer NOT NULL,
	    name varchar(50) NOT NULL,
	    description varchar(50) NOT NULL,
	    announcement text NOT NULL,
	    created_at integer NOT NULL,
	    updated_at integer NOT NULL);`)

	// topic
	db.Exec(`CREATE TABLE topics (
	    id integer PRIMARY KEY AUTOINCREMENT,
	    uid integer NOT NULL,
	    title varchar(50) NOT NULL,
	    content text NOT NULL,
	    group_id integer NOT NULL,
	    created_at integer NOT NULL,
	    updated_at integer NOT NULL);`)

	// comment
	db.Exec(`DROP TABLE IF EXISTS comments;`)
	db.Exec(`CREATE TABLE comments (
    id integer PRIMARY KEY AUTOINCREMENT,
    uid integer NOT NULL,
    content text NOT NULL,
    topic_id integer NOT NULL,
    created_at integer NOT NULL,
    updated_at integer NOT NULL);`)
}
