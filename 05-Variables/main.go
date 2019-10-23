package main

import (
	
	"text/template"
	"os"
	"log"
)
var tpl *template.Template
func init(){
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
func main(){
	
	log.Fatalln(tpl.ExecuteTemplate(os.Stdout,"tpl.gohtml","Rajat Singh Rawat"))

}