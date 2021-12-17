package variable

import (
	"fmt"
	varconfig "ttslight/config/variables"
)

type Variable struct {
	Name string
	Type int
	Path string
}

func New(name string, vartype int, path string) Variable {
	e := Variable{name, vartype, path}
	return e
}

func NewFromConf(varcon varconfig.Variable) Variable {
	e := Variable{varcon.Name, varcon.Type, varcon.Path}
	return e
}

func (v Variable) ToString() string {
	var result string = fmt.Sprintf("[Variable: name: %v, type: %v, path: %v]", v.Name, v.Type, v.Path)
	return result
}
