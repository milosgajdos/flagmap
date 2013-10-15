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
	flagmap.Var(flagmap.Option("foo"), "foo", "This is a foo option")
}

func main() {
	options := flagmap.Parse()

	fmt.Printf("Options: %v", options["foo"])	
}
```

Once built, you can test the above code like this:
```
$ ./compiledbinary -foo="test"
Options: map[foo:test]
```
