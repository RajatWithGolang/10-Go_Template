package main

import (
	"os"
	"text/template"
	"log"

)
var tmpl *template.Template

func init(){
	tmpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}
func main(){
    log.Fatalln(tmpl.ExecuteTemplate(os.Stdout,"tpl.gohtml","Rajat"))
}