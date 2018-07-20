package main

import (
	"github.com/factorapp/structure"
)

type HelloController struct {
	structure.BasicController
	Name string `source:"#name"`
	Output string `target:"#span"`
}

// type OtherThingController struct {
// 	structure.
// }

func main() {
	structure.RegisterController("hello", &HelloController{})
	structure.Run()
}
