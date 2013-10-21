package flagmap

import (
	"flag"
	"fmt"
)

type FlagMap map[string]string

var options = make(FlagMap)

type MapValue struct {
	name string
	m   FlagMap
}

func Parse() FlagMap {
	flag.Parse()
	return options
}

func Option(name string, defaultVal string, usage string) {
        mapVal := &MapValue{name, options}
        mapVal.m[name] = defaultVal
        flag.Var(mapVal, mapVal.String(), usage)
}

func (m *MapValue) Set(value string) error {
	m.m[m.name] = value
	return nil
}

func (m MapValue) String() string {
	return fmt.Sprint(m.name)
}

func (f FlagMap) Defined(name string) bool {
        if _, ok := f[name]; !ok {
                return false
        }

        if f[name] == "" {
                return false
        }
        return true
}
