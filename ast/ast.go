package ast

import (
	"bytes"
	"fmt"
)

type JSONData interface {
	String() string
	jsonData()
}

type Object struct {
	Value map[string]JSONData
}

func (o *Object) jsonData() {}

func (o *Object) String() string {
	var out bytes.Buffer
	out.WriteString("{")
	for k, v := range o.Value {
		out.WriteString(fmt.Sprintf("%v:%v,", k, v.String()))
	}
	out.WriteString("}")
	return out.String()
}

type Array struct {
	Value []JSONData
}

type Number struct {
	Value float64
}

func (n *Number) jsonData() {}

func (n *Number) String() string {
	return fmt.Sprintf("%v", n.Value)
}

type String struct {
	Value string
}

func (s *String) jsonData() {}

func (s *String) String() string {
	return s.Value
}

type Boolean struct {
	Value bool
}

func (b *Boolean) jsonData() {}

func (b *Boolean) String() string {
	return fmt.Sprintf("%v", b.Value)
}

type Null struct {
}

func (n *Null) jsonData() {}

func (n *Null) String() string {
	return "null"
}
