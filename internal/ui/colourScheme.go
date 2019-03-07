package ui

import "github.com/inkyblackness/imgui-go"

// ApplyImguiStyles sets the Imgui layout and theme.
func ApplyImguiStyles(themeId int) {
	style := imgui.CurrentStyle()

	commonLayout(&style)

	switch themeId {
	case 1:
		darkTheme(&style)
	default:
		lightTheme(&style)
	}
}

func commonLayout(style *imgui.Style) {
	style.ScaleAllSizes(DPIScale())
	fontAtlas := imgui.CurrentIO().Fonts()
	fontConfig := imgui.NewFontConfig()
	fontConfig.SetSize(DPIScale() * float32(DefaultFontSize))
	fontAtlas.AddFontFromFileTTF("./assets/fonts/FiraCode-Regular.ttf", DPIScale() * DefaultFontSize)
	fontAtlas.AddFontDefaultV(fontConfig)
	//fontAtlas.AddFontDefault()
}

func lightTheme(style *imgui.Style) {
	style.SetColor(imgui.StyleColorText, imgui.Vec4{0.31, 0.25, 0.24, 1.00})
	style.SetColor(imgui.StyleColorWindowBg, imgui.Vec4{0.94, 0.94, 0.94, 1.00})
	style.SetColor(imgui.StyleColorMenuBarBg, imgui.Vec4{0.74, 0.74, 0.94, 1.00})
	style.SetColor(imgui.StyleColorChildBg, imgui.Vec4{0.68, 0.68, 0.68, 0.00})
	style.SetColor(imgui.StyleColorBorder, imgui.Vec4{0.50, 0.50, 0.50, 0.60})
	style.SetColor(imgui.StyleColorBorderShadow, imgui.Vec4{0.00, 0.00, 0.00, 0.00})
	style.SetColor(imgui.StyleColorFrameBg, imgui.Vec4{0.62, 0.70, 0.72, 0.56})
	style.SetColor(imgui.StyleColorFrameBgHovered, imgui.Vec4{0.95, 0.33, 0.14, 0.47})
	style.SetColor(imgui.StyleColorFrameBgActive, imgui.Vec4{0.97, 0.31, 0.13, 0.81})
	style.SetColor(imgui.StyleColorTitleBg, imgui.Vec4{0.42, 0.75, 1.00, 0.53})
	style.SetColor(imgui.StyleColorTitleBgCollapsed, imgui.Vec4{0.40, 0.65, 0.80, 0.20})
	style.SetColor(imgui.StyleColorScrollbarBg, imgui.Vec4{0.40, 0.62, 0.80, 0.15})
	style.SetColor(imgui.StyleColorScrollbarGrab, imgui.Vec4{0.39, 0.64, 0.80, 0.30})
	style.SetColor(imgui.StyleColorScrollbarGrabHovered, imgui.Vec4{0.28, 0.67, 0.80, 0.59})
	style.SetColor(imgui.StyleColorScrollbarGrabActive, imgui.Vec4{0.25, 0.48, 0.53, 0.67})
	style.SetColor(imgui.StyleColorWindowBg, imgui.Vec4{0.89, 0.98, 1.00, 0.99})
	style.SetColor(imgui.StyleColorCheckMark, imgui.Vec4{0.48, 0.47, 0.47, 0.71})
	style.SetColor(imgui.StyleColorSliderGrabActive, imgui.Vec4{0.31, 0.47, 0.99, 1.00})
	style.SetColor(imgui.StyleColorButton, imgui.Vec4{1.00, 0.79, 0.18, 0.78})
	style.SetColor(imgui.StyleColorButtonHovered, imgui.Vec4{0.42, 0.82, 1.00, 0.81})
	style.SetColor(imgui.StyleColorButtonActive, imgui.Vec4{0.72, 1.00, 1.00, 0.86})
	style.SetColor(imgui.StyleColorHeader, imgui.Vec4{0.65, 0.78, 0.84, 0.80})
	style.SetColor(imgui.StyleColorHeaderHovered, imgui.Vec4{0.75, 0.88, 0.94, 0.80})
	style.SetColor(imgui.StyleColorHeaderActive, imgui.Vec4{0.55, 0.68, 0.74, 0.80}) //imgui.Vec4{0.46, 0.84, 0.90, 1.00}})
	style.SetColor(imgui.StyleColorResizeGrip, imgui.Vec4{0.60, 0.60, 0.80, 0.30})
	style.SetColor(imgui.StyleColorResizeGripHovered, imgui.Vec4{1.00, 1.00, 1.00, 0.60})
	style.SetColor(imgui.StyleColorResizeGripActive, imgui.Vec4{1.00, 1.00, 1.00, 0.90})
	style.SetColor(imgui.StyleColorTextSelectedBg, imgui.Vec4{1.00, 0.99, 0.54, 0.43})
}

func darkTheme(style *imgui.Style) {
	style.SetColor(imgui.StyleColorText, imgui.Vec4{1.000, 1.000, 1.000, 1.000})
	style.SetColor(imgui.StyleColorTextDisabled, imgui.Vec4{0.500, 0.500, 0.500, 1.000})
	style.SetColor(imgui.StyleColorWindowBg, imgui.Vec4{0.180, 0.180, 0.180, 1.000})
	style.SetColor(imgui.StyleColorChildBg, imgui.Vec4{0.280, 0.280, 0.280, 0.000})
	style.SetColor(imgui.StyleColorPopupBg, imgui.Vec4{0.313, 0.313, 0.313, 1.000})
	style.SetColor(imgui.StyleColorBorder, imgui.Vec4{0.266, 0.266, 0.266, 1.000})
	style.SetColor(imgui.StyleColorBorderShadow, imgui.Vec4{0.000, 0.000, 0.000, 0.000})
	style.SetColor(imgui.StyleColorFrameBg, imgui.Vec4{0.160, 0.160, 0.160, 1.000})
	style.SetColor(imgui.StyleColorFrameBgHovered, imgui.Vec4{0.200, 0.200, 0.200, 1.000})
	style.SetColor(imgui.StyleColorFrameBgActive, imgui.Vec4{0.280, 0.280, 0.280, 1.000})
	style.SetColor(imgui.StyleColorTitleBg, imgui.Vec4{0.148, 0.148, 0.148, 1.000})
	style.SetColor(imgui.StyleColorTitleBgActive, imgui.Vec4{0.148, 0.148, 0.148, 1.000})
	style.SetColor(imgui.StyleColorTitleBgCollapsed, imgui.Vec4{0.148, 0.148, 0.148, 1.000})
	style.SetColor(imgui.StyleColorMenuBarBg, imgui.Vec4{0.195, 0.195, 0.195, 1.000})
	style.SetColor(imgui.StyleColorScrollbarBg, imgui.Vec4{0.160, 0.160, 0.160, 1.000})
	style.SetColor(imgui.StyleColorScrollbarGrab, imgui.Vec4{0.277, 0.277, 0.277, 1.000})
	style.SetColor(imgui.StyleColorScrollbarGrabHovered, imgui.Vec4{0.300, 0.300, 0.300, 1.000})
	style.SetColor(imgui.StyleColorScrollbarGrabActive, imgui.Vec4{1.000, 0.391, 0.000, 1.000})
	style.SetColor(imgui.StyleColorCheckMark, imgui.Vec4{1.000, 1.000, 1.000, 1.000})
	style.SetColor(imgui.StyleColorSliderGrab, imgui.Vec4{0.391, 0.391, 0.391, 1.000})
	style.SetColor(imgui.StyleColorSliderGrabActive, imgui.Vec4{1.000, 0.391, 0.000, 1.000})
	style.SetColor(imgui.StyleColorButton, imgui.Vec4{1.000, 1.000, 1.000, 0.000})
	style.SetColor(imgui.StyleColorButtonHovered, imgui.Vec4{1.000, 1.000, 1.000, 0.156})
	style.SetColor(imgui.StyleColorButtonActive, imgui.Vec4{1.000, 1.000, 1.000, 0.391})
	style.SetColor(imgui.StyleColorHeader, imgui.Vec4{0.313, 0.313, 0.313, 1.000})
	style.SetColor(imgui.StyleColorHeaderHovered, imgui.Vec4{0.469, 0.469, 0.469, 1.000})
	style.SetColor(imgui.StyleColorHeaderActive, imgui.Vec4{0.469, 0.469, 0.469, 1.000})
	style.SetColor(imgui.StyleColorSeparator, imgui.Vec4{0.266, 0.266, 0.266, 1.000})
	style.SetColor(imgui.StyleColorSeparatorHovered, imgui.Vec4{0.391, 0.391, 0.391, 1.000})
	style.SetColor(imgui.StyleColorSeparatorActive, imgui.Vec4{1.000, 0.391, 0.000, 1.000})
	style.SetColor(imgui.StyleColorResizeGrip, imgui.Vec4{1.000, 1.000, 1.000, 0.250})
	style.SetColor(imgui.StyleColorResizeGripHovered, imgui.Vec4{1.000, 1.000, 1.000, 0.670})
	style.SetColor(imgui.StyleColorResizeGripActive, imgui.Vec4{1.000, 0.391, 0.000, 1.000})
	//style.SetColor(imgui.StyleColorTab, imgui.Vec4{0.098, 0.098, 0.098, 1.000})
	//style.SetColor(imgui.StyleColorTabHovered, imgui.Vec4{0.352, 0.352, 0.352, 1.000})
	//style.SetColor(imgui.StyleColorTabActive, imgui.Vec4{0.195, 0.195, 0.195, 1.000})
	//style.SetColor(imgui.StyleColorTabUnfocused, imgui.Vec4{0.098, 0.098, 0.098, 1.000})
	//style.SetColor(imgui.StyleColorTabUnfocusedActive, imgui.Vec4{0.195, 0.195, 0.195, 1.000})
	//style.SetColor(imgui.StyleColorDockingPreview, imgui.Vec4{1.000, 0.391, 0.000, 0.781})
	//style.SetColor(imgui.StyleColorDockingEmptyBg, imgui.Vec4{0.180, 0.180, 0.180, 1.000})
	style.SetColor(imgui.StyleColorPlotLines, imgui.Vec4{0.469, 0.469, 0.469, 1.000})
	style.SetColor(imgui.StyleColorPlotLinesHovered, imgui.Vec4{1.000, 0.391, 0.000, 1.000})
	style.SetColor(imgui.StyleColorPlotHistogram, imgui.Vec4{0.586, 0.586, 0.586, 1.000})
	style.SetColor(imgui.StyleColorPlotHistogramHovered, imgui.Vec4{1.000, 0.391, 0.000, 1.000})
	style.SetColor(imgui.StyleColorTextSelectedBg, imgui.Vec4{1.000, 1.000, 1.000, 0.156})
	style.SetColor(imgui.StyleColorDragDropTarget, imgui.Vec4{1.000, 0.391, 0.000, 1.000})
	style.SetColor(imgui.StyleColorNavHighlight, imgui.Vec4{1.000, 0.391, 0.000, 1.000})
	style.SetColor(imgui.StyleColorNavWindowingHighlight, imgui.Vec4{1.000, 0.391, 0.000, 1.000})
	//style.SetColor(imgui.StyleColorNavWindowingDimBg, imgui.Vec4{0.000, 0.000, 0.000, 0.586})
	//style.SetColor(imgui.StyleColorModalWindowDimBg, imgui.Vec4{0.000, 0.000, 0.000, 0.586})
	imgui.PushStyleVarFloat(imgui.StyleVarChildRounding, 4)
	imgui.PushStyleVarFloat(imgui.StyleVarFrameBorderSize, 1.0)
	imgui.PushStyleVarFloat(imgui.StyleVarFrameRounding, 2.0)
	imgui.PushStyleVarFloat(imgui.StyleVarGrabMinSize, 7.0)
	imgui.PushStyleVarFloat(imgui.StyleVarPopupRounding, 2.0)
	imgui.PushStyleVarFloat(imgui.StyleVarScrollbarRounding, 12.0)
	imgui.PushStyleVarFloat(imgui.StyleVarScrollbarSize, 13.0)
	//imgui.PushStyleVarFloat(imgui.StyleVarTabBorderSize, 1.0 )
	//imgui.PushStyleVarFloat(imgui.StyleVarTabRounding, 0.0 )
	imgui.PushStyleVarFloat(imgui.StyleVarWindowRounding, 4.0)
}
