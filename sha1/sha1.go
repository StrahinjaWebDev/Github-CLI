package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	sig, err := sha1Sum("http.log.gz")

	if err != nil {
		log.Fatalf("error: %s", err)
	}

	fmt.Println(sig)

	sig, err = sha1Sum("sha1.go")

	if err != nil {
		log.Fatalf("error: %s", err)
	}

	fmt.Println(sig)
}

// $ cat http.log.gz.gz | gunzip | sha1sum
func sha1Sum(fileName string) (string, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return "", nil
	}

	defer file.Close()
	var r io.Reader = file

	if strings.HasSuffix(fileName, ".gz") {
		gz, err := gzip.NewReader(file)

		if err != nil {
			return "", err
		}

		defer file.Close()
		r = gz
	}

	w := sha1.New()

	if _, err := io.Copy(w, r); err != nil {
		return "", err
	}

	sig := w.Sum(nil)

	return fmt.Sprintf("%x", sig), nil
}
