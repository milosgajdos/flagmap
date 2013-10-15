flagmap
=======

Wrapper around flag package which returns a map of parsed flag values with flag names as keys

Example usage
==============

package main

```go
import (
	"fmt"
	"github.com/milosgajdos83/flagmap"
)

func init() {
        flagmap.Var(flagmap.Option("foo1"), "This is a foo1 option")
        flagmap.Var(flagmap.Option("foo2"), "This is a foo2 option")
}

func main() {
	options := flagmap.Parse()

	fmt.Printf("Options: %v\n", options)
}
```

Once built, you can test the above code like this:
```
$ ./flagmap -foo1="test1" -foo2="test2"
Options: map[foo1:test1 foo2:test2]
```
