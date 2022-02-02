package lrc

import (
	"net/url"
)

func GetSearchKey(data MediaData) string {
	var key string
	if data.Artist != "" {
		key = data.Artist + "-" + data.Title
	} else if data.Album != "" {
		key = data.Album + "-" + data.Title
	}
	return url.QueryEscape(key)
}
