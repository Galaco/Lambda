package filesystem

import (
	lambdaFS "github.com/galaco/lambda-core/filesystem"
	"github.com/galaco/lambda-core/lib/gameinfo"
	"github.com/galaco/lambda-core/lib/util"
	"github.com/galaco/lambda-core/resource"
)

// New builds a new filesystem from a game directory root.
// It loads a gameinfo.txt and attempts to find listed resourced
// in it.
func New(gameDir string) lambdaFS.IFileSystem {
	gameInfo, err := gameinfo.LoadConfig(gameDir)
	if err != nil {
		util.Logger().Panic(err)
	}

	// Register GameInfo.txt referenced resource paths
	// Filesystem module needs to know about all the possible resource
	// locations it can search.
	fs := lambdaFS.CreateFilesystemFromGameInfoDefinitions(gameDir, gameInfo)

	// Explicity define fallbacks for missing resources
	// Defaults are defined, but if HL2 assets are not readable, then
	// the default may not be readable
	resource.Manager().SetErrorModelName("models/props/de_dust/du_antenna_A.mdl")
	resource.Manager().SetErrorTextureName("materials/error.vtf")

	return fs
}
