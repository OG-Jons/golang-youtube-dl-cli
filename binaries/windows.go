//go:build windows

package binaries

import _ "embed"

//go:embed resources/yt-dlp.exe
var YoutubeDL []byte
