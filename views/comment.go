package views

import (
	"appleforum/models/db"
	"appleforum/templates"
	"net/http"
	"time"
)

//######################################################################

// CommentEditHandler handles comment
var CommentEditHandler = LoginRequiredHandler(func(w http.ResponseWriter, r *http.Request, session *Session) {
	if r.Method == http.MethodPost {
		topicID := r.FormValue("tid")
		content := r.PostFormValue("content")
		action := r.PostFormValue("action")
		commentID := r.PostFormValue("id")
		userID := session.UserID
		if action == "update" {
			db.Exec("UPDATE comments SET content = ?, updated_at = ? WHERE id = ?", content,
				time.Now(), commentID)
			http.Redirect(w, r, "/topics?id="+topicID, http.StatusSeeOther)
			return
		}
		db.Exec(`INSERT INTO comments (uid, content, topic_id, created_at, updated_at) 
	VALUES (?, ?, ?, ?, ?)`, userID, content, topicID, time.Now(), time.Now())
		http.Redirect(w, r, "/topics?id="+topicID, http.StatusSeeOther)
		return
	}

	commentID := r.FormValue("id")
	type Comment struct {
		ID      int64
		Content string
	}
	var comment Comment
	var topicID int64
	err := db.QueryRow("SELECT id, content, topic_id FROM comments WHERE id = ?", commentID).Scan(&comment.ID,
		&comment.Content, &topicID)
	if err != nil {

	}
	type Topic struct {
		ID      int64
		Title   string
		Content string
	}
	var topic Topic
	err = db.QueryRow("SELECT id, title, content FROM topics WHERE id = ?", topicID).Scan(&topic.ID,
		&topic.Title, &topic.Content)
	if err != nil {

	}
	templates.Render(w, "commentEdit", map[string]interface{}{
		"Comment": comment,
		"Topic":   topic,
	})

})
