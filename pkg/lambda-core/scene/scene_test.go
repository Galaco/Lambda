package scene

import (
	"github.com/galaco/Lambda/pkg/lambda-core/mesh"
	"github.com/galaco/Lambda/pkg/lambda-core/model"
	"reflect"
	"testing"
)

func TestNewScene(t *testing.T) {
	sut := NewScene(model.Bsp{}, make([]model.StaticProp, 0))
	if reflect.TypeOf(sut) != reflect.TypeOf(&Scene{}) {
		t.Error("unexpected type for NewScene")
	}
}

func TestScene_Bsp(t *testing.T) {
	bspMesh := mesh.NewMesh()
	bsp := model.NewBsp(bspMesh)

	sut := NewScene(*bsp, make([]model.StaticProp, 0))
	if sut.Bsp().Mesh() != bsp.Mesh() {
		t.Error("unexpected model contained in scene")
	}
}

func TestScene_StaticProps(t *testing.T) {
	bspMesh := mesh.NewMesh()
	prop := model.StaticProp{}
	bsp := model.NewBsp(bspMesh)

	sut := NewScene(*bsp, []model.StaticProp{prop, prop, prop})
	if len(sut.StaticProps()) != 3 {
		t.Error("unexpected props returned")
	}
}
