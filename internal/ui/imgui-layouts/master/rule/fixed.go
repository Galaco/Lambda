package rule

import (
	"github.com/galaco/Lambda/internal/ui/imgui-layouts/master"
)

type RuleFixedWidth struct {
	Width int
}

func (rule *RuleFixedWidth) Resolve(panel *master.Panel) {
	panel.Size[0] = float32(rule.Width)
}

func NewRuleFixedWidth(width int) *RuleFixedWidth {
	return &RuleFixedWidth{
		Width: width,
	}
}

type RuleFixedHeight struct {
	Height int
	asPercent bool
	fixedOffset int
}

func (rule *RuleFixedHeight) Resolve(panel *master.Panel) {
	if !rule.asPercent {
		panel.Size[1] = float32(rule.Height)
	} else {
		panel.Size[1] = (float32(rule.Height)/100) * panel.WindowSize[1]
	}
	panel.Size[1] = panel.Size[1] - float32(rule.fixedOffset)
}

func NewRuleFixedHeight(height int, asPercent bool, fixedOffset int) *RuleFixedHeight {
	return &RuleFixedHeight{
		Height: height,
		asPercent: asPercent,
		fixedOffset: fixedOffset,
	}
}
