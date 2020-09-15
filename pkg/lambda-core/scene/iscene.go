package scene

import (
	"github.com/galaco/Lambda/pkg/lambda-core/model"
)

// IScene
type IScene interface {
	// Bsp
	Bsp() *model.Bsp
	// StaticProps
	StaticProps() []model.StaticProp
}
