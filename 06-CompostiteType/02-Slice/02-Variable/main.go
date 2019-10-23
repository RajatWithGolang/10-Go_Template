package main

import (
	"os"
	"text/template"
	"log"
)
var tmpl *template.Template

func init(){
	tmpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
func main(){
	xi := []int{10,20,30,40,50}
    log.Fatalln(tmpl.ExecuteTemplate(os.Stdout,"tpl.gohtml",xi))
}