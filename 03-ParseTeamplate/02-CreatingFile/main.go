package main

import (
	
	"text/template"
	"os"
	"log"
)
func main(){
	tpl,err := template.ParseFiles("tpl.gohtml")
	if err != nil{
		log.Fatalln(err)
	}
	file,err := os.Create("index.html")
    if err!= nil{
		log.Fatalln("Error in Creating File : ",err)
	}
	defer file.Close()
	tpl.Execute(file,nil)

}