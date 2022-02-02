package netease

import (
	"lyric/lrc"
)

type Netease struct {
}

const (
	baseUrl         = "http://music.163.com/api/"
	searchUrlFormat = baseUrl + "search/get?s=%s&type=1&offset=0&limit=5"
	lyricUrlFormat  = baseUrl + "song/lyric?os=pc&id=%d&lv=-1&kv=-1&tv=-1"
)

func New() *Netease {
	return &Netease{}
}

func (*Netease) GetLyric(data lrc.MediaData) string {

	return ""
}

func getLyricUrl(data lrc.MediaData) string {
	return ""
}
