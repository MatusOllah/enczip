// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package zip_test

import (
	"bytes"
	"compress/flate"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/MatusOllah/enczip/zip"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/japanese"
)

func ExampleWriter() {
	// Create a buffer to write our archive to.
	buf := new(bytes.Buffer)

	// Create a new Shift-JIS zip archive.
	w := zip.NewWriter(buf, japanese.ShiftJIS)

	// Add some files to the archive.
	var files = []struct {
		Name, Body string
	}{
		{"お読みください.txt", "このアーカイブにはいくつかのテキストファイルが含まれています。"},
		{"ボーカロイド一覧.txt", "ボーカロイドの名前:\nミク\nリン\nレン\nルカ\nMEIKO\nKAITO"},
	}

	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}

	// Make sure to check the error on Close.
	err := w.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleReader() {
	// Open a Shift-JIS zip archive for reading.
	r, err := zip.OpenReader("testdata/shiftjis.zip", japanese.ShiftJIS)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	// Iterate through the files in the archive,
	// printing some of their contents.
	for _, f := range r.File {
		fmt.Printf("Contents of %s:\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.CopyN(os.Stdout, rc, 68)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		rc.Close()
		fmt.Println()
	}
	// Output:
	// Contents of 初音ミクが好きです.txt:
	// oo ee oo~
	// Contents of サブフォルダ/readme.txt:
	// Hello world!
}

func ExampleWriter_RegisterCompressor() {
	// Override the default Deflate compressor with a higher compression level.

	// Create a buffer to write our archive to.
	buf := new(bytes.Buffer)

	// Create a new zip archive.
	w := zip.NewWriter(buf, encoding.Nop)

	// Register a custom Deflate compressor.
	w.RegisterCompressor(zip.Deflate, func(out io.Writer) (io.WriteCloser, error) {
		return flate.NewWriter(out, flate.BestCompression)
	})

	// Proceed to add files to w.
}
