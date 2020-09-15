package material

type IMaterial interface {
	//Bind()
	Width() int
	Height() int
	FilePath() string
}
