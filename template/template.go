package template

import (
	"bytes"
	"fmt"
	"html/template"
	"syscall/js"

	dom "github.com/gowasm/livedom"
)

type Renderer struct {
	elt *dom.Element
}

func NewRenderer(elt *dom.Element) Renderer {
	return Renderer{
		elt: elt,
	}
}

func (p Renderer) Render(tplName string, data map[string]interface{}) (string, error) {
	tplElt := p.elt.Call("querySelector", "#"+tplName)
	//tplElt := p.elt.QuerySelector("#" + tplName)
	if tplElt == js.Null() {
		return "", fmt.Errorf("no template %s found", tplName)
	}
	tplStrVal := tplElt.Get("innerHTML") //InnerHTML()
	if tplStrVal == js.Null() {
		return "", fmt.Errorf("no inner HTML in %s", tplName)
	}
	tplStr := tplStrVal.String()
	tpl, err := template.New(tplStr).Parse(tplStr)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err := tpl.Execute(buf, data); err != nil {
		return "", err
	}

	return string(buf.Bytes()), nil
}
