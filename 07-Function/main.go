package main

import (
	"text/template"
	"os"
	"log"
	"strings"
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
var myFunc = template.FuncMap{ //created a FuncMap to register functions.
	"uc" : strings.ToUpper, //"uc" is what the func will be called in the template and is a ToUpper func from package strings
	"inc"  : incSalaryCont, // "inc" is a func I declared this will get the salary and increament it with multipels of 5
	 "inp" : incSalaryPerm, // "inc" is a func I declared this will get the salary and increament it by 3000
	}

func incSalaryCont(sal int) int{
	newSal := sal * 10
	return newSal
}
func incSalaryPerm(sal int) int{
	newSal := sal + 3000
	return newSal
}
var tmpl *template.Template

func init(){
	tmpl = template.Must(template.New("").Funcs(myFunc).ParseFiles("tpl.gohtml"))
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

 log.Fatalln(tmpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", e))
}