package auth

import "github.com/zuzuka28/music_land_api/pkg/uniqkey"

var identityKey = uniqkey.Gen() //nolint:gockecknoglobals

func IdentityKey() string {
	return identityKey.String()
}
