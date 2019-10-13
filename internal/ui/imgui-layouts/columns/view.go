package columns

import (
	"errors"
	"github.com/inkyblackness/imgui-go"
)

// View represents a column view. It wraps imgui's Column functionality to
// provide a simplified mechanism for rendering columns.
type View struct {
	numColumns     int
	columnContents []func()
	columnWidths   []*columnWidth
}

// Render draws the current configuration of columns
func (view *View) Render(width int, height int) {
	imgui.Columns(view.numColumns, "Columns")

	for idx, renderFunc := range view.columnContents {
		if !view.columnWidths[idx].AsPercentage {
			imgui.SetColumnWidth(idx, float32(view.columnWidths[idx].Width))
			renderFunc()
		} else {
			imgui.PushItemWidth(float32(view.columnWidths[idx].Width))
			renderFunc()
			imgui.PopItemWidth()
		}

		imgui.NextColumn()
	}
}

// SetColumnContents receives a renderable function to associate with a
// particular column, and how that column should be sized.
func (view *View) SetColumnContents(idx int, renderFunc func(), width *columnWidth) error {
	if idx < 0 || idx > view.numColumns {
		return errors.New("column index out of bounds")
	}

	view.columnContents[idx] = renderFunc

	if width != nil {
		view.columnWidths[idx] = width
	} else {
		view.columnWidths[idx] = NewColumnWidth(-1, true)
	}

	return nil
}

// NewColumns returns a new columns view
func NewColumns(num int) *View {
	return &View{
		numColumns:     num,
		columnContents: make([]func(), num),
		columnWidths:   make([]*columnWidth, num),
	}
}

// columnWidth provides data about how a column should be sized.
type columnWidth struct {
	AsPercentage bool
	Width        int
}

// NewColumnWidth returns a new columnWidth struct
func NewColumnWidth(width int, asPercentage bool) *columnWidth {
	return &columnWidth{
		Width:        width,
		AsPercentage: asPercentage,
	}
}
