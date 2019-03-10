package templates

//######################################################################
import (
	"html/template"
	"io"
	"log"
)

var tmpls map[string]*template.Template = make(map[string]*template.Template)

func init() {
	tmpls["login"] = template.Must(template.New("login").Parse(loginTemplate))
	tmpls["groupEdit"] = template.Must(template.New("groupEdit").Parse(groupEditTemplate))
	tmpls["groupList"] = template.Must(template.New("groupList").Parse(groupListTemplate))
	tmpls["topicEdit"] = template.Must(template.New("topicEdit").Parse(topicEditTemplate))
	tmpls["topicList"] = template.Must(template.New("topicList").Parse(topicListTemplate))
	tmpls["index"] = template.Must(template.New("index").Parse(indexTemplate))
	tmpls["commentEdit"] = template.Must(template.New("commentEdit").Parse(commentEditTemplate))
}

// Render render a template
func Render(w io.Writer, tmpl string, data interface{}) {
	err := tmpls[tmpl].Execute(w, data)
	if err != nil {
		log.Panicf("render template error: %v\n", err)
	}
}
