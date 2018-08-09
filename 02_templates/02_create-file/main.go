package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "Blake Sager"

	str := fmt.Sprintf(`<!DOCTYPE html>
<html>
	<body>
		<h1>` + name + `</h1>
	</body>
</html>
	`)

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error createing file", err)
	}
	defer nf.Close()

	io.Copy(nf, strings.NewReader(str))
}
