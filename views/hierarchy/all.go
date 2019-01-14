package hierarchy

import (
	"fmt"
	"github.com/inkyblackness/imgui-go"
	"log"
)

const (
	entityFilterNone = 0
	entityFilterPointOnly = 1
	entityFilterBrushOnly = 2
	entityFilterPropOnly = 3
)

type allEntities struct {
	nodes []Item
	filtered []*Item
}

func (mod *allEntities) Render() {
	imgui.BeginChild("HierarchyMainScrolling")
	for _, row := range mod.nodes {
		row.Render()
	}
	imgui.EndChild()
}

func (mod *allEntities) Filter(filterMode int) {
	log.Println(fmt.Sprintf("Filter mode: %d", filterMode))
}

func (mod *allEntities) AddEntity(id int, classname string, targetname string) {
	mod.nodes = append(mod.nodes, NewItem(id, classname, targetname))
}


