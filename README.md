flagmap
=======

Simplified wrapper package around Go flag package which returns a map of parsed flag values with flag names as keys and flag values as strings
You can supply a default value as a second argument go ```Options()``` function. See the example below.
```Defined()`` method of ```FlagMap``` type evaluates to false IF the specified parameter is empty string or does not exist in the FlagMap.

Example usage
==============


```go
package main

import (
	"fmt"
	"github.com/milosgajdos83/flagmap"
)

func init() {
        flagmap.Option("foo1", "", "This is a foo1 option")
        flagmap.Option("foo2", "foo2val", "This is a foo2 option")
}

func main() {
	options := flagmap.Parse()

	fmt.Printf("Options: %v\n", options)
        fmt.Printf("Just foo2: %v\n", options["foo2"])
}
```

Once built, you can test the above code like this:
```
$ ./yourbinary -foo2="test2"
Options: map[foo1: foo2:test2]
foo1 defined: false
```
