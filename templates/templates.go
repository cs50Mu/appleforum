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
}

// Render render a template
func Render(w io.Writer, tmpl string, data interface{}) {
	err := tmpls[tmpl].Execute(w, data)
	if err != nil {
		log.Panicf("render template error: %v\n", err)
	}
}
