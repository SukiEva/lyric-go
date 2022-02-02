# LYRIC-GO

This library provides an API to search for lyrics from various providers in China.

## Supported Providers

- KuGou Music
- QQ Music
- Netease Music

## Usage

Before using, you need to install it

```
go get github.com/SukiEva/lyric-go
```

### Basic Usage

```go
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
```

### Custom Usage

```go
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
```

## Contributing

You are more than welcome to contribute to this project. Fork and make a Pull Request, or create an Issue if you see any problem or want to propose a feature.