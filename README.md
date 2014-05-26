go-lil-t
========

Little testing-helpers for go

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
