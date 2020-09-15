package vgui

type MasterPanel struct {
	Panel
}

func (panel *MasterPanel) Resize(width, height float64) {
	for _, child := range panel.children {
		child.Resize(width, height)
	}
}
