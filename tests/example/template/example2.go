package main

import (
	"html/template"

)

func main()  {
	tx := template.Must(template.New("hh").Parse(
		`{{range $x := . -}}
		{{println $x}}
		{{- end}}
		`))
		s := []int{11, 22, 33, 44, 55}
		_ = tx.Execute(os.Stdout, s)
}