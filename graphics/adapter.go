package graphics

type Adapter interface {
	Viewport(x, y, width, height int32)
	ClearColor(r, g, b, a float32)
	Clear(mask uint32)
}
