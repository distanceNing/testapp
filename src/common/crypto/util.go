package crypto

import "strings"

func GenerateECKeyFromString(keyStr string) string {
	last := 0
	var a []string
	for i := 64; i < len(keyStr); i += 64 {
		a = append(a, keyStr[last:i])
		last = i
	}
	if last < len(keyStr) {
		a = append(a, keyStr[last:])
	}
	priKey := "-----BEGIN PRIVATE KEY-----\n"
	priKey += strings.Join(a, "\n")
	priKey += "\n-----END PRIVATE KEY-----"
	return priKey
}
