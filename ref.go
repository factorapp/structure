package structure

type Ref interface {
	Name() string
	Set(interface{}) error
	Value() interface{}
}

type StringRef struct {
	str string
}

func NewStringRef(s string) *StringRef {
	return &StringRef{str: s}
}

func (s StringRef) Name() string { return s.str }
func (s *StringRef) Set(iface interface{}) error {
	// TODO: check type assert
	s.str = iface.(string)
	return nil
}

func (s StringRef) Value() interface{} {
	return s.str
}
