# 📦 enczip

[![Go Reference](https://pkg.go.dev/badge/github.com/MatusOllah/enczip/zip.svg)](https://pkg.go.dev/github.com/MatusOllah/enczip/zip) [![CI (Go)](https://github.com/MatusOllah/enczip/actions/workflows/ci.yml/badge.svg)](https://github.com/MatusOllah/enczip/actions/workflows/ci.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/MatusOllah/enczip)](https://goreportcard.com/report/github.com/MatusOllah/enczip) [![GitHub license](https://img.shields.io/github/license/MatusOllah/enczip)](LICENSE) ![Made in Slovakia](https://raw.githubusercontent.com/pedromxavier/flag-badges/refs/heads/main/badges/SK.svg)

**enczip** is a fork of the Go stdlib package `archive/zip` to add support for various text encodings other than just UTF-8.

## 🤔 Why?

> ⚠️ rant incoming

Because Shift-JIS ZIP archives exist.
Sometimes, especially if you tinker with Japanese singing voice synthesis, life hands you an ancient ZIP archive that has been faithfully preserved in glorious Shift-JIS. And sometimes that ZIP archive belongs to a certain UTAU voicebank (*wink wink, Yamine Renri*).

The Go stdlib's `archive/zip` package has a strict "UTF-8 or else" policy, which is great in theory, until you're having to read the most Shift-JIS'd, mojibaked, Go-unfriendly ZIP archive ever produced by 2000s-era software from Japan that apparently never got the memo. This happened to me while working on [one of my other projects](https://github.com/MatusOllah/gotau).
The result? Mojibake, random errors, confusion, and a lot of time spent questioning why I chose this hobby and why I didn't just cough up the ~100 EUR for a more commercial SVS engine like Synthesizer V and call it a day.

So, I forked `archive/zip`, lifted the UTF-8-only restriction, and added better support for obscure text encodings beyond just UTF-8.

Thus, **enczip** - a fork of `archive/zip` with proper handling for non-UTF-8 filename encodings.

tl;dr: Shift-JIS sucks. UTAU should use UTF-8 already.

## 🌟 Usage

enczip is very similar to the original `archive/zip`, except that functions also accept an `encoding.Encoding` value.

### 📖 Reading

```go
package main

import (
    "bytes"
    "fmt"
    "io"
    "os"

    "github.com/MatusOllah/enczip/zip"
    "golang.org/x/text/encoding/japanese"
)

func main() {
    // Open a Shift-JIS zip archive for reading.
    r, err := zip.OpenReader("testdata/shiftjis.zip", japanese.ShiftJIS)
    if err != nil {
        panic(err)
    }
    defer r.Close()

    // Do anything we want with the archive
    // e.g. iterate, FS, ...
    for _, f := range r.File {
        fmt.Printf("Contents of %s:\n", f.Name)
        rc, err := f.Open()
        if err != nil {
            panic(err)
        }
        _, err = io.CopyN(os.Stdout, rc, 68)
        if err != nil && err != io.EOF {
            panic(err)
        }
        rc.Close()
        fmt.Println()
    }
}
```

### ✏️ Writing

```go
package main

import (
    "bytes"
    "fmt"
    "io"
    "os"

    "github.com/MatusOllah/enczip/zip"
    "golang.org/x/text/encoding/japanese"
)

func main() {
    buf := new(bytes.Buffer)

    // Create a new Shift-JIS zip archive (why though?)
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
            panic(err)
        }
        _, err = f.Write([]byte(file.Body))
        if err != nil {
            panic(err)
        }
    }

    if err := w.Close(); err != nil {
        panic(err)
    }
}
```

enczip works of course with any text encoding, not just Shift-JIS :)

## ⚖️ License

Copyright © 2025 Matúš Ollah

Licensed under the **BSD 3-Clause License** (see [LICENSE](LICENSE))

### Go

<https://cs.opensource.google/go/go>

```txt
Copyright 2009 The Go Authors.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are
met:

   * Redistributions of source code must retain the above copyright
notice, this list of conditions and the following disclaimer.
   * Redistributions in binary form must reproduce the above
copyright notice, this list of conditions and the following disclaimer
in the documentation and/or other materials provided with the
distribution.
   * Neither the name of Google LLC nor the names of its
contributors may be used to endorse or promote products derived from
this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
```
