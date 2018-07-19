package structure

type HelloController struct {
	Controller
	Name string `source:"input#first-name"`
}
type Hello2Controller struct {
	Controller
	Name string `target:"input#first-name"`
}
