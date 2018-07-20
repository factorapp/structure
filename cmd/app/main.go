//go:generate bash -c "cp $DOLLAR(go env GOROOT)/misc/wasm/wasm_exec.js ./app/wasm_exec.js"
package main

import (
	"github.com/factorapp/structure"
)

type HelloController struct {
	structure.BasicController
	Name   string `source:"Name"`
	Output string `target:"Output"`
}

// type OtherThingController struct {
// 	structure.
// }

func main() {
	structure.RegisterController("hello", &HelloController{})
	structure.Run()
}
