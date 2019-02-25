package thumbcache

import (
	"testing"
)

func TestInitCache(t *testing.T) {
	path := "./foo"
	sut, err := InitCache(path)

	if err != nil {
		t.Error(err)
	}
	if sut.fsPath == "" {
		t.Error("Cache created improperly; no local path")
	}
}

func TestCache_Add(t *testing.T) {
	t.Skip()
	path := "./foo"
	sut, err := InitCache(path)
	if err != nil {
		t.Error(err)
	}

	if nil == sut.Add(path+"/ika.vtf", 16, 16, []byte{}) {
		t.Error("no entry returned for added thumbnail")
	}
	if len(sut.cacheFile.FilePaths) == 0 {
		t.Error("failed to add new path to cache FilePath list")
	}
}

func TestCache_Exists(t *testing.T) {
	t.Skip()
	path := "./foo"
	sut, err := InitCache(path)
	if err != nil {
		t.Error(err)
	}

	if nil == sut.Add(path+"/ika.vtf", 16, 16, []byte{}) {
		t.Error("no entry returned for added thumbnail")
	}

	if false == sut.Exists(path+"/ika.vtf") {
		t.Error("no entry found for known existing thumbnail")
	}
}

func TestCache_FlushToDisk(t *testing.T) {
	t.Skip("skipping filesystem writing")
}

func TestCache_Get(t *testing.T) {
	t.Skip()
	path := "./foo"
	sut, err := InitCache(path)
	if err != nil {
		t.Error(err)
	}

	if nil == sut.Add(path+"/ika.vtf", 16, 16, []byte{}) {
		t.Error("no entry returned for added thumbnail")
	}

	if nil == sut.Get(path+"/ika.vtf") {
		t.Error("no entry returned for known existing thumbnail")
	}
}

func TestCache_ThumbnailAtlas(t *testing.T) {
	path := "./foo"
	sut, err := InitCache(path)
	if err != nil {
		t.Error(err)
	}

	sut.ThumbnailAtlas()
}
