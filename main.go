package main

import (
	"fmt"
	"golang-youtube-dl-cli/binaries"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	// Determine the OS-specific binary name.
	var binaryName string
	switch {
	case isWindows():
		binaryName = "yt-dlp.exe"
	default:
		binaryName = "yt-dlp"
	}

	// Now you can use the binary as needed, e.g., save it to a file or execute it.
	// Example: save the binary to a file
	dname, fname := saveBinary(binaryName, binaries.YoutubeDL)

	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {
			panic(err)
		}
	}(dname)

	youtubeURL := "https://www.youtube.com/watch?v=YIkjjLtJP8I"

	cmd := exec.Command(fname, youtubeURL)

	// Redirect standard output and error to capture the output.
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command.
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running youtubedl:", err)
		return
	}

	fmt.Println("Video downloaded successfully!")
}

func isWindows() bool {
	return os.PathSeparator == '\\'
}

func saveBinary(filename string, data []byte) (string, string) {
	dname, err := os.MkdirTemp("", "youtube")
	if err != nil {
		panic(err)
	}

	fname := filepath.Join(dname, filename)
	err = os.WriteFile(fname, data, 0777)

	return dname, fname
}
