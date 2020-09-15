package resource

// IResource Generic filesystem object. If it was loaded from a path, it should
// implement this.
type IResource interface {
	FilePath() string
}
