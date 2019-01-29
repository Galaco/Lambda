package thumbcache

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unsafe"
)

type Cache struct {
	fsPath    string
	cacheFile cacheFile
}

// ThumbnailAtlas return the atlas data for the cache
func (cache *Cache) ThumbnailAtlas() *[]byte {
	return &cache.cacheFile.Atlas
}

// Exists returns whether a particular CacheEntry exists in the cache
func (cache *Cache) Exists(filePath string) bool {
	for _,cacheEntry := range cache.cacheFile.FilePaths {
		if cacheEntry == filePath {
			return true
		}
	}

	return false
}

// Get returns the positional data within the cached atlas image
func (cache *Cache) Get(filePath string) *CacheEntry {
	for idx,cacheEntry := range cache.cacheFile.FilePaths {
		 if cacheEntry == filePath {
		 	return &cache.cacheFile.Entries[idx]
		 }
	}

	return nil
}

// Add adds a new item to the cache
func (cache *Cache) Add(filePath string, width int, height int, color []byte) *CacheEntry {
	// @TODO write me
	return nil
}

// FlushToDisk will write the existing cache data to disk
func (cache *Cache) FlushToDisk() error {
	if err := cache.deleteFSCopy(); err != nil {
		return err
	}
	if err := cache.writeFSCopy(); err != nil {
		return err
	}
	return nil
}

// deleteFSCopy deletes existing cache on local fs
func (cache *Cache) deleteFSCopy() error {
	if _,err := os.Stat(cache.fsPath); err == nil {
		return os.Remove(cache.fsPath)
	}
	return nil
}

// writeFSCopy to file exports the current cache in memory to a file
func (cache *Cache) writeFSCopy() error {
	fsCacheBuffer := make([]byte, 0)

	// write header last (it contains offsets for later data)
	offset := int(unsafe.Sizeof(cache.cacheFile.Header))

	cache.cacheFile.Header.Version = int32(cacheVersion)
	cache.cacheFile.Header.FilePathBlockOffset = int32(offset)

	// write filepaths
	for _,path := range cache.cacheFile.FilePaths {
		fsCacheBuffer = append(fsCacheBuffer, []byte(path)...)
		fsCacheBuffer = append(fsCacheBuffer, []byte("\x00")...)
		offset += len(path) + len("\x00")
	}
	cache.cacheFile.Header.FilePathBlockLength = int32(offset) - cache.cacheFile.Header.FilePathBlockOffset
	cache.cacheFile.Header.EntryBlockOffset = int32(offset)

	// Write positional entries
	entrySize := int(unsafe.Sizeof(CacheEntry{}))
	for _,entry := range cache.cacheFile.Entries {
		var buf bytes.Buffer
		err := binary.Write(&buf, binary.LittleEndian, entry)
		if err != nil {
			return err
		}
		fsCacheBuffer = append(fsCacheBuffer, buf.Bytes()...)

		offset += entrySize
	}
	cache.cacheFile.Header.NumEntries = int32(len(cache.cacheFile.Entries))
	cache.cacheFile.Header.EntryBlockLength = int32(offset) - cache.cacheFile.Header.EntryBlockOffset
	cache.cacheFile.Header.AtlasOffset = int32(offset)

	// Write atlas pixel data
	cache.cacheFile.Header.AtlasLength = int32(len(cache.cacheFile.Atlas))
	fsCacheBuffer = append(fsCacheBuffer, cache.cacheFile.Atlas...)

	// Write header last
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.LittleEndian, cache.cacheFile.Header)
	if err != nil {
		return err
	}
	fsCacheBuffer = append(fsCacheBuffer, buf.Bytes()...)

	return ioutil.WriteFile(cache.fsPath, fsCacheBuffer, 0644)
}

// loadFromFile opens an on-disk cache file and loads it into memory
func (cache *Cache) loadFromFile() error {
	f,err := os.Open(cache.fsPath)
	if err != nil {
		return err
	}
	defer f.Close()

	raw,err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	// Read header
	head := header{}
	err = binary.Read(bytes.NewBuffer(raw), binary.LittleEndian, &head)
	if err != nil {
		return err
	}
	cache.cacheFile.Header = head

	// Read filepaths
	paths := strings.Split(string(raw[head.FilePathBlockOffset:head.FilePathBlockOffset+head.FilePathBlockLength]), "\x00")
	if len(paths) != int(head.NumEntries) {
		return errors.New(fmt.Sprintf("cache entry data corrupted for cache %s", cache.fsPath))
	}

	// Read entries
	entries := make([]CacheEntry, head.NumEntries)
	err = binary.Read(bytes.NewBuffer(raw[head.EntryBlockOffset:head.EntryBlockOffset+head.EntryBlockLength]), binary.LittleEndian, &entries)
	if err != nil {
		return err
	}
	cache.cacheFile.Entries = entries

	// Read Cache Texture data
	cache.cacheFile.Atlas = raw[head.AtlasOffset:head.AtlasOffset+head.AtlasLength]

	return nil
}

// InitCache prepares a cache for a directory.
// It will attempt to load an existing cachefile if it exists
func InitCache(cachePath string) (*Cache, error) {
	cache := &Cache{}
	cache.fsPath = cacheFilenamePrefix + hashDirectory(cachePath)
	if _,err := os.Stat(cachePath); err == nil {
		// Cache exists, attempt to load it
		err = cache.loadFromFile()
		if err != nil {
			// Cache is invalid
			return nil,err
		}
	} else if os.IsNotExist(err) {
		// no cache exists to load
		cache.cacheFile = cacheFile{}
	} else {
		return nil, errors.New(fmt.Sprintf("failed to initialize cache for path: %s", cachePath))
	}

	return cache,nil
}