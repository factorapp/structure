package main

import "github.com/factorapp/structure"

type HelloController struct {
	DOMObjectController
	Name string `source:"input#first-name"`
}

func main() {
	structure.RegisterController(&HelloController{})
}
