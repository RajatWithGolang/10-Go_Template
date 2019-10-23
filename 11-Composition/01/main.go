package main

import (
	"text/template"
	"os"
	"log"
	//"time"
)

type Person struct{
	Name string
	Age int
}
type Occupation struct{
	Person
	Engineeer bool
}

var tmpl *template.Template


func init(){
	tmpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main(){
  o1:= Occupation{Person{
	  "Rajat",
	   29},
	  true,
  }
err := tmpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", o1)
	if err != nil {
		log.Fatalln(err)
	}

}