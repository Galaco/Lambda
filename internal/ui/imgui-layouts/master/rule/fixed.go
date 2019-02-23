package rule

import (
	"github.com/galaco/Lambda/ui/imgui-layouts/master"
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
