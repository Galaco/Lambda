package filesystem

import (
	lambdaFS "github.com/galaco/Lambda-Core/core/filesystem"
	"github.com/galaco/Lambda-Core/core/logger"
	"github.com/galaco/Lambda-Core/core/resource"
	"github.com/galaco/Lambda-Core/lib/gameinfo"
)

func New(gameDir string) *lambdaFS.FileSystem {
	gameInfo, err := gameinfo.LoadConfig(gameDir)
	if err != nil {
		logger.Fatal(err)
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
