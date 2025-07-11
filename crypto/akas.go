package crypto

import (
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

var salt string = "7d4e9c2f8b1a5036d2c8f4e9b7a3d1c6e5f8a2b4d7c9e3f6a8b2d5c7e9f4a1b3"

func getDynSalt() string {
	uuid := strings.ReplaceAll(uuid.New().String(), "-", "")
	sum := sha256.Sum256([]byte(uuid + salt))
	return fmt.Sprintf("%x", sum)
}

func GenAkas() (ak, as string) {
	ak = strings.ReplaceAll(uuid.New().String(), "-", "")
	uuid := strings.ReplaceAll(uuid.New().String(), "-", "")
	as = fmt.Sprintf("%x", sha1.Sum([]byte(uuid+getDynSalt())))
	return
}
