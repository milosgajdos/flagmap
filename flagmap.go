package flagmap

import (
	"flag"
	"fmt"
)

var options = make(map[string]string)

type MapValue struct {
	name string
	m    map[string]string
}

func Var(value flag.Value, usage string) {
	flag.Var(value, value.String(), usage)
}

func Parse() map[string]string {
	flag.Parse()
	return options
}

func Option(name string) *MapValue {
	return &MapValue{name, options}
}

func (m *MapValue) Set(value string) error {
	m.m[m.name] = value
	return nil
}

func (m MapValue) String() string {
	return fmt.Sprint(m.name)
}