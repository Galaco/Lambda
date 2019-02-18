package graphics

import "unsafe"

type Adapter interface {
	Init()

	//Lambda simplified
	LambdaCreateTexture2D(id *uint32, width, height int32, data []byte)
	LambdaBindTexture2DToFramebuffer(framebufferId uint32)
	LambdaBindFramebuffer(framebufferId uint32)
	LambdaBindTexture2D(id uint32)

	// General
	Viewport(x, y, width, height int32)
	ClearColor(r, g, b, a float32)
	Clear(mask uint32)
	ClearAll()

	// Framebuffer
	CreateFramebuffers(n int32, framebuffers *uint32)
	BindFramebuffer(target uint32, framebuffer uint32)
	DeleteFramebuffers(n int32, framebuffers *uint32)
	FramebufferTexture2D(target uint32, attachment uint32, textarget uint32, texture uint32, level int32)

	// Texture
	DeleteTextures(n int32, textures *uint32)
	GenTextures(n int32, textures *uint32)
	BindTexture(target uint32, texture uint32)
	TexImage2D(target uint32, level int32, internalformat int32, width int32, height int32, border int32, format uint32, xtype uint32, pixels unsafe.Pointer)
	TexParameteri(target uint32, pname uint32, param int32)

	// Drawing
	DrawTriangleArray(offset int32, count int32)

	// Misc
	EnableBlend()
	EnableDepthTest()
	EnableCullFaceBack()

	// Uniforms
	SendUniformMat4(uniform int32, matrix *float32)

	Error() bool
}
