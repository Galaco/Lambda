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

uniform mat4 projection;
uniform mat4 view;
uniform mat4 model;

layout(location = 0) in vec3 vertexPosition_modelspace;

void main() {
	//gl_Position = projection * view * model * vec4(vertexPosition_modelspace, 1.0);
	gl_Position = vec4(vertexPosition_modelspace, 1.0);
}
` + "\x00"

//language=glsl
var fragmentSource = `
#version 410

out vec4 color;
void main() {
  color = vec4(1,0,0, 0.03);
}
` + "\x00"
