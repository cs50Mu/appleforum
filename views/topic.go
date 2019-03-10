package views

import (
	"appleforum/models/db"
	"appleforum/templates"
	"net/http"
	"time"
)

//######################################################################

// TopicEditHandler handles group edit/add action
var TopicEditHandler = LoginRequiredHandler(func(w http.ResponseWriter, r *http.Request, session *Session) {
	// if POST
	if r.Method == http.MethodPost {
		// when create new topic
		groupID := r.FormValue("gid")
		// when edit topic
		topicID := r.FormValue("id")
		title := r.PostFormValue("title")
		content := r.PostFormValue("content")
		action := r.PostFormValue("action")
		userID := session.UserID
		if action == "create" {
			db.Exec(`INSERT INTO topics (uid, title, content, group_id, created_at, updated_at) 
		VALUES (?, ?, ?, ?, ?, ?)`, userID, title, content, groupID, time.Now(), time.Now())
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		db.Exec(`UPDATE topics SET title = ?, content = ?, updated_at = ? WHERE id = ?`,
			title, content, time.Now(), topicID)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	// else render form
	topicID := r.FormValue("id")
	type Topic struct {
		ID      int64
		Title   string
		Content string
	}
	var topic Topic
	isUpdate := false
	if topicID != "" {
		isUpdate = true
		err := db.QueryRow("SELECT id, title, content FROM topics WHERE id = ?", topicID).Scan(&topic.ID, &topic.Title, &topic.Content)
		if err != nil {
			// handle err
		}
	}
	templates.Render(w, "topicEdit", map[string]interface{}{
		"Topic":    topic,
		"IsUpdate": isUpdate,
	})
})

// TopicListHandler handles topic display
var TopicListHandler = SessionHandler(func(w http.ResponseWriter, r *http.Request, session *Session) {
	//
	topicID := r.FormValue("id")
	type Topic struct {
		ID      int64
		Title   string
		Content string
	}
	var topic Topic
	err := db.QueryRow("SELECT id, title, content FROM topics WHERE id = ?", topicID).Scan(&topic.ID, &topic.Title, &topic.Content)
	if err != nil {
		// handle err
	}
	type Comment struct {
		ID      int64
		Content string
	}
	var comment Comment
	comments := make([]Comment, 0)
	rows := db.Query("SELECT id, content FROM comments WHERE topic_id = ?", topicID)
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&comment.ID, &comment.Content)
		if err != nil {

		}
		comments = append(comments, comment)
	}
	templates.Render(w, "topicList", map[string]interface{}{
		"Topic":    topic,
		"Comments": comments,
	})
})
