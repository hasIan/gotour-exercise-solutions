package main

import (
	"io"
	"log"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r13r rot13Reader) Read(bytes []byte) (int, error) {
	length, err := r13r.r.Read(bytes)
	if err != nil && err != io.EOF {
		log.Fatal("Error reading stream", err)
	}

	for i := 0; i < length; i++ {
		bytes[i] = rot13(bytes[i])
	}

	return length, err
}

func isAlpha(b byte) (bool, bool) {
	return isUpper(b) || isLower(b), isUpper(b)
}

func rot13(b byte) byte {
	switch {
	case isUpper(b):
		return rotTransform(b, 'A')
	case isLower(b):
		return rotTransform(b, 'a')
	default:
		return b
	}
}

func rotTransform(b, transformer byte) byte {
	return (((b - transformer) + 13) % 26) + transformer
}

func isUpper(b byte) bool {
	return b >= 'A' && b <= 'Z'
}

func isLower(b byte) bool {
	return b >= 'a' && b <= 'z'
}

func main() {
	s := strings.NewReader(
		"Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
