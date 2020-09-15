package game

// IGame interface represents a game
type IGame interface {
	// RegisterEntityClasses should setup any game entity classes
	// for use with the engine when loading entdata
	RegisterEntityClasses()
}
