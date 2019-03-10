package templates

//######################################################################

const commentEditTemplate = `<h1><a href="/topics?id={{ .Topic.ID }}">{{ .Topic.Title }}</a></h1>
<p>{{ .Topic.Content }}</p>
<form method="POST">
    <input type="hidden" name="id" value="{{ .Comment.ID }}">
	<input type="hidden" name="tid" value="{{ .Topic.ID }}">
	<textarea name="content" rows="12">{{ .Comment.Content }}</textarea><br />
	<input type="submit" name="action" value="update">
</form>
`
