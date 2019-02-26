package list

import (
	"github.com/inkyblackness/imgui-go"
)

// FilteredList contains references to rows in the master
// list. It represents a subset of all rows in another list.
type FilteredList struct {
	rows []*row
	selected int
}

// Render renders all the rows in this list
func (filteredList *FilteredList) Render() {
	imgui.BeginChild("HierarchyMainScrolling")
	for _, row := range filteredList.rows {
		if row.Id != filteredList.selected {
			row.selected = false
		}
		if row.render() {
			filteredList.selected = row.Id
		}
	}
	imgui.EndChild()
}

// NewFilteredList returns a new FilteredList populated from the
// passed row slice
func NewFilteredList(rows []*row) *FilteredList {
	return &FilteredList{
		rows: rows,
	}
}
