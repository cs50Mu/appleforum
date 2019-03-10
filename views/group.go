package views

import (
	"appleforum/models/db"
	"appleforum/templates"
	"log"
	"net/http"
	"time"
)

//######################################################################

// GroupEditHandler handles group edit/add action
var GroupEditHandler = LoginRequiredHandler(func(w http.ResponseWriter, r *http.Request, session *Session) {
	// if POST
	if r.Method == http.MethodPost {
		groupName := r.PostFormValue("group_name")
		description := r.PostFormValue("description")
		announcement := r.PostFormValue("announcement")
		userID := session.UserID
		db.Exec(`INSERT INTO groups (uid, name, description, announcement, created_at, updated_at) 
		VALUES (?, ?, ?, ?, ?, ?)`, userID, groupName, description, announcement, time.Now(), time.Now())
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	// else render form
	templates.Render(w, "groupEdit", nil)
})

// GroupListHandler handles group list
var GroupListHandler = SessionHandler(func(w http.ResponseWriter, r *http.Request, session *Session) {
	groupName := r.FormValue("name")
	var gid int64
	err := db.QueryRow("SELECT id FROM groups WHERE name = ?", groupName).Scan(&gid)
	if err != nil {
		http.NotFound(w, r)
	}
	log.Printf("gid: %v", gid)
	type Topic struct {
		ID    int64
		Title string
	}
	var topic Topic
	topics := make([]Topic, 0)
	rows := db.Query("SELECT id, title FROM topics WHERE group_id = ?", gid)
	defer rows.Close()
	for rows.Next() {
		// how to conver a database row into a struct
		// refer: https://stackoverflow.com/questions/17265463/how-do-i-convert-a-database-row-into-a-struct-in-go
		err = rows.Scan(&topic.ID, &topic.Title)
		if err != nil {
			// empty result
		}
		log.Printf("topic: %+v", topic)
		topics = append(topics, topic)
	}
	templates.Render(w, "groupList", map[string]interface{}{
		"Topics":  topics,
		"GroupID": gid,
	})
})
