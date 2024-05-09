package logging

import "github.com/zuzuka28/music_land_api/pkg/uniqkey"

var requestIDKey = uniqkey.Gen() //nolint:gockecknoglobals

func RequestIDKey() string {
	return requestIDKey.String()
}
