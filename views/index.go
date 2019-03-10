package views

import (
	"appleforum/models/db"
	"appleforum/templates"
	"log"
	"net/http"
)

//######################################################################

// IndexHandler handles index
var IndexHandler = SessionHandler(func(w http.ResponseWriter, r *http.Request, session *Session) {
	type Group struct {
		Name string
		Desc string
	}
	var group Group
	groups := make([]Group, 0)
	rows := db.Query("SELECT name, description FROM groups ORDER BY created_at DESC;")
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&group.Name, &group.Desc)
		if err != nil {
			// empty
		}
		groups = append(groups, group)
	}

	type Topic struct {
		ID    int64
		Title string
	}
	var topic Topic
	topics := make([]Topic, 0)
	rows = db.Query("SELECT id, title FROM topics ORDER BY created_at DESC;")
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&topic.ID, &topic.Title)
		if err != nil {
			// empty result
		}
		log.Printf("topic: %+v", topic)
		topics = append(topics, topic)
	}
	templates.Render(w, "index", map[string]interface{}{
		"Groups": groups,
		"Topics": topics,
	})
})
