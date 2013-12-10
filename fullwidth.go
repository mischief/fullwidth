package fullwidth

import (
	"bufio"
	"bytes"
	"io"
)

// Convert ascii to full-width
type FullEncoder struct {
	buf *bufio.Writer
}

// Make a new full-width encoder
func Encoder(wr io.Writer) *FullEncoder {
	return &FullEncoder{buf: bufio.NewWriter(wr)}
}

func (f *FullEncoder) Write(p []byte) (n int, err error) {
	str := string(p)

	for _, r := range str {
		if r < 0x21 || r > 0x7e {
			if _, err := f.buf.WriteRune(r); err != nil {
				return 0, err
			}
		} else {
			if _, err := f.buf.WriteRune(r + 0xfee0); err != nil {
				return 0, err
			}
		}
	}

	f.buf.Flush()

	return len(p), nil
}

// Convert full-width to ascii
type FullDecoder struct {
	buf *bufio.Writer
}

// Make a new full-width decoder
func Decoder(wr io.Writer) *FullDecoder {
	return &FullDecoder{buf: bufio.NewWriter(wr)}
}

func (f *FullDecoder) Write(p []byte) (n int, err error) {
	str := string(p)

	for _, r := range str {
		if r < 0xff01 || r > 0xff5e {
			if _, err := f.buf.WriteRune(r); err != nil {
				return 0, err
			}
		} else {
			if _, err := f.buf.WriteRune(r - 0xfee0); err != nil {
				return 0, err
			}
		}
	}

	f.buf.Flush()

	return len(p), nil
}

// Convenience function to convert one string to full width
func FullWidth(str string) string {
	out := new(bytes.Buffer)
	e := Encoder(out)
	io.WriteString(e, str)
	return out.String()
}

// Convenience function to convert one string from half to full width
func HalfWidth(str string) string {
	out := new(bytes.Buffer)
	d := Decoder(out)
	io.WriteString(d, str)
	return out.String()
}
