package main

import (
	"bytes"
	"flag"
	"html/template"
	"log"
	"net/http"

	"github.com/gregoryv/align"
)

func main() {
	bind := ":9100"
	flag.StringVar(&bind, "b", bind, "")
	flag.StringVar(&bind, "bind", bind, "listen interface and port")
	flag.Parse()

	http.HandleFunc("/", serve)

	log.Print("listen on", bind)
	http.ListenAndServe(bind, nil)
}

func serve(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		a := r.FormValue("a")
		b := r.FormValue("b")
		if len(a) == 0 || len(b) == 0 {

			_ = tpl.ExecuteTemplate(w, "START", "")
			_ = tpl.ExecuteTemplate(w, "ERR", "oups")
			_ = tpl.ExecuteTemplate(w, "END", "")
			return
		}
		result := align.NeedlemanWunsch([]rune(a), []rune(b))
		_ = tpl.ExecuteTemplate(w, "START", "")

		var origins bytes.Buffer
		result.PrintOrigins(&origins)

		var alignment bytes.Buffer
		result.PrintAlignment(&alignment)

		var matrix bytes.Buffer
		result.PrintScoreMatrix(&matrix)

		view := map[string]any{
			"Alignment":   alignment.String(),
			"Origins":     origins.String(),
			"ScoreMatrix": matrix.String(),
			"MaxScore":    result.MaxScore(),
		}

		_ = tpl.ExecuteTemplate(w, "RESULT", view)
		_ = tpl.ExecuteTemplate(w, "END", "")

	default:
		_ = tpl.ExecuteTemplate(w, "START", "")
		_ = tpl.ExecuteTemplate(w, "NEW", "")
		_ = tpl.ExecuteTemplate(w, "END", "")
	}
}

const page = `

{{define "START"}}
<html>
<a href="/">New</a> | <a href="https://en.wikipedia.org/wiki/Needleman%E2%80%93Wunsch_algorithm">
wikipedia Needleman-Wunsch_algorithm
</a>
<br>
<br>
{{end}}


{{define "NEW"}}
<form method="POST">
Sequence A: <input name="a" value="GCATGCG"/>
<br>
Sequence B: <input name="b" value="GATTACA"/>
<button>Align</button>
</form>
{{end}}


{{define "RESULT"}}
<div>
MAX SCORE:
<pre>{{.MaxScore}}</pre>
</div>
<div>
ALIGNMENTS:
<pre>{{.Alignment}}</pre>
</div>
<div>
ORIGINS:
<pre>{{.Origins}}</pre>
</div>
<div>
SCORE MATRIX:
<pre>{{.ScoreMatrix}}</pre>
</div>
<style>
div {
float: left;
margin-right: 4em;
}
</style>
{{end}}


{{define "ERR"}}
Error: {{.}}
{{end}}


{{define "END"}}
<html>
{{end}}

`

var tpl = template.Must(template.New("foo").Parse(page))
