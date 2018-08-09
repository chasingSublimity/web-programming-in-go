package main

import "fmt"

func main() {
	name := "Blake Sager"

	tpl := `
	<!DOCTYPE html>
	<html>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
	`
	fmt.Println(tpl)
}
