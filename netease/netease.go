package netease

import (
	"fmt"
	"github.com/SukiEva/lyric-go/lrc"
	"log"
)

type Netease struct {
}

const (
	baseUrl         = "http://music.163.com/api/"
	searchUrlFormat = baseUrl + "search/get?s=%s&type=1&offset=0&limit=5"
	lyricUrlFormat  = baseUrl + "song/lyric?os=pc&id=%d&lv=-1&kv=-1&tv=-1"
)

var c = lrc.NewClient()

func New() *Netease {
	return &Netease{}
}

func (*Netease) GetLyric(data lrc.MediaData) string {
	lyricUrl := getLyricUrl(data)
	if lyricUrl == "" {
		log.Println("Error not finding lyric url")
		return ""
	}
	res, err := c.Get(lyricUrl, "")
	if err != nil {
		log.Printf("Error getting %s : %s\n", lyricUrl, err)
		return ""
	}
	defer res.Body.Close()
	var jsonObj map[string]interface{}
	if err := lrc.ParseJson(res, &jsonObj); err != nil {
		return ""
	}
	lyric := jsonObj["lrc"].(map[string]interface{})["lyric"].(string)
	return lyric
}

func getLyricUrl(data lrc.MediaData) string {
	searchUrl := fmt.Sprintf(searchUrlFormat, lrc.GetSearchKey(data))
	res, err := c.Get(searchUrl, "")
	if err != nil {
		log.Printf("Error getting %s : %s\n", searchUrl, err)
		return ""
	}
	defer res.Body.Close()
	var jsonObj map[string]interface{}
	if err := lrc.ParseJson(res, &jsonObj); err != nil {
		return ""
	}
	status := jsonObj["code"].(float64)
	if status != 200 {
		return ""
	}
	candidates := jsonObj["result"].(map[string]interface{})["songs"].([]interface{})
	id := -1.0
	for _, cand := range candidates {
		candidate := cand.(map[string]interface{})
		name := candidate["name"].(string)
		if data.Title == name {
			id = candidate["id"].(float64)
		}
	}
	if id == -1.0 {
		return ""
	}
	return fmt.Sprintf(lyricUrlFormat, int(id))
}
