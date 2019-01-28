package renderer

type Window struct {
	width int
	height int
	frameBuffer *fbo
}


func (win* Window) DrawFrame() {
	win.frameBuffer.Bind()

	//RENDER

	win.frameBuffer.Unbind()
}

func (win* Window) SetSize(width int, height int) {
	win.width = width
	win.height = height
	win.frameBuffer.Destroy()
	win.frameBuffer = NewFbo(width, height)
	win.frameBuffer.Bind()
}

func NewWindow(width int, height int) *Window {
	return &Window {
		width: width,
		height: height,
		frameBuffer: NewFbo(width, height),
	}
}