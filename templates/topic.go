package templates

//######################################################################

const topicEditTemplate = `<h1>Add / Update Topic</h1>
	<form method="POST">
		<label>Title:</label><br />
		<input type="text" name="title" value="{{ .Topic.Title }}"><br />
		<label>Content:</label><br />
		<input type="text" name="content" value="{{ .Topic.Content }}"><br />
		<input type="submit" name="action" value="{{ if .IsUpdate }}update{{ else }}create{{ end }}">
	</form>`

const topicListTemplate = `<h1>{{ .Topic.Title }}</h1>
<hr >
<p>{{ .Topic.Content }}</p>
<a href="/topics/edit?id={{ .Topic.ID }}">Edit Topic</a>
<h2>Comments</h2>
{{ range .Comments }}
<p>{{ .Content }}</p>
<p><a href="/comments/edit?id={{ .ID }}">Edit</a></p>
<hr >
{{ else }}
no comment right now
{{ end }}
<form action="/comments/add?tid={{ .Topic.ID }}" method="POST">
	<textarea name="content" rows="12" placeholder="Your comment..."></textarea><br />
	<input type="submit" name="action">
</form>
`
