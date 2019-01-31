


##### Usage

```go
package main

import (
	"github.com/galaco/Lambda/lib/imgui-layouts/keyvalues"
	"log"
)

func main() {
	view := keyvalues.NewKeyValues()
	for i := 0; i < 4; i++ {
		view.AddKeyValue(keyvalues.NewKeyValue("foo", "value", func(k, v string) {
			//on change
		}))
	}
	
	//... do things
	
	// render
	view.Render()
}
```