package thumbcache

import (
	"crypto/md5"
	"encoding/hex"
)

func hashDirectory(directory string) string {
	hasher := md5.New()
	hasher.Write([]byte(directory))
	return hex.EncodeToString(hasher.Sum(nil))
}