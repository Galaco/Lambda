package hierarchy

import (
	"fmt"
	"github.com/inkyblackness/imgui-go"
	"log"
	"strings"
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
	currentFilterMode int
}

func (mod *allEntities) Render() {
	imgui.BeginChild("HierarchyMainScrolling")
	for _, row := range mod.filtered {
		row.Render()
	}
	imgui.EndChild()
}

func (mod *allEntities) Filter(filterMode int) {
	switch filterMode {
	case entityFilterPropOnly:
		mod.filtered = make([]*Item, 0)
		for idx,n := range mod.nodes {
			if strings.HasPrefix(n.Classname, "prop_") == true {
				mod.filtered = append(mod.filtered, &mod.nodes[idx])
			}
		}
		mod.currentFilterMode = entityFilterPropOnly
	default:
		mod.filtered = make([]*Item, len(mod.nodes))
		for idx := range mod.nodes {
			mod.filtered[idx] = &mod.nodes[idx]
		}
		mod.currentFilterMode = entityFilterNone
	}
	log.Println(fmt.Sprintf("Filter mode: %d", filterMode))
}

func (mod *allEntities) AddEntity(id int, classname string, targetname string) {
	mod.nodes = append(mod.nodes, NewItem(id, classname, targetname))
	mod.Filter(mod.currentFilterMode)
}


