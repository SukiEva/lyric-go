package lyric

import (
	"errors"
	"github.com/SukiEva/lyric-go/kugou"
	"github.com/SukiEva/lyric-go/lrc"
	"github.com/SukiEva/lyric-go/netease"
	"github.com/SukiEva/lyric-go/qmusic"
)

type provider interface {
	GetLyric(data lrc.MediaData) string
}

type Lyric struct {
	providers []provider
}

type MediaData lrc.MediaData

func New() *Lyric {
	return &Lyric{}
}

func Default() *Lyric {
	return &Lyric{
		providers: []provider{
			kugou.New(),
			netease.New(),
			qmusic.New(),
		},
	}
}

func (l *Lyric) GetLyric(data MediaData) (string, error) {
	if len(l.providers) == 0 {
		return "", errors.New("no provider selected")
	}
	for _, p := range l.providers {
		lyric := p.GetLyric(lrc.MediaData(data))
		if len(lyric) > 5 {
			return lyric, nil
		}
	}
	return "", errors.New("lyric not found")
}

func (l *Lyric) AddKugou() {
	l.providers = append(l.providers, kugou.New())
}
func (l *Lyric) AddNetease() {
	l.providers = append(l.providers, netease.New())
}
func (l *Lyric) AddQQmusic() {
	l.providers = append(l.providers, qmusic.New())
}
