package main

import (
	"github.com/mischief/fullwidth"
	"io"
	"os"
)

func main() {
	e := fullwidth.Encoder(os.Stdout)
	io.Copy(e, os.Stdin)
}
