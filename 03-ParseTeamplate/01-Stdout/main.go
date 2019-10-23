package main

import (
	"fmt"
	"text/template"
	"os"
	"log"
)
func main(){
	tpl,err := template.ParseFiles("tpl.gohtml")
      if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(os.Stdout,nil)
	fmt.Println("Rajat")
	if err != nil {
		log.Fatalln(err)
	}
}