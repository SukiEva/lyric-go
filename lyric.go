package lyric

import (
	"github.com/SukiEva/lyric-go/kugou"
	"github.com/SukiEva/lyric-go/lrc"
	"github.com/SukiEva/lyric-go/netease"
	"github.com/SukiEva/lyric-go/qmusic"
)

type provider interface {
	GetLyric(data lrc.MediaData) string
}

var (
	defaultProviders = []provider{
		kugou.New(),
		netease.New(),
		qmusic.New(),
	}
)

type Lyric struct {
	providers []provider
}

func New() *Lyric {
	return &Lyric{}
}

func (l *Lyric) GetLyric(data lrc.MediaData) string {
	l.providers = append(l.providers, kugou.New())
	return l.providers[0].GetLyric(data)
}
