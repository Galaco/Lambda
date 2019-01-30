


##### Usage

```go
package main

import (
	"github.com/galaco/Lambda/lib/imgui-layouts/columns"
	"log"
)

func main() {
	numColumns := 4
	view := columns.NewColumns(numColumns)
	
	for i := 0; i < 4; i++ {
		err := view.SetColumnContents(i, func() {
		    // render contents here
		}, nil)
		if err != nil {
		    panic(err)
		}
	}
	
	//... do things
	
	// render
	view.Render(1024, 768)
}
```