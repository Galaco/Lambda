##### Directory


```go
package main

import (
	"log"
	"github.com/galaco/imgui-layouts/directory"
)

func main() {
	thumbs := make([]directory.Thumbnail, 100)
	for i := 0; i < 100; i++ {
		thumbs[i] = *directory.NewThumbnail("foo", false, func() {
			log.Println(i)
		})
	}
	
	view := directory.NewDirectory(thumbs)
	
	//... do things
	
	// render
	view.Render()
}





```