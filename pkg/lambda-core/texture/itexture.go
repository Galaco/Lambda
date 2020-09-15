package texture

// ITexture Interface for a GPU texture
type ITexture interface {
	Width() int
	Height() int
	Format() uint32
	PixelDataForFrame(int) []byte
	FilePath() string
	Thumbnail() []byte
}
