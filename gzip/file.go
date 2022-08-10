package gzip

import (
	"bufio"
	"compress/gzip"
	"log"
	"os"
)

type F struct {
	f  *os.File
	gf *gzip.Writer
	fw *bufio.Writer
}

func CreateGZ(s string) (f F, err error) {
	fi, err := os.OpenFile(s, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		log.Printf("Error in Create: %s", err.Error())
		return F{}, err
	}
	gf := gzip.NewWriter(fi)
	fw := bufio.NewWriter(gf)
	f = F{fi, gf, fw}
	return
}

func WriteGZ(f F, data []byte) {
	(f.fw).Write(data)
}

func CloseGZ(f F) {
	f.fw.Flush()
	// Close the gzip first.
	f.gf.Close()
	f.f.Close()
}
