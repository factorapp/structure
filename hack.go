package structure

type HelloController struct {
	DOMObjectController
	Name string `source:"input#first-name"`
}
type Hello2Controller struct {
	DOMObjectController
	Name string `target:"input#first-name"`
}
