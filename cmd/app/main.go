//go:generate bash -c "cp $DOLLAR(go env GOROOT)/misc/wasm/wasm_exec.js ./app/wasm_exec.js"
package main

import (
	"fmt"
	"github.com/factorapp/structure"
)

type SlideshowController struct {
	structure.BasicController
	Index int
}

func (s *SlideshowController) currentSlide() int {
	return s.Index+1
}
func (s *SlideshowController) Next() {
	tlen := len(s.Targets()["slide"])
	if s.Index >= tlen {
		s.Index= 0
	}
	s.Index++
	fmt.Println("INDEX:", s.Index)
}

func (s *SlideshowController) Previous() {

	s.Index--

	if r.Index <= 0{
		s.Index = len(s.Targets()["slide"]) 
		fmt.Println("INDEX:", s.Index)
		return
	}
	fmt.Println("INDEX:", s.Index)
}
// type OtherThingController struct {
// 	structure.
// }

func main() {
	structure.RegisterController("slideshow", &SlideshowController{})
	structure.Run()
}
