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