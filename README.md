flagmap
=======

Simplified wrapper package around Go flag package which returns a map of parsed flag values with flag names as keys and flag values as interface slices

This package allows you to "aggregate" passed flags as oppose to "overriding" them which gives you more flexibility as you can see in the example code snippet below.

Example usage
==============


```go
package main

import (
	"fmt"
	"github.com/milosgajdos83/flagmap"
)

func init() {
        flagmap.Option("foo1", "This is a foo1 option")
        flagmap.Option("foo2", "This is a foo2 option")
}

func main() {
	options := flagmap.Parse()

	fmt.Printf("Options: %v\n", options)
        fmt.Printf("Just foo2: %v\n", options["foo2"])
}
```

Once built, you can test the above code like this:
```
$ ./yourbinary -foo1="test1" -foo2="test2" -foo2=20
All options: map[foo1:[test1] foo2:[test2 20]]
Just foo2: [test2 20]
```
