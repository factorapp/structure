package main

import "github.com/factorapp/structure"

type HelloController struct {
	structure.BasicController
	Name string `source:"input#first-name"`
}

// type OtherThingController struct {
// 	structure.
// }

func main() {
	structure.RegisterController(&HelloController{})

	//structure.RegisterController(NewHTMLAJAXController("div#myhtml", "/mystuff..."))
}
