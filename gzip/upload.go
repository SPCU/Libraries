package gzip

import (
	"bufio"
	"compress/gzip"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// SendGzipFile
// Sends a file using gzip compression through HTTP call
func SendGzipFile(targetUrl string, methodType string, filename string) error {
	// Open the file
	in, err := os.Open(filename)
	if err != nil {
		return err
	}

	// gzip writes to pipe, http reads from pipe
	pr, pw := io.Pipe()

	// buffer readers from file, writes to pipe
	buf := bufio.NewReader(in)

	// gzip wraps buffer writer and wr
	gw := gzip.NewWriter(pw)

	// Actually start reading from the file and writing to gzip
	go func() {
		log.Println("Start compression...")
		n, wErr := buf.WriteTo(gw)
		if wErr != nil {
			log.Println(wErr)
		}
		_ = gw.Close()
		_ = pw.Close()
		log.Printf("Done writing: %d", n)
	}()

	// Create new request
	req, err := http.NewRequest(methodType, targetUrl, pr)
	if err != nil {
		return err
	}

	// Do the HTTP Request
	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Printf("File has been sent successfully: %s", string(respBody))

	return nil
}
