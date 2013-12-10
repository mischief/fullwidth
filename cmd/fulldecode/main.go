package main

import (
	"github.com/mischief/fullwidth"
	"io"
	"os"
)

func main() {
	d := fullwidth.Decoder(os.Stdout)
	io.Copy(d, os.Stdin)
}
