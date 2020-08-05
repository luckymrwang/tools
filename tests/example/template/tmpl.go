package main

import (
	"os"
	"text/template"
)

type Friend struct {
	Fname string
}
type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

func main() {
	f1 := Friend{Fname: "xiaofang"}
	f2 := Friend{Fname: "wugui"}
	t := template.New("test")
	t = template.Must(t.Parse(
		`hello {{ .UserName }} !
{{- /* this line is a comment */}}
{{range .Emails}}
an email {{ . }}
{{- end }}
{{ with .Friends }}
{{- range . }}
my friend name is {{.Fname}}
{{- end }}
{{- end }}
{{- "put" | printf "%s%s" "out" | printf "%q"}}
{{$a:=true}}
{{- $b:="abc"}}
{{- $c:="def"}}
{{- if $a }}
{{- $b = (print $b "/" $c)}}
{{- end}}
{{ print $b}}
`))
	p := Person{UserName: "longshuai",
		Emails:  []string{"a1@qq.com", "a2@gmail.com"},
		Friends: []*Friend{&f1, &f2}}
	t.Execute(os.Stdout, p)
}
