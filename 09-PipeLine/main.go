package main

import (
	"text/template"
	"os"
	"log"
	"time"
)



var tmpl *template.Template

var myFunc = template.FuncMap{
	"fdateMDY": monthDayYear,
}

func monthDayYear(t time.Time) string{
  return t.Format("01-02-2006")
}
func init(){
	tmpl = template.Must(template.New("").Funcs(myFunc).ParseFiles("tpl.gohtml"))
}

func main(){

err := tmpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}

}