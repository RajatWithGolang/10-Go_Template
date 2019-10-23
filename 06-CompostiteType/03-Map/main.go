package main

import(
	"os"
	"log"
	"text/template"
)

var tmpl *template.Template

func init(){
	tmpl = template.Must(template.ParseFiles("tpl.gohtml"))
	}
type Person struct{
	Name string
	Age int
}
func main(){
	myMap := map[int]string{
		1 : "Rajat",
		2 : "Rishabh",
		3:  "Richa",
		4 : "Simi",
	}
        log.Fatalln(tmpl.ExecuteTemplate(os.Stdout,"tpl.gohtml",myMap))
}