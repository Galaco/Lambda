package renderer

import (
	"github.com/galaco/Lambda-Core/core/logger"
	"github.com/galaco/gosigl"
)

func LoadShader() *gosigl.Context {
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

layout(location = 0) in vec3 vertex;
layout(location = 1) in vec3 normal;
layout(location = 2) in vec2 uv;
layout(location = 3) in vec3 tangent;

out vec2 UV;

void main() {
	gl_Position = projection * view * model * vec4(vertex, 1.0);
	
	UV = uv;
}
` + "\x00"

//language=glsl
var fragmentSource = `
#version 410

uniform sampler2D albedoSampler;

in vec2 UV;

out vec4 frag_colour;

void AddAlbedo(inout vec4 fragColour, in sampler2D sampler, in vec2 uv) 
{
	fragColour = texture(sampler, uv).rgba;
}

void main() {
	AddAlbedo(frag_colour, albedoSampler, UV);
}
` + "\x00"
