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
	 p1 :=Person{"Rajat",29}
        log.Fatalln(tmpl.ExecuteTemplate(os.Stdout,"tpl.gohtml",p1))
}