package list

import (
	"strings"
)

// List contains a slice of rows that can be rendered
// using imgui
type List struct {
	nodes    []row
	selected int
}

// Filter applies a filtermode on the stored rows and returns
// a FilteredList that contains all rows that pass the filter.
func (mod *List) Filter(filterMode int) *FilteredList {
	filtered := []*row{}
	switch filterMode {
	case EntityFilterPropOnly:
		for idx, n := range mod.nodes {
			if strings.Contains(n.label, " prop_") == true {
				filtered = append(filtered, &mod.nodes[idx])
			}
		}
	default:
		filtered = make([]*row, len(mod.nodes))
		for idx := range mod.nodes {
			filtered[idx] = &mod.nodes[idx]
		}
	}

	return NewFilteredList(filtered)
}

// AddRow adds a new row to the end of the list.
func (mod *List) AddRow(id int, label string, onClick func(id int)) {
	mod.nodes = append(mod.nodes, newRow(id, label, func(id int) {
		mod.selected = id
		onClick(id)
	}))
}
