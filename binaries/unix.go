//go:build linux || unix || darwin

package binaries

import _ "embed"

//go:embed resources/yt-dlp_linux
var YoutubeDL []byte
