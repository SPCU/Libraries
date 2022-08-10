package gzip

import (
	"compress/gzip"
	"io/ioutil"
	"net/http"
)

// ParserHandler
// is an HTTP handler that returns
// the data of the uploaded gzip in []byte.
func ParserHandler(r *http.Request) ([]byte, error) {
	defer r.Body.Close()

	// Read gzip file
	gr, err := gzip.NewReader(r.Body)
	if err != nil {
		return nil, err
	}
	defer gr.Close()

	return ioutil.ReadAll(gr)
}
