package loader

import (
	"github.com/galaco/Lambda/pkg/lambda-core/event"
	matloader "github.com/galaco/Lambda/pkg/lambda-core/loader/material"
	"github.com/galaco/Lambda/pkg/lambda-core/material"
	"github.com/galaco/Lambda/pkg/lambda-core/mesh"
	"github.com/galaco/Lambda/pkg/lambda-core/model"
	"github.com/galaco/Lambda/pkg/lambda-core/resource"
	"github.com/galaco/Lambda/pkg/lambda-core/resource/message"
	"github.com/galaco/Lambda/pkg/lambda-core/scene"
	"github.com/galaco/Lambda/pkg/lambda-core/texture"
	"github.com/galaco/bsp"
	"github.com/galaco/bsp/lumps"
	"github.com/galaco/bsp/primitives/common"
	"github.com/galaco/bsp/primitives/dispinfo"
	"github.com/galaco/bsp/primitives/dispvert"
	"github.com/galaco/bsp/primitives/face"
	"github.com/galaco/bsp/primitives/plane"
	"github.com/galaco/bsp/primitives/texinfo"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/golang-source-engine/filesystem"
	"math"
	"strings"
	"unsafe"
)

type bspstructs struct {
	faces     []face.Face
	planes    []plane.Plane
	vertexes  []mgl32.Vec3
	surfEdges []int32
	edges     [][2]uint16
	texInfos  []texinfo.TexInfo
	dispInfos []dispinfo.DispInfo
	dispVerts []dispvert.DispVert
	game      *lumps.Game
	lightmap  []common.ColorRGBExponent32
}

// LoadMap is the gateway into loading the core static level. Entities are loaded
// elsewhere
// It loads in the following order:
// BSP Geometry
// BSP Materials
// StaticProps (materials loaded as required)
func LoadMap(fs *filesystem.FileSystem, file *bsp.Bsp) scene.IScene {
	ResourceManager := resource.Manager()
	bspStructure := bspstructs{
		faces:     file.Lump(bsp.LumpFaces).(*lumps.Face).GetData(),
		planes:    file.Lump(bsp.LumpPlanes).(*lumps.Planes).GetData(),
		vertexes:  file.Lump(bsp.LumpVertexes).(*lumps.Vertex).GetData(),
		surfEdges: file.Lump(bsp.LumpSurfEdges).(*lumps.Surfedge).GetData(),
		edges:     file.Lump(bsp.LumpEdges).(*lumps.Edge).GetData(),
		texInfos:  file.Lump(bsp.LumpTexInfo).(*lumps.TexInfo).GetData(),
		dispInfos: file.Lump(bsp.LumpDispInfo).(*lumps.DispInfo).GetData(),
		dispVerts: file.Lump(bsp.LumpDispVerts).(*lumps.DispVert).GetData(),
		game:      file.Lump(bsp.LumpGame).(*lumps.Game),
		lightmap:  file.Lump(bsp.LumpLighting).(*lumps.Lighting).GetData(),
	}

	//MATERIALS
	stringTable := matloader.LoadMaterials(
		fs,
		file.Lump(bsp.LumpTexDataStringData).(*lumps.TexDataStringData),
		file.Lump(bsp.LumpTexDataStringTable).(*lumps.TexDataStringTable),
		&bspStructure.texInfos)

	// BSP FACES
	bspMesh := mesh.NewMesh()
	bspObject := model.NewBsp(bspMesh)
	bspFaces := make([]mesh.Face, len(bspStructure.faces))
	dispFaces := make([]int, 0)
	lightMaps := make([]texture.ITexture, len(bspFaces))

	for idx, f := range bspStructure.faces {
		if f.DispInfo > -1 {
			// This face is a displacement
			bspFaces[idx] = generateDisplacementFace(&f, &bspStructure, bspMesh)
			dispFaces = append(dispFaces, idx)
		} else {
			bspFaces[idx] = generateBspFace(&f, &bspStructure, bspMesh)
		}

		// Prepare lightmaps for each face
		if len(bspStructure.lightmap) < 1 {
			continue
		}

		lightMaps[idx] = texture.LightmapFromColorRGBExp32(
			int(f.LightmapTextureSizeInLuxels[0]+1),
			int(f.LightmapTextureSizeInLuxels[1]+1),
			lightmapSamplesFromFace(&f, &bspStructure.lightmap))

		//bspFaces[idx].AddLightmap(texture.LightmapFromColorRGBExp32(
		//	int(f.LightmapTextureSizeInLuxels[0]+1),
		//	int(f.LightmapTextureSizeInLuxels[1]+1),
		//	lightmapSamplesFromFace(&f, &bspStructure.lightmap)))
		//bspFaces[idx].Lightmap().Finish()
	}

	// Add MATERIALS TO FACES
	for idx, bspFace := range bspFaces {
		faceVmt, _ := stringTable.FindString(int(bspStructure.texInfos[bspStructure.faces[idx].TexInfo].TexData))
		var mat material.IMaterial
		if ResourceManager.HasMaterial(faceVmt) {
			mat = ResourceManager.Material(faceVmt).(material.IMaterial)
		} else {
			mat = ResourceManager.Material(resource.Manager().ErrorTextureName()).(material.IMaterial)
		}

		lightMat := bspFaces[idx].Lightmap()
		bspFaces[idx].AddMaterial(mat)
		// Generate texture coordinates
		bspMesh.AddUV(
			texCoordsForFaceFromTexInfo(
				bspMesh.Vertices()[bspFace.Offset()*3:(bspFace.Offset()*3)+(bspFace.Length()*3)],
				&bspStructure.texInfos[bspStructure.faces[idx].TexInfo], mat.Width(), mat.Height())...)
		if lightMat != nil {
			bspMesh.AddLightmapCoordinate(
				lightmapCoordsForFaceFromTexInfo(
					bspMesh.Vertices()[bspFace.Offset()*3:(bspFace.Offset()*3)+(bspFace.Length()*3)],
					&bspStructure.faces[idx],
					&bspStructure.texInfos[bspStructure.faces[idx].TexInfo], lightMat.Width(), lightMat.Height())...)
		}

		if strings.HasPrefix(faceVmt, "TOOLS/") {
			bspFaces[idx].AddMaterial(nil)
		}
	}

	//lightmapTexture := texture.NewAtlas(4096, 4096)
	//_,err := lightmapTexture.PackTextures(lightMaps, 1)
	//if err != nil {
	//	logger.Error(err)
	//}

	// Finish the bsp object.
	event.Manager().Dispatch(message.LoadedMap(bspObject))
	//bspMesh.Finish()

	cl := model.ClusterLeaf{
		Id:        0,
		Faces:     bspFaces,
		DispFaces: dispFaces,
	}
	bspObject.SetClusterLeafs([]model.ClusterLeaf{cl})

	// Get static props
	staticProps := LoadStaticProps(bspStructure.game.GetStaticPropLump(), fs)

	return scene.NewScene(*bspObject, staticProps)
}

// generateBspFace Create primitives from face data in the bsp
func generateBspFace(f *face.Face, bspStructure *bspstructs, bspMesh mesh.IMesh) mesh.Face {
	offset := int32(len(bspMesh.Vertices())) / 3
	length := int32(0)

	planeNormal := bspStructure.planes[f.Planenum].Normal
	// All surfedges associated with this face
	// surfEdges are basically indices into the edges lump
	faceSurfEdges := bspStructure.surfEdges[f.FirstEdge:(f.FirstEdge + int32(f.NumEdges))]
	rootIndex := uint16(0)
	for idx, surfEdge := range faceSurfEdges {
		edge := bspStructure.edges[int(math.Abs(float64(surfEdge)))]
		e1 := 0
		e2 := 1
		if surfEdge < 0 {
			e1 = 1
			e2 = 0
		}
		//Capture root indice
		if idx == 0 {
			rootIndex = edge[e1]
		} else {
			// Just create a triangle for every edge now
			bspMesh.AddVertex(bspStructure.vertexes[rootIndex].X(), bspStructure.vertexes[rootIndex].Y(), bspStructure.vertexes[rootIndex].Z())
			bspMesh.AddNormal(planeNormal.X(), planeNormal.Y(), planeNormal.Z())

			bspMesh.AddVertex(bspStructure.vertexes[edge[e1]].X(), bspStructure.vertexes[edge[e1]].Y(), bspStructure.vertexes[edge[e1]].Z())
			bspMesh.AddNormal(planeNormal.X(), planeNormal.Y(), planeNormal.Z())

			bspMesh.AddVertex(bspStructure.vertexes[edge[e2]].X(), bspStructure.vertexes[edge[e2]].Y(), bspStructure.vertexes[edge[e2]].Z())
			bspMesh.AddNormal(planeNormal.X(), planeNormal.Y(), planeNormal.Z())

			length += 3 // num verts (3 b/c face triangles
		}
	}

	return mesh.NewFace(offset, length, nil, nil)
}

// generateDisplacementFace Create Primitive from Displacement face
// This is based on:
// https://github.com/Metapyziks/VBspViewer/blob/master/Assets/VBspViewer/Scripts/Importing/VBsp/VBspFile.cs
func generateDisplacementFace(f *face.Face, bspStructure *bspstructs, bspMesh mesh.IMesh) mesh.Face {
	corners := make([]mgl32.Vec3, 4)
	normal := bspStructure.planes[f.Planenum].Normal

	info := bspStructure.dispInfos[f.DispInfo]
	size := int(1 << uint32(info.Power))
	firstCorner := int32(0)
	firstCornerDist2 := float32(math.MaxFloat32)

	offset := int32(len(bspMesh.Vertices())) / 3
	length := int32(0)

	for surfId := f.FirstEdge; surfId < f.FirstEdge+int32(f.NumEdges); surfId++ {
		surfEdge := bspStructure.surfEdges[surfId]
		edgeIndex := int32(math.Abs(float64(surfEdge)))
		edge := bspStructure.edges[edgeIndex]
		vert := bspStructure.vertexes[edge[0]]
		if surfEdge < 0 {
			vert = bspStructure.vertexes[edge[1]]
		}
		corners[surfId-f.FirstEdge] = vert

		dist2tmp := info.StartPosition.Sub(vert)
		dist2 := (dist2tmp.X() * dist2tmp.X()) + (dist2tmp.Y() * dist2tmp.Y()) + (dist2tmp.Z() * dist2tmp.Z())
		if dist2 < firstCornerDist2 {
			firstCorner = surfId - f.FirstEdge
			firstCornerDist2 = dist2
		}
	}

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			a := generateDispVert(int(info.DispVertStart), x, y, size, corners, firstCorner, &bspStructure.dispVerts)
			b := generateDispVert(int(info.DispVertStart), x, y+1, size, corners, firstCorner, &bspStructure.dispVerts)
			c := generateDispVert(int(info.DispVertStart), x+1, y+1, size, corners, firstCorner, &bspStructure.dispVerts)
			d := generateDispVert(int(info.DispVertStart), x+1, y, size, corners, firstCorner, &bspStructure.dispVerts)

			// Split into triangles
			bspMesh.AddVertex(a.X(), a.Y(), a.Z(), b.X(), b.Y(), b.Z(), c.X(), c.Y(), c.Z())
			bspMesh.AddNormal(normal.X(), normal.Y(), normal.Z(), normal.X(), normal.Y(), normal.Z(), normal.X(), normal.Y(), normal.Z())
			bspMesh.AddVertex(a.X(), a.Y(), a.Z(), c.X(), c.Y(), c.Z(), d.X(), d.Y(), d.Z())
			bspMesh.AddNormal(normal.X(), normal.Y(), normal.Z(), normal.X(), normal.Y(), normal.Z(), normal.X(), normal.Y(), normal.Z())

			length += 6 // 6 b/c quad = 2*triangle
		}
	}

	return mesh.NewFace(offset, length, nil, nil)
}

// generateDispVert Create a displacement vertex
func generateDispVert(offset int, x int, y int, size int, corners []mgl32.Vec3, firstCorner int32, dispVerts *[]dispvert.DispVert) mgl32.Vec3 {
	vert := (*dispVerts)[offset+x+y*(size+1)]

	tx := float32(x) / float32(size)
	ty := float32(y) / float32(size)
	sx := 1.0 - tx
	sy := 1.0 - ty

	cornerA := corners[(0+firstCorner)&3]
	cornerB := corners[(1+firstCorner)&3]
	cornerC := corners[(2+firstCorner)&3]
	cornerD := corners[(3+firstCorner)&3]

	origin := ((cornerB.Mul(sx).Add(cornerC.Mul(tx))).Mul(ty)).Add((cornerA.Mul(sx).Add(cornerD.Mul(tx))).Mul(sy))

	return origin.Add(vert.Vec.Mul(vert.Dist))
}

// texCoordsForFaceFromTexInfo Generate texturecoordinates for face data
func texCoordsForFaceFromTexInfo(vertexes []float32, tx *texinfo.TexInfo, width int, height int) (uvs []float32) {
	for idx := 0; idx < len(vertexes); idx += 3 {
		//u = tv0,0 * x + tv0,1 * y + tv0,2 * z + tv0,3
		u := ((tx.TextureVecsTexelsPerWorldUnits[0][0] * vertexes[idx]) +
			(tx.TextureVecsTexelsPerWorldUnits[0][1] * vertexes[idx+1]) +
			(tx.TextureVecsTexelsPerWorldUnits[0][2] * vertexes[idx+2]) +
			tx.TextureVecsTexelsPerWorldUnits[0][3]) / float32(width)

		//v = tv1,0 * x + tv1,1 * y + tv1,2 * z + tv1,3
		v := ((tx.TextureVecsTexelsPerWorldUnits[1][0] * vertexes[idx]) +
			(tx.TextureVecsTexelsPerWorldUnits[1][1] * vertexes[idx+1]) +
			(tx.TextureVecsTexelsPerWorldUnits[1][2] * vertexes[idx+2]) +
			tx.TextureVecsTexelsPerWorldUnits[1][3]) / float32(height)

		uvs = append(uvs, u, v)
	}

	return uvs
}

// lightmapCoordsForFaceFromTexInfo create lightmap coordinates from TexInfo
func lightmapCoordsForFaceFromTexInfo(vertexes []float32, faceInfo *face.Face, tx *texinfo.TexInfo, width int, height int) (uvs []float32) {
	for idx := 0; idx < len(vertexes); idx += 3 {
		u := (mgl32.Vec3{vertexes[idx], vertexes[idx+1], vertexes[idx+2]}).Dot(
			mgl32.Vec3{
				tx.LightmapVecsLuxelsPerWorldUnits[0][0],
				tx.LightmapVecsLuxelsPerWorldUnits[0][1],
				tx.LightmapVecsLuxelsPerWorldUnits[0][2],
			}) + tx.LightmapVecsLuxelsPerWorldUnits[0][3]
		v := (mgl32.Vec3{vertexes[idx], vertexes[idx+1], vertexes[idx+2]}).Dot(
			mgl32.Vec3{
				tx.LightmapVecsLuxelsPerWorldUnits[1][0],
				tx.LightmapVecsLuxelsPerWorldUnits[1][1],
				tx.LightmapVecsLuxelsPerWorldUnits[1][2],
			}) + tx.LightmapVecsLuxelsPerWorldUnits[1][3]

		u -= float32(faceInfo.LightmapTextureMinsInLuxels[0]) - .5
		v -= float32(faceInfo.LightmapTextureMinsInLuxels[1]) - .5
		u /= float32(faceInfo.LightmapTextureSizeInLuxels[0]) + 1
		v /= float32(faceInfo.LightmapTextureSizeInLuxels[1]) + 1

		//u *= float32(width) // lightmapRect.width
		//v *= float32(height) //lightmapRect.height
		//u += lightmapRect.x
		//v += lightmapRect.y

		uvs = append(uvs, u, v)
	}

	return uvs
}

// lightmapSamplesFromFace create a lightmap rectangle for a face
func lightmapSamplesFromFace(f *face.Face, samples *[]common.ColorRGBExponent32) []common.ColorRGBExponent32 {
	sampleSize := int32(unsafe.Sizeof((*samples)[0]))
	numLuxels := (f.LightmapTextureSizeInLuxels[0] + 1) * (f.LightmapTextureSizeInLuxels[1] + 1)
	firstSampleIdx := f.Lightofs / sampleSize

	return (*samples)[firstSampleIdx : firstSampleIdx+numLuxels]
}
