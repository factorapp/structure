package core

// Template represents a parsed template
type Template interface {
	// Render renders a template to id with the given data
	RenderTo(id string, data interface{}) error
}
