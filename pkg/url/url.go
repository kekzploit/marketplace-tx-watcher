package url

import (
	"crypto/md5"
	"fmt"
)

func GenUrl(title string) string {
	url := fmt.Sprintf("%x", md5.Sum([]byte(title)))
	return url
}
