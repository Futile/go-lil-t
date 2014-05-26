go-lil-t
========

I wrote these two little helpers, `If` and `IfNot`, to replace testing code such as:

```go
    if result == nil {
        t.Errorf("error")
    }
```

with the much shorter:

```go
    If(result == nil).Errorf("error")
```

which is much more readable. Also it helps to remove clutter from tests.

Example:
---------

```go
package example

import (
	"testing"

	"github.com/futile/go-lil-t"
)


func TestMyStuff(t *testing.T) {
    If, IfNot := lilt.New(t)

    If(false).Errorf("critical bug")

    IfNot(true).Errorf("another critical bug")
}
```
