package core

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

type structFieldRef struct {
	name string
	get  func() interface{}
	set  func(interface{}) error
}

func newStructFieldRef(
	fieldName string,
	get func() interface{},
	set func(interface{}) error,
) Ref {
	return &structFieldRef{name: fieldName, get: get, set: set}
}

func (s structFieldRef) Name() string       { return s.name }
func (s structFieldRef) Value() interface{} { return s.get() }
func (s structFieldRef) Set(iface interface{}) error {
	return s.set(iface)
}
