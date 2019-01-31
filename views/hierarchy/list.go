package hierarchy

import "github.com/galaco/Lambda/views/hierarchy/list"

// entityList contains the master list of entities and provides
// a gateway for managing filtering the masterlist into subsets.
type entityList struct {
	entityList        list.List
	filteredList      *list.FilteredList
	currentFilterMode int
}

// addEntity adds a new entity to the list
func (l *entityList) addEntity(id int, classname string, targetname string, onClick func(id int)) {
	l.entityList.AddRow(id, classname, targetname, onClick)
	l.filter()
}

// getFiltered returns a renderable list of entities filtered by the
// current filter options
func (l *entityList) getFiltered() *list.FilteredList {
	if l.filteredList == nil {
		l.filter()
	}
	return l.filteredList
}

// setFilterMode sets th current filter
func (l *entityList) setFilterMode(mode int) {
	if l.currentFilterMode != mode {
		l.currentFilterMode = mode
		l.filter()
	}
}

// filter rebuilds the list of filtered entities
func (l *entityList) filter() {
	l.filteredList = l.entityList.Filter(l.currentFilterMode)
}
