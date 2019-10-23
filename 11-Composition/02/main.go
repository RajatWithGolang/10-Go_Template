

package main

import (
	"text/template"
	"os"
	"log"
	//"time"
)

type course struct {
	Number, Name, Units string
}
type semester struct {
	Term    string
	Courses []course
}
type year struct {
	Fall, Spring, Summer semester
}

var tmpl *template.Template

func init(){
	tmpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main(){
 y := year{
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				course{"01","Course1","3",},
				course{"02","Course2","2",},
				course{"03","Course3","1",},
				
			},
		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				course{"04","Course4","6",},
				course{"05","Course5","5",},
				course{"06","Course6","4",},
			},
		},
			Summer: semester{
			Term: "Spring",
			Courses: []course{
				course{"07","Course7","4",},
				course{"08","Course8","5",},
				course{"09","Course9","6",},
			},
		},
	}
err := tmpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", y)
	if err != nil {
		log.Fatalln(err)
	}

}