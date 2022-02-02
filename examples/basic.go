package main

import (
	"fmt"
	"github.com/SukiEva/lyric-go"
)

func main() {
	l := lyric.Default() // use all three providers
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
