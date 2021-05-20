package main

import (
	"crypto/rand"
	"io"
	"io/ioutil"
	"log"

	"github.com/cheggaaa/pb/v3"
)

func main() {

	var limit int64 = 1024 * 1024 * 500
	// we will copy 200 Mb from /dev/rand to /dev/null
	reader := io.LimitReader(rand.Reader, limit)
	writer := ioutil.Discard

	// start new bar
	bar := pb.Full.Start64(limit)
	// create proxy reader
	barReader := bar.NewProxyReader(reader)
	// copy from proxy reader
	if _, err := io.Copy(writer, barReader); err != nil {
		log.Fatal(err)
	}
	// finish bar
	bar.Finish()
}
