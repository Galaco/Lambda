package rule

import (
	"github.com/galaco/Lambda/internal/ui/imgui-layouts/master"
)

const (
	ClampTop    = 0
	ClampRight  = 1
	ClampBottom = 2
	ClampLeft   = 3
)

type RuleClampToEdge struct {
	Edge   int
	Offset int
}

func (rule *RuleClampToEdge) Resolve(panel *master.Panel) {
	switch rule.Edge {
	case ClampTop:
		panel.TopLeft[1] = float32(rule.Offset)
	case ClampLeft:
		panel.TopLeft[0] = float32(rule.Offset)
	case ClampBottom:
		panel.BottomRight[1] = float32(rule.Offset)
	case ClampRight:
		panel.BottomRight[0] = float32(rule.Offset)
	}
}

func NewRuleClampToEdge(edge, offset int) *RuleClampToEdge {
	return &RuleClampToEdge{
		Edge:   edge,
		Offset: offset,
	}
}
