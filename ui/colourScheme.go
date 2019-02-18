package ui

import "github.com/inkyblackness/imgui-go"

func ApplyImguiStyles(themeId int) {
	fontAtlas := imgui.CurrentIO().Fonts()
	fontAtlas.AddFontFromFileTTF("./FiraCode-Regular.ttf", 16)
	style := imgui.CurrentStyle()

	commonLayout()

	switch themeId {
	case 1:
		darkTheme(&style)
	default:
		lightTheme(&style)
	}
}

func commonLayout() {
	imgui.PushStyleVarVec2(imgui.StyleVarWindowPadding, imgui.Vec2{15, 15})
	imgui.PushStyleVarFloat(imgui.StyleVarWindowRounding, 0)
	imgui.PushStyleVarVec2(imgui.StyleVarFramePadding, imgui.Vec2{5, 5})
	imgui.PushStyleVarFloat(imgui.StyleVarFrameRounding, 0)
	imgui.PushStyleVarVec2(imgui.StyleVarItemSpacing, imgui.Vec2{12, 8})
	imgui.PushStyleVarVec2(imgui.StyleVarItemInnerSpacing, imgui.Vec2{8, 6})
	imgui.PushStyleVarFloat(imgui.StyleVarIndentSpacing, 25)
	imgui.PushStyleVarFloat(imgui.StyleVarScrollbarSize, 16)
	imgui.PushStyleVarFloat(imgui.StyleVarScrollbarRounding, 0)
	imgui.PushStyleVarFloat(imgui.StyleVarGrabMinSize, 5)
	imgui.PushStyleVarFloat(imgui.StyleVarGrabRounding, 0)
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
	//style.SetColor(imgui.StyleColorCloseButton, imgui.Vec4{0.41, 0.75, 0.98, 0.50}})
	//style.SetColor(imgui.StyleColorCloseButtonHovered, imgui.Vec4{1.00, 0.47, 0.41, 0.60}})
	//style.SetColor(imgui.StyleColorCloseButtonActive, imgui.Vec4{1.00, 0.16, 0.00, 1.00}})
	style.SetColor(imgui.StyleColorTextSelectedBg, imgui.Vec4{1.00, 0.99, 0.54, 0.43})
	//style.SetColor(imgui.StyleColorTooltipBg, imgui.Vec4{0.82, 0.92, 1.00, 0.90}})
	//style.Alpha = 1.0
	//style.WindowFillAlphaDefault = 1.0
	//style.FrameRounding = 4
	//style.IndentSpacing = 12.0
}

func darkTheme(style *imgui.Style) {
	style.SetColor(imgui.StyleColorText, imgui.Vec4{0.80, 0.80, 0.83, 1.00})
	style.SetColor(imgui.StyleColorTextDisabled, imgui.Vec4{0.24, 0.23, 0.29, 1.00})
	style.SetColor(imgui.StyleColorWindowBg, imgui.Vec4{0.06, 0.05, 0.07, 1.00})
	style.SetColor(imgui.StyleColorPopupBg, imgui.Vec4{0.07, 0.07, 0.09, 1.00})
	style.SetColor(imgui.StyleColorBorder, imgui.Vec4{0.80, 0.80, 0.83, 0.88})
	style.SetColor(imgui.StyleColorBorderShadow, imgui.Vec4{0.92, 0.91, 0.88, 0.00})
	style.SetColor(imgui.StyleColorFrameBg, imgui.Vec4{0.10, 0.09, 0.12, 1.00})
	style.SetColor(imgui.StyleColorFrameBgHovered, imgui.Vec4{0.24, 0.23, 0.29, 1.00})
	style.SetColor(imgui.StyleColorFrameBgActive, imgui.Vec4{0.56, 0.56, 0.58, 1.00})
	style.SetColor(imgui.StyleColorTitleBg, imgui.Vec4{0.10, 0.09, 0.12, 1.00})
	style.SetColor(imgui.StyleColorTitleBgCollapsed, imgui.Vec4{1.00, 0.98, 0.95, 0.75})
	style.SetColor(imgui.StyleColorTitleBgActive, imgui.Vec4{0.07, 0.07, 0.09, 1.00})
	style.SetColor(imgui.StyleColorMenuBarBg, imgui.Vec4{0.10, 0.09, 0.12, 1.00})
	style.SetColor(imgui.StyleColorScrollbarBg, imgui.Vec4{0.10, 0.09, 0.12, 1.00})
	style.SetColor(imgui.StyleColorScrollbarGrab, imgui.Vec4{0.80, 0.80, 0.83, 0.31})
	style.SetColor(imgui.StyleColorScrollbarGrabHovered, imgui.Vec4{0.56, 0.56, 0.58, 1.00})
	style.SetColor(imgui.StyleColorScrollbarGrabActive, imgui.Vec4{0.06, 0.05, 0.07, 1.00})
	//style.SetColor(imgui.StyleColorComboBg, imgui.Vec4{0.19, 0.18, 0.21, 1.00})
	style.SetColor(imgui.StyleColorCheckMark, imgui.Vec4{0.80, 0.80, 0.83, 0.31})
	style.SetColor(imgui.StyleColorSliderGrab, imgui.Vec4{0.80, 0.80, 0.83, 0.31})
	style.SetColor(imgui.StyleColorSliderGrabActive, imgui.Vec4{0.06, 0.05, 0.07, 1.00})
	style.SetColor(imgui.StyleColorButton, imgui.Vec4{0.10, 0.09, 0.12, 1.00})
	style.SetColor(imgui.StyleColorButtonHovered, imgui.Vec4{0.24, 0.23, 0.29, 1.00})
	style.SetColor(imgui.StyleColorButtonActive, imgui.Vec4{0.56, 0.56, 0.58, 1.00})
	style.SetColor(imgui.StyleColorHeader, imgui.Vec4{0.10, 0.09, 0.12, 1.00})
	style.SetColor(imgui.StyleColorHeaderHovered, imgui.Vec4{0.56, 0.56, 0.58, 1.00})
	style.SetColor(imgui.StyleColorHeaderActive, imgui.Vec4{0.06, 0.05, 0.07, 1.00})
	//style.SetColor(imgui.StyleColorColumn, imgui.Vec4{0.56, 0.56, 0.58, 1.00})
	//style.SetColor(imgui.StyleColorColumnHovered, imgui.Vec4{0.24, 0.23, 0.29, 1.00})
	//style.SetColor(imgui.StyleColorColumnActive, imgui.Vec4{0.56, 0.56, 0.58, 1.00})
	style.SetColor(imgui.StyleColorResizeGrip, imgui.Vec4{0.00, 0.00, 0.00, 0.00})
	style.SetColor(imgui.StyleColorResizeGripHovered, imgui.Vec4{0.56, 0.56, 0.58, 1.00})
	style.SetColor(imgui.StyleColorResizeGripActive, imgui.Vec4{0.06, 0.05, 0.07, 1.00})
	//style.SetColor(imgui.StyleColorCloseButton, imgui.Vec4{0.40, 0.39, 0.38, 0.16})
	//style.SetColor(imgui.StyleColorCloseButtonHovered, imgui.Vec4{0.40, 0.39, 0.38, 0.39})
	//style.SetColor(imgui.StyleColorCloseButtonActive, imgui.Vec4{0.40, 0.39, 0.38, 1.00})
	style.SetColor(imgui.StyleColorPlotLines, imgui.Vec4{0.40, 0.39, 0.38, 0.63})
	style.SetColor(imgui.StyleColorPlotLinesHovered, imgui.Vec4{0.25, 1.00, 0.00, 1.00})
	style.SetColor(imgui.StyleColorPlotHistogram, imgui.Vec4{0.40, 0.39, 0.38, 0.63})
	style.SetColor(imgui.StyleColorPlotHistogramHovered, imgui.Vec4{0.25, 1.00, 0.00, 1.00})
	style.SetColor(imgui.StyleColorTextSelectedBg, imgui.Vec4{0.25, 1.00, 0.00, 0.43})
	style.SetColor(imgui.StyleColorModalWindowDarkening, imgui.Vec4{1.00, 0.98, 0.95, 0.73})
}
