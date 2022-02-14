package hasher

import (
	"crypto/md5"
	"encoding/hex"
)

type md5Hasher struct{}

func NewMd5Hasher() *md5Hasher {
	return &md5Hasher{}
}

func (h *md5Hasher) Hash(data string) string {
	hasher := md5.New()
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}
