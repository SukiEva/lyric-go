package qmusic

import (
	"encoding/base64"
	"fmt"
	"github.com/SukiEva/lyric-go/lrc"
	"log"
)

type Qmusic struct {
}

const (
	baseUrl         = "https://c.y.qq.com/"
	refererUrl      = "https://y.qq.com"
	searchUrlFormat = baseUrl + "soso/fcgi-bin/client_search_cp?w=%s&format=json"
	lyricUrlFormat  = baseUrl + "lyric/fcgi-bin/fcg_query_lyric_yqq.fcg?songmid=%s&format=json"
)

var c = lrc.NewClient()

func New() *Qmusic {
	return &Qmusic{}
}

func (*Qmusic) GetLyric(data lrc.MediaData) string {
	lyricUrl := getLyricUrl(data)
	if lyricUrl == "" {
		log.Println("Error not finding lyric url")
		return ""
	}
	res, err := c.Get(lyricUrl, refererUrl)
	if err != nil {
		log.Printf("Error getting %s : %s\n", lyricUrl, err)
		return ""
	}
	defer res.Body.Close()
	var jsonObj map[string]interface{}
	if err := lrc.ParseJson(res, &jsonObj); err != nil {
		return ""
	}
	encoded := jsonObj["lyric"].(string)
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		log.Println("Error decoding lyric")
		return ""
	}
	return string(decoded)
}

func getLyricUrl(data lrc.MediaData) string {
	searchUrl := fmt.Sprintf(searchUrlFormat, lrc.GetSearchKey(data))
	res, err := c.Get(searchUrl, refererUrl)
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
	if status != 0 {
		return ""
	}
	candidates := jsonObj["data"].(map[string]interface{})["song"].(map[string]interface{})["list"].([]interface{})
	mid := ""
	for _, cand := range candidates {
		candidate := cand.(map[string]interface{})
		name := candidate["songname"].(string)
		if data.Title == name {
			mid = candidate["songmid"].(string)
		}
	}
	if mid == "" {
		return ""
	}
	return fmt.Sprintf(lyricUrlFormat, mid)
}
