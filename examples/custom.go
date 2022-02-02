package main

import (
	"fmt"
	"github.com/SukiEva/lyric-go"
)

func main() {
	l := lyric.New() // blanket provider
	l.AddKugou()
	l.AddNetease()
	l.AddQQmusic()
	data := lyric.MediaData{
		Title:    "孤独",  // Must
		Artist:   "邓紫棋", // Suggest
		Album:    "",    // Suggest
		Duration: 0,
	}
	lrc, err := l.GetLyric(data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(lrc)
}
