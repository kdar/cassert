cassert
=======

GoConvey test assertions without the BDD framework

## Examples

```
package main

import (
	. "github.com/kdar/cassert"
	"testing"
)

func TestSomething(t *testing.T) {
	So(t, 5, ShouldEqual, 5)
	So(t, true, ShouldBeTrue)
}

```

```
package main

import (
	"github.com/kdar/cassert"
	"testing"
)

func TestSomething(t *testing.T) {
	So := cassert.New(t)
	So(5, cassert.ShouldEqual, 5)
	So(true, cassert.ShouldBeTrue)
}

```