package main

import (
    "os"
    "text/template"
)

func main() {
    t1 := template.New("test1")
    tmpl, _ := t1.Parse(
`
{{- define "T1"}}ONE {{println .}}{{end}}
{{- define "T2"}}{{template "T1" $}}{{end}}
{{- template "T2" . -}}
`)
    _ = tmpl.Execute(os.Stdout, "hello world")
}