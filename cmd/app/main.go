//go:generate bash -c "cp $DOLLAR(go env GOROOT)/misc/wasm/wasm_exec.js ./app/wasm_exec.js"
package main

import (
	"github.com/factorapp/structure"
)

type SlideshowController struct {
	structure.BasicController
	Index int
}

func (s *SlideshowController) Next() {
	s.Index++
}

func (s *SlideshowController) Previous() {
	s.Index--
}
// type OtherThingController struct {
// 	structure.
// }

func main() {
	structure.RegisterController("slideshow", &SlideshowController{})
	structure.Run()
}
