package master

import "github.com/go-gl/mathgl/mgl32"

const (
	axisWidth  = 0
	axisHeight = 1
)

// Panel represents a rectangle positioned within a window.
// It is provided a set of rules that resolve the rectangles size and
// position on screen.
type Panel struct {
	// TopLeft is the computed offset from the top left after rule resolution
	TopLeft mgl32.Vec2
	// BottomRight is the computed offset from the  bottom right after rule resolution
	BottomRight mgl32.Vec2
	// Size is the computed width and height
	Size mgl32.Vec2

	// Rules is the ruleset to compute layout information from
	Rules []Rulable
}

// AddRules adds a new rule to the resolver
func (pane *Panel) AddRule(rule Rulable) {
	pane.Rules = append(pane.Rules, rule)
}

// Resolve generates this panels position and size based on the provided
// ruleset.
// Incompatible rules wont error, but unexpected results may be generated.
func (pane *Panel) Resolve(width int, height int) (offset, size mgl32.Vec2) {
	for _, rule := range pane.Rules {
		rule.Resolve(pane)
	}

	// Resolve size
	size[0] = pane.resolveAxis(width, axisWidth)
	size[1] = pane.resolveAxis(height, axisHeight)

	// Resolve offset
	offset[0] = pane.resolveOffset(int(width), int(size[0]), axisWidth)
	offset[1] = pane.resolveOffset(int(height), int(size[1]), axisHeight)

	return offset, size
}

// resolveAxis computes the true size of a given axis based on rule resolution
// data.
func (pane *Panel) resolveAxis(size int, axis int) float32 {
	// Fixed width
	if pane.Size[axis] != -1 {
		return pane.Size[axis]
	}

	// Compute width based on clamping
	computedSize := float32(size)
	if pane.TopLeft[axis] != -1 {
		computedSize -= pane.TopLeft[axis]
	}
	if pane.BottomRight[axis] != -1 {
		computedSize -= pane.BottomRight[axis]
	}

	return computedSize
}

// resolveOffset resolves the position onscreen for a given axis
func (pane *Panel) resolveOffset(size, computedSize, axis int) float32 {
	// Offset top-left corner
	if pane.TopLeft[axis] != -1 {
		return pane.TopLeft[axis]
	}
	// Offset bottom-right
	if pane.BottomRight[axis] != -1 {
		return float32(size - (computedSize + int(pane.BottomRight[axis])))
	}

	return 0
}

// NewPanel returns a new panel
func NewPanel() *Panel {
	return &Panel{
		TopLeft:     mgl32.Vec2{-1, -1},
		BottomRight: mgl32.Vec2{-1, -1},
		Size:        mgl32.Vec2{-1, -1},
	}
}
