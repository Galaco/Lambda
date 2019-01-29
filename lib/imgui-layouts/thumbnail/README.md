##### Columns


```go
package main

import (
	"log"
	"github.com/galaco/imgui-layouts/thumbnail"
)

func main() {
	thumbs := make([]thumbnail.Thumbnail, 100)
	for i := 0; i < 100; i++ {
		thumbs[i] = *thumbnail.NewThumbnail("foo", false, func() {
			log.Println(i)
		})
	}
	
	view := thumbnail.NewDirectory(thumbs)
	
	//...
	
	// render
	view.Render()
}





```