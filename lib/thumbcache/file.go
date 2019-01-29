package thumbcache

var cacheVersion = 1
var cacheFilenamePrefix = "dir_cache_"

// cacheFile represents an in-memory/on-disk cache file
// It begins with a header that defines the rest of the
// cache contents.
// Filepath block containing all filenames included in the cache, NULL-terminated.
// This is followed by positional data for each entry that lies within
// the texture atlas
// Finally is lump pixel data that contains all thumbnail data
type cacheFile struct {
	// Header provides metadata about a cache
	Header header
	// FilePaths is a slice of all filepaths stored in the cache
	FilePaths []string
	// Entries provides data about thumbnail data location in the atlas, order
	// matches FilePaths order
	Entries []CacheEntry
	// Atlas is a raw lump of color data bytes
	Atlas []byte
}

// header provides information about file format and offsets
// of different datastructures contained in the cache file.
type header struct {
	// Version is the cache format version
	Version int32
	// NumEntries is the count of thumbnails in this cache
	NumEntries int32
	// FilePathBlockOffset is offset in bytes to the start of filepath lump
	FilePathBlockOffset int32
	// FilePathBlockLength is size in bytes of filepath lump
	FilePathBlockLength int32
	// EntryBlockOffset is offset in bytes to the start of the thumbnail atlas info
	EntryBlockOffset int32
	// EntryBlockLength size of atlas entry info in bytes
	EntryBlockLength int32
	// AtlasOffset is offset in bytes to raw image data of atlas
	AtlasOffset int32
	// AtlasLength is size of atlas color data in bytes
	AtlasLength int32
}

// CacheEntry provides data about the location of an item in the
// cache atlas
type CacheEntry struct {
	// X is X offset into the cache atlas image
	X float32
	// Y is Y offset into the cache atlas image
	Y float32
	// Width is width of thumbnail inside the cache atlas image
	Width float32
	// Height is height of thumbnail inside the cache atlas image
	Height float32
}