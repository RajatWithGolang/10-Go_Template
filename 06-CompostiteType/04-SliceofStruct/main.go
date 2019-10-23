package main

import(
	"os"
	"log"
	"text/template"
)
type Person struct{
	Name string
	Age int
}
var tmpl *template.Template

func init(){
	tmpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
func main(){

 p1:= Person{"Rajat",29}
 p2:= Person{"Rishabh",22}
 p3:= Person{"Richa",21}
 p4:= Person{"Simi",29}
 p5:= Person{"Riyansh",1}  
 p := []Person{p1,p2,p3,p4,p5}
 log.Fatalln(tmpl.Execute(os.Stdout,p))
 

}