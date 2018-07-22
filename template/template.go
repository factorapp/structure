package template

import (
	"path/filepath"

	"github.com/gobuffalo/packr"
	"github.com/gobuffalo/plush"
)

type Renderer struct {
	componentName string
	box           packr.Box
}

func NewRenderer(componentsBox packr.Box, componentName string) Renderer {
	return Renderer{
		componentName: componentName,
		box:           componentsBox,
	}
}

func (p Renderer) Render(tplName string, data map[string]interface{}) (string, error) {
	tplStr, err := p.box.MustString(filepath.Join("components", p.componentName, "_"+tplName))
	if err != nil {
		return "", err
	}
	plushCtx := plush.NewContextWith(data)
	return plush.Render(tplStr, plushCtx)
}

func NewComponentsBox(componentsBaseDir string) (*packr.Box, error) {
	// use absolute path so that packr doesn't look in the GOPATH.
	// see https://github.com/gobuffalo/packr/blob/ee34b116572778801ca4a9f6355eda4577cabce8/box.go#L26
	abs, err := filepath.Abs(componentsBaseDir)
	if err != nil {
		return nil, err
	}
	box := packr.NewBox(abs)
	return &box, nil
}
