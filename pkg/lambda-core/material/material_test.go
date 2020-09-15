package material

import (
	"github.com/galaco/Lambda/pkg/lambda-core/texture"
	"testing"
)

func TestMaterial_FilePath(t *testing.T) {
	sut := Material{
		filePath: "foo/bar.vmt",
	}

	if sut.FilePath() != "foo/bar.vmt" {
		t.Errorf("incorrect filepath returned. Expected %s, but received: %s", "foo/bar.vmt", sut.FilePath())
	}
}

func TestMaterial_Height(t *testing.T) {
	sut := Material{
		filePath: "foo/bar.vmt",
	}
	sut.Textures.Albedo = texture.NewError("error.vtf")

	if sut.Height() != sut.Textures.Albedo.Height() {
		t.Error("material height doesnt match basetextures height")
	}
}

func TestMaterial_Width(t *testing.T) {
	sut := Material{
		filePath: "foo/bar.vmt",
	}
	sut.Textures.Albedo = texture.NewError("error.vtf")

	if sut.Width() != sut.Textures.Albedo.Width() {
		t.Error("material width doesnt match basetextures width")
	}
}
