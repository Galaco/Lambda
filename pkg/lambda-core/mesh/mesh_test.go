package mesh

import (
	"github.com/galaco/Lambda/pkg/lambda-core/material"
	"github.com/galaco/Lambda/pkg/lambda-core/texture"
	"reflect"
	"testing"
)

func TestNewMesh(t *testing.T) {
	if reflect.TypeOf(NewMesh()) != reflect.TypeOf(&Mesh{}) {
		t.Errorf("unexpected type returned for NewMesh. Expected: %s, but received: %s", reflect.TypeOf(&Mesh{}), reflect.TypeOf(NewMesh()))
	}
}

func TestMesh_AddLightmapCoordinate(t *testing.T) {
	sut := Mesh{}
	expected := []float32{
		1, 2, 3, 4,
	}
	sut.AddLightmapCoordinate(expected...)

	for i := 0; i < len(expected); i++ {
		if sut.LightmapCoordinates()[i] != expected[i] {
			t.Error("unexpected lightmap coordinate")
		}
	}
}

func TestMesh_AddNormal(t *testing.T) {
	sut := Mesh{}
	expected := []float32{
		1, 2, 3, 4,
	}
	sut.AddNormal(expected...)

	for i := 0; i < len(expected); i++ {
		if sut.Normals()[i] != expected[i] {
			t.Error("unexpected normal")
		}
	}
}

func TestMesh_AddTextureCoordinate(t *testing.T) {
	sut := Mesh{}
	expected := []float32{
		1, 2, 3, 4,
	}
	sut.AddUV(expected...)

	for i := 0; i < len(expected); i++ {
		if sut.UVs()[i] != expected[i] {
			t.Error("unexpected texture coordinate")
		}
	}
}

func TestMesh_AddVertex(t *testing.T) {
	sut := Mesh{}
	expected := []float32{
		1, 2, 3, 4,
	}
	sut.AddVertex(expected...)

	for i := 0; i < len(expected); i++ {
		if sut.Vertices()[i] != expected[i] {
			t.Error("unexpected vertex")
		}
	}
}

func TestMesh_Lightmap(t *testing.T) {
	sut := Mesh{}
	expected := &texture.Lightmap{}
	sut.SetLightmap(expected)

	if expected != sut.Lightmap() {
		t.Error("unexpected lightmap applied to mesh")
	}
}

func TestMesh_Material(t *testing.T) {
	sut := Mesh{}
	expected := material.NewMaterial("foo.vmt")
	sut.SetMaterial(expected)

	if expected != sut.Material() {
		t.Error("unexpected material applied to mesh")
	}
}

func TestMesh_LightmapCoordinates(t *testing.T) {
	sut := Mesh{}
	expected := []float32{
		1, 2, 3, 4,
	}
	sut.AddLightmapCoordinate(expected...)

	for i := 0; i < len(expected); i++ {
		if sut.LightmapCoordinates()[i] != expected[i] {
			t.Error("unexpected lightmap coordinate")
		}
	}
}

func TestMesh_Normals(t *testing.T) {
	sut := Mesh{}
	expected := []float32{
		1, 2, 3, 4,
	}
	sut.AddNormal(expected...)

	for i := 0; i < len(expected); i++ {
		if sut.Normals()[i] != expected[i] {
			t.Error("unexpected normal")
		}
	}
}

func TestMesh_SetLightmap(t *testing.T) {
	sut := Mesh{}
	expected := &texture.Lightmap{}
	sut.SetLightmap(expected)

	if expected != sut.Lightmap() {
		t.Error("unexpected lightmap applied to mesh")
	}
}

func TestMesh_SetMaterial(t *testing.T) {
	sut := Mesh{}
	expected := material.NewMaterial("foo.vmt")
	sut.SetMaterial(expected)

	if expected != sut.Material() {
		t.Error("unexpected material applied to mesh")
	}
}

func TestMesh_TextureCoordinates(t *testing.T) {
	sut := Mesh{}
	expected := []float32{
		1, 2, 3, 4,
	}
	sut.AddUV(expected...)

	for i := 0; i < len(expected); i++ {
		if sut.UVs()[i] != expected[i] {
			t.Error("unexpected texture coordinate")
		}
	}
}

func TestMesh_Vertices(t *testing.T) {
	sut := Mesh{}
	expected := []float32{
		1, 2, 3, 4,
	}
	sut.AddVertex(expected...)

	for i := 0; i < len(expected); i++ {
		if sut.Vertices()[i] != expected[i] {
			t.Error("unexpected vertex")
		}
	}
}
