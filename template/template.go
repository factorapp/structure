package template

import (
	"bytes"
	"fmt"
	"html/template"

	dom "github.com/gowasm/go-js-dom"
)

type Renderer struct {
	elt dom.Element
}

func NewRenderer(elt dom.Element) Renderer {
	return Renderer{
		elt: elt,
	}
}

func (p Renderer) Render(tplName string, data map[string]interface{}) (string, error) {
	tplElt := p.elt.QuerySelector("#" + tplName)
	if tplElt == nil {
		return "", fmt.Errorf("no template %s found", tplName)
	}
	tplStr := tplElt.InnerHTML()
	tpl, err := template.New(tplName).Parse(tplStr)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err := tpl.Execute(buf, data); err != nil {
		return "", err
	}

	return string(buf.Bytes()), nil
}
