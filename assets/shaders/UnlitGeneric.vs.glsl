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