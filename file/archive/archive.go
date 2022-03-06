package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
)

func main() {

	//zrc, err := zip.OpenReader("fs.zip")
	zrc, err := zip.OpenReader("readme.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer zrc.Close()

	for _, file := range zrc.File {
		fmt.Printf("file %s:\n", file.Name)

		rc, err := file.Open()
		if err != nil {
			log.Fatalln(err)
		}
		defer rc.Close()

		content, err := readFile(rc)
		if err != nil {
			log.Fatalln(err)
		}

		displayFile(content)

	}
}

func readFile(rc io.ReadCloser) ([]byte, error) {

	b, err := io.ReadAll(rc)
	if err != nil {
		return nil, err
	}
	return b, nil

}

func displayFile(b []byte) {
	fmt.Printf(("%s\n\n"), b)
	//fmt.Println(string(content))
}
