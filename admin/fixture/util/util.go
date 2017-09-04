package util

import (
	"os"
	"bufio"
	"encoding/base64"
)

func ImageToBase64(filename string) string {
	image, _ := os.Open(filename)
	defer image.Close()

	info, _ := image.Stat()

	var size int64 = info.Size()
	buffer := make([]byte, size)

	fReader := bufio.NewReader(image)
	fReader.Read(buffer)

	return base64.StdEncoding.EncodeToString(buffer)
}
