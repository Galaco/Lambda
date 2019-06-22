package thumbcache

import (
	"crypto/md5"
	"encoding/hex"
)

func hashDirectory(directory string) string {
	hasher := md5.New()
	_,_ = hasher.Write([]byte(directory))
	return hex.EncodeToString(hasher.Sum(nil))
}
