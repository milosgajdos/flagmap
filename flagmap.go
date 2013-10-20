package flagmap

import (
	"flag"
	"fmt"
)

type FlagMap map[string][]interface{}

var options = make(FlagMap)

type MapValue struct {
	name string
	m   FlagMap
}

func Var(value flag.Value, usage string) {
	flag.Var(value, value.String(), usage)
}

func Parse() map[string][]interface{} {
	flag.Parse()
	return options
}

func Option(name string) *MapValue {
	return &MapValue{name, options}
}

func (m *MapValue) Set(value string) error {
	m.m[m.name] = append(m.m[m.name], value)
	return nil
}

func (m MapValue) String() string {
	return fmt.Sprint(m.name)
}

func (FlagMap) Defined(name string) bool {
        if _, ok := options[name]; !ok {
                return false
        }
        return true
}
