package templates

//######################################################################

const indexTemplate = `<h1>Index</h1>
<h2>Groups</h2>
    <ul>
		{{ range .Groups }}
        <li><a href="/groups?name={{ .Name }}">{{ .Name }}</a></li>
		{{ else }}
		no group to display
		{{ end }}
    </ul>
	<a href="/groups/add">New Group</a>
	<h2>Recent Topics</h2>
    <ul>
		{{ range .Topics }}
        <li><a href="/topics?id={{ .ID }}">{{ .Title }}</a></li>
		{{ else }}
		no topic to display
		{{ end }}
    </ul>
	`
