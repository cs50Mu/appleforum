package templates

//######################################################################

const groupEditTemplate = `	<h1>Add Group</h1>
	<form method="POST">
		<label>Groupname:</label><br />
		<input type="text" name="group_name"><br />
		<label>Description:</label><br />
		<input type="text" name="description"><br />
		<label>Announcement:</label><br />
		<input type="text" name="announcement"><br />
		<input type="submit">
	</form>`

const groupListTemplate = `<h1>Group List</h1>
    <ul>
		{{ range .Topics }}
        <li><a href="/topics?id={{ .ID }}">{{ .Title }}</a></li>
		{{ else }}
		no topic to display
		{{ end }}
    </ul>
	<a href="/topics/add?gid={{ .GroupID }}">New topic</a>`
