package kugou

import (
	"encoding/base64"
	"fmt"
	"log"
	"lyric/lrc"
)

type Kugou struct {
}

const (
	baseUrl         = "http://lyrics.kugou.com/"
	searchUrlFormat = baseUrl + "search?ver=1&man=yes&client=pc&keyword=%s&duration=%d"
	lyricUrlFormat  = baseUrl + "download?ver=1&client=pc&id=%s&accesskey=%s&fmt=lrc&charset=utf8"
)

var c = lrc.NewClient()

func New() *Kugou {
	return &Kugou{}
}

func (*Kugou) GetLyric(data lrc.MediaData) string {
	lyricUrl := getLyricUrl(data)
	if lyricUrl == "" {
		log.Println("Error not finding lyric url")
		return ""
	}
	res, err := c.Get(lyricUrl)
	if err != nil {
		log.Printf("Error getting %s : %s\n", lyricUrl, err)
		return ""
	}
	defer res.Body.Close()
	var jsonObj map[string]interface{}
	if err := lrc.ParseJson(res, &jsonObj); err != nil {
		return ""
	}
	encoded := jsonObj["content"].(string)
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		log.Println("Error decoding lyric")
		return ""
	}
	return string(decoded)
}

func getLyricUrl(data lrc.MediaData) string {
	searchUrl := fmt.Sprintf(searchUrlFormat, lrc.GetSearchKey(data), data.Duration)
	res, err := c.Get(searchUrl)
	if err != nil {
		log.Printf("Error getting %s : %s\n", searchUrl, err)
		return ""
	}
	defer res.Body.Close()
	var jsonObj map[string]interface{}
	if err := lrc.ParseJson(res, &jsonObj); err != nil {
		return ""
	}
	status := jsonObj["status"].(float64)
	if status != 200 {
		return ""
	}
	candidates := jsonObj["candidates"].([]interface{})
	maxScore := float64(-1)
	id := ""
	accessKey := ""
	for _, cand := range candidates {
		candidate := cand.(map[string]interface{})
		score := candidate["score"].(float64)
		if score > maxScore {
			maxScore = score
			id = candidate["id"].(string)
			accessKey = candidate["accesskey"].(string)
		}
	}
	if maxScore == -1 {
		return ""
	}
	return fmt.Sprintf(lyricUrlFormat, id, accessKey)
}
