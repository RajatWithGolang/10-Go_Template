package main

import(
	"os"
	"log"
	"text/template"
)
type Contract struct{
	Ename string
	Salary int
}

type Permanent struct{
	Ename string
	Salary int
}

type Employee struct{
	Con []Contract
	Per  []Permanent
}
var tmpl *template.Template

func init(){
	tmpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
func main(){

 c1:= Contract{"Rajat",10000}
 c2:= Contract{"Rishabh",12000}
 c3:= Contract{"Aman",15000}

 p1:= Permanent{"Sajid",50000}
 p2:= Permanent{"Mudit",60000}  

 c:= []Contract{c1,c2,c3}

 p:= []Permanent{p1,p2}

 e := Employee{c,p}

 log.Fatalln(tmpl.Execute(os.Stdout,e))
 

}