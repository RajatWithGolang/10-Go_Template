package main


import "fmt"

func main(){
	name := "Rajat Singh Rawat"

	str := `
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
	`
	fmt.Println(str)
}