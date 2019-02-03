package renderer

import (
	"github.com/galaco/Lambda-Core/core/logger"
	"github.com/galaco/gosigl"
)

func loadShader() *gosigl.Context {
	shader := gosigl.NewShader()
	err := shader.AddShader(vertexSource, gosigl.VertexShader)
	if err != nil {
		logger.Fatal(err)
	}
	err = shader.AddShader(fragmentSource, gosigl.FragmentShader)
	if err != nil {
		logger.Fatal(err)
	}
	shader.Finalize()

	return &shader
}

// language=glsl
var vertexSource = `
#version 410

layout(location = 0) in vec3 vertexPosition_modelspace;

void main() {
  gl_Position.xyz = vertexPosition_modelspace;
  gl_Position.w = 1.0;
}
` + "\x00"

//language=glsl
var fragmentSource = `
#version 410

out vec3 color;
void main() {
  color = vec3(1,0,0);
}
` + "\x00"