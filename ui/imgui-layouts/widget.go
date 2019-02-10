package imgui_layouts

const imguiTitleBarHeight = 20
const imguiMenuBarHeight = 20
const imguiBorderSize = 2

type Panel struct {
	DisplayProperties struct {
		HasTitleBar bool
		HasMenuBar bool
	}
	Size [2]int
	Position [2]int

}

func (panel *Panel) PanelSize() (x,y int) {
	return panel.Size[0], panel.Size[1]
}

func (panel *Panel) InternalSize() (x,y int){
	x = panel.Size[0] - (2 * imguiBorderSize)
	y = panel.Size[1]- (2 * imguiBorderSize)
	if panel.DisplayProperties.HasTitleBar == true {
		y -= imguiTitleBarHeight
	}
	if panel.DisplayProperties.HasMenuBar == true {
		y -= imguiMenuBarHeight
	}

	return x,y
}

func (panel *Panel) SetPosition(x, y int) {

}

func (panel *Panel) Resize(x,y int) {
	if panel.Size[0] != x {
		panel.Size[0] = x
	}

	if panel.Size[1] != y {
		panel.Size[1] = y
	}
}