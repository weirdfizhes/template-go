package echohttp

import (
	"net/http"
)

// GetHeaderKeyValue is for getting header key and value from request data
func GetHeaderKeyValue(header http.Header, key string) (value string) {
	for k, v := range header {
		if k == key {
			for _, vLoop := range v {
				value = vLoop
			}
		}
	}

	return
}
