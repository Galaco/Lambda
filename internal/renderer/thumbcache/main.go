package thumbcache

//
//
//func main() {
//	cache,err := InitCache("foo")
//	if err != nil {
//		panic(err)
//	}
//
//
//	filesInDirectory := []string{}
//
//	for _,file := range filesInDirectory {
//		getThumbnail(cache, file)
//	}
//}
//
//func getThumbnail(cache *Cache, filepath string) *CacheEntry {
//	if cache.Exists(filepath) {
//		return cache.Get(filepath)
//	}
//
//	thumb := loadVtf(filepath)
//	return cache.Add(filepath, thumb.LowResImageWidth, thumb.LowResImageHeight, thumb.LowResImageData)
//}
//
//func loadVtf(filepath string) []byte{
//	return []byte{}
//}
