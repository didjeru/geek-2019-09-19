package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (m MyReader) Read(b []byte) (n int, err error) {
	for n = range b {
		b[n] = 'A'
	}
	return
}

func main() {
	reader.Validate(MyReader{})
}
