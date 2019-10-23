package main


import (
	"fmt"
	"os"
	"io"
	"strings"
)

func main(){
	name := "Rajat Singh Rawat"

	str := fmt.Sprintf(`
	 <!DOCTYPE html>
	<html lang="en">
<head>
    <meta charset="UTF-8">
    
    <title>Document</title>
</head>
<body>
    <h1>` + name +` </h1>
</body>
</html>
	`)
	file,_  := os.Create("index.html")
	defer file.Close()
	io.Copy(file,strings.NewReader(str))

	fmt.Println(str)
}