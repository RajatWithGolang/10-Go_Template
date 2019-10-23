package main

import (
	"log"
	"os"
	"text/template"
)

type Person struct {
	Name   string
	Age int
	Caption string 
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	P1 := Person{
		Name:   "Rajat",
		Age: 29,
		Caption:   `<script>alert("somthing wrong!");</script>`,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", P1)
	if err != nil {
		log.Fatalln(err)
	}
}