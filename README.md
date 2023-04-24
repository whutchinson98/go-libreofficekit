# go-libreofficekit

CGo bindings to [LibreOfficeKit](https://docs.libreoffice.org/libreofficekit.html)

# Install
```bash
# Latest version of LibreOffice (5.2) is required
$ sudo add-apt-repository ppa:libreoffice/ppa
$ sudo apt-get update
$ sudo apt-get install libreoffice libreofficekit-dev
$ go get github.com/docsbox/go-libreofficekit
```

# Usage

This example demonstrates how to convert Microsoft Office document to PDF

```go
package main

import "github.com/dveselov/go-libreofficekit"

func main() {
    office, _ := libreofficekit.NewOffice("/path/to/libreoffice")

    document, _ := office.LoadDocument("kittens.docx")
    document.SaveAs("kittens.pdf", "pdf", "skipImages")

    document.Close()
    office.Close()
}

```
