package util

import (
	"strings"
)

func ToWebsafeBase64(str string) string {
	r := strings.NewReplacer("+", "-", "/", "_")

	return r.Replace(strings.Trim(str, "="))
}
