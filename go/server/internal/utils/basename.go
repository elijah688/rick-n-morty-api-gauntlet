package utils

import (
	"strings"
)

func Basename(url string) string {
	parts := strings.Split(url, "/")
	return parts[len(parts)-1]
}
