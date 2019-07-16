package opengl

import (
	"github.com/galaco/gosigl"
	"github.com/galaco/lambda-core/lib/util"
	"github.com/go-gl/gl/v4.1-core/gl"
	"io/ioutil"
	"log"
	"unsafe"
)

type OpenGL struct {
}

func (ogl *OpenGL) Init() {
	if err := gl.Init(); err != nil {
		panic(err)
	}
}

func (ogl *OpenGL) LambdaCreateTextureStorage2D(id *uint32, width, height int32) {
	ogl.GenTextures(1, id)
	ogl.BindTexture(gl.TEXTURE_2D, *id)
	gl.TexStorage2D(gl.TEXTURE_2D, 1, gl.RGBA8, width, height)

	ogl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	ogl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	ogl.BindTexture(gl.TEXTURE_2D, 0)
}

func (ogl *OpenGL) LambdaCreateTexture2D(id *uint32, width, height int32, data []byte) {
	ogl.GenTextures(1, id)
	ogl.BindTexture(gl.TEXTURE_2D, *id)

	var p unsafe.Pointer
	if data != nil && len(data) > 0 {
		p = gl.Ptr(data)
	} else {
		p = nil
	}

	ogl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA8, width, height, 0, gl.RGBA, gl.UNSIGNED_BYTE, p)

	ogl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	ogl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	ogl.BindTexture(gl.TEXTURE_2D, 0)
}

func (ogl *OpenGL) LambdaBindTexture2DToFramebuffer(texId uint32) {
	ogl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0, gl.TEXTURE_2D, texId, 0)
}

func (ogl *OpenGL) LambdaBindDepthBufferToFramebuffer(texId uint32) {
	gl.FramebufferRenderbuffer(gl.FRAMEBUFFER, gl.DEPTH_ATTACHMENT, gl.RENDERBUFFER, texId)
}

func (ogl *OpenGL) LambdaBindColourBufferToFramebuffer(texId uint32) {
	gl.FramebufferRenderbuffer(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0, gl.RENDERBUFFER, texId)
}

func (ogl *OpenGL) LambdaBindTexture2D(id uint32) {
	ogl.BindTexture(gl.TEXTURE_2D, id)
}

func (ogl *OpenGL) LambdaBindFramebuffer(framebufferId uint32) {
	ogl.BindFramebuffer(gl.FRAMEBUFFER, framebufferId)
}

func (ogl *OpenGL) LambdaDrawBuffers() {
	gl.DrawBuffer(gl.COLOR_ATTACHMENT0)
}

func (ogl *OpenGL) LambdaCreateRenderbufferStorageMultiSampledColour(width, height int32) uint32 {
	var buf uint32
	msaa := int32(4)

	gl.GenRenderbuffers(1, &buf)
	gl.BindRenderbuffer(gl.RENDERBUFFER, buf)
	gl.RenderbufferStorageMultisample(gl.RENDERBUFFER, msaa, gl.RGB8, width, height)

	return buf
}

func (ogl *OpenGL) LambdaCreateRenderbufferStorageMultiSampledDepth(width, height int32) uint32 {
	var buf uint32
	msaa := int32(4)

	gl.GenRenderbuffers(1, &buf)
	gl.BindRenderbuffer(gl.RENDERBUFFER, buf)
	gl.RenderbufferStorageMultisample(gl.RENDERBUFFER, msaa, gl.DEPTH_COMPONENT, width, height)

	return buf
}

func (ogl *OpenGL) LambdaCreateRenderbufferStorageColour(width, height int32) uint32 {
	var buf uint32

	gl.GenRenderbuffers(1, &buf)
	gl.BindRenderbuffer(gl.RENDERBUFFER, buf)
	gl.RenderbufferStorage(gl.RENDERBUFFER, gl.RGB8, width, height)

	return buf
}

func (ogl *OpenGL) LambdaCreateRenderbufferStorageDepth(width, height int32) uint32 {
	var buf uint32

	gl.GenRenderbuffers(1, &buf)
	gl.BindRenderbuffer(gl.RENDERBUFFER, buf)
	gl.RenderbufferStorage(gl.RENDERBUFFER, gl.DEPTH_COMPONENT, width, height)

	return buf
}

func (ogl *OpenGL) LambdaLoadSimpleShader(path string) *gosigl.Context {
	vs, err := ioutil.ReadFile(path + ".vs.glsl")
	if err != nil {
		util.Logger().Panic(err)
	}
	fs, err := ioutil.ReadFile(path + ".fs.glsl")
	if err != nil {
		util.Logger().Panic(err)
	}

	shader := gosigl.NewShader()
	if err = shader.AddShader(string(vs)+"\x00", gosigl.VertexShader); err != nil {
		util.Logger().Panic(err)
	}
	if err = shader.AddShader(string(fs)+"\x00", gosigl.FragmentShader); err != nil {
		util.Logger().Panic(err)
	}
	shader.Finalize()

	return &shader
}

func (ogl *OpenGL) Viewport(x, y, width, height int32) {
	gl.Viewport(x, y, width, height)
}

func (ogl *OpenGL) ClearColor(r, g, b, a float32) {
	gl.ClearColor(r, g, b, a)
}

func (ogl *OpenGL) Clear(mask uint32) {
	gl.Clear(mask)
}

func (ogl *OpenGL) ClearAll() {
	ogl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (ogl *OpenGL) CreateFramebuffers(n int32, framebuffers *uint32) {
	gl.GenFramebuffers(n, framebuffers)
}
func (ogl *OpenGL) BindFramebuffer(target uint32, framebuffer uint32) {
	gl.BindFramebuffer(target, framebuffer)
}
func (ogl *OpenGL) DeleteFramebuffers(n int32, framebuffers *uint32) {
	gl.DeleteFramebuffers(n, framebuffers)
}
func (ogl *OpenGL) FramebufferTexture(target uint32, attachment uint32, texture uint32, level int32) {
	gl.FramebufferTexture(target, attachment, texture, level)
}
func (ogl *OpenGL) FramebufferTexture2D(target uint32, attachment uint32, textarget uint32, texture uint32, level int32) {
	gl.FramebufferTexture2D(target, attachment, textarget, texture, level)
}

func (ogl *OpenGL) DeleteRenderBuffer(n int32, target *uint32) {
	gl.DeleteRenderbuffers(n, target)
}

// Textures

func (ogl *OpenGL) DeleteTextures(n int32, textures *uint32) {
	gl.DeleteTextures(n, textures)
}
func (ogl *OpenGL) GenTextures(n int32, textures *uint32) {
	gl.GenTextures(n, textures)
}
func (ogl *OpenGL) BindTexture(target uint32, texture uint32) {
	gl.BindTexture(target, texture)
}
func (ogl *OpenGL) TexImage2D(target uint32, level int32, internalformat int32, width int32, height int32, border int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	gl.TexImage2D(target, level, internalformat, width, height, border, format, xtype, pixels)
}
func (ogl *OpenGL) TexParameteri(target uint32, pname uint32, param int32) {
	gl.TexParameteri(target, pname, param)
}

func (ogl *OpenGL) EnableBlend() {
	gosigl.EnableBlend()
}
func (ogl *OpenGL) EnableDepthTest() {
	gosigl.EnableDepthTest()
}
func (ogl *OpenGL) EnableCullFaceBack() {
	gosigl.EnableCullFace(gosigl.Back, gosigl.WindingClockwise)
}

func (ogl *OpenGL) DrawTriangleArray(offset int32, count int32) {
	gl.DrawArrays(gl.TRIANGLES, offset, count)
}

func (ogl *OpenGL) SendUniformMat4(uniform int32, matrix *float32) {
	gl.UniformMatrix4fv(uniform, 1, false, matrix)
}

func (ogl *OpenGL) Error() bool {
	if err := gl.GetError(); err != 0 {
		log.Println(err)
		return true
	}
	return false
}
