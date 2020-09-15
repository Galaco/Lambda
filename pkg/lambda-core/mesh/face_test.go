package mesh

import (
	"github.com/galaco/Lambda/pkg/lambda-core/material"
	"github.com/galaco/Lambda/pkg/lambda-core/texture"
	"reflect"
	"testing"
)

func TestFace_IsLightmapped(t *testing.T) {
	sut := Face{}
	sut.AddLightmap(&texture.Lightmap{})

	if sut.IsLightmapped() != true {
		t.Error("face has lightmap, but isn't marked as being lightmapped")
	}
}

func TestFace_Length(t *testing.T) {
	sut := NewFace(32, 64, nil, nil)

	if sut.Length() != 64 {
		t.Error("unexpected length for face")
	}
}

func TestFace_AddMaterial(t *testing.T) {
	sut := Face{}
	expected := material.NewMaterial("foo.vmt")
	sut.AddMaterial(expected)

	if expected != sut.Material() {
		t.Error("unexpected material applied to face")
	}
}

func TestFace_AddLightmap(t *testing.T) {
	sut := Face{}
	expected := &texture.Lightmap{}
	sut.AddLightmap(expected)

	if expected != sut.Lightmap() {
		t.Error("unexpected lightmap applied to face")
	}
}

func TestFace_Lightmap(t *testing.T) {
	sut := Face{}
	expected := &texture.Lightmap{}
	sut.AddLightmap(expected)

	if expected != sut.Lightmap() {
		t.Error("unexpected lightmap applied to face")
	}
}

func TestFace_Material(t *testing.T) {
	sut := Face{}
	expected := material.NewMaterial("foo.vmt")
	sut.AddMaterial(expected)

	if expected != sut.Material() {
		t.Error("unexpected material applied to face")
	}
}

func TestFace_Offset(t *testing.T) {
	sut := NewFace(32, 64, nil, nil)

	if sut.Offset() != 32 {
		t.Error("unexpected offset for face")
	}
}

func TestNewFace(t *testing.T) {
	sut := NewFace(32, 64, nil, nil)
	if reflect.TypeOf(sut) != reflect.TypeOf(Face{}) {
		t.Errorf("unexpceted type returned. Expected %s, but received: %s", reflect.TypeOf(Face{}), reflect.TypeOf(sut))
	}
}
