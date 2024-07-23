package main

import (
	"fmt"
	"io"
)

type rot13Reader struct {
	r io.Reader
}

const (
	ALPHABET_LENGTH byte = 26
	ROT13_OFFSET    byte = 13
)

// Read reads from the underlying reader and modifies the buffer that was read using the rot13 algorithm.
func (rot13 rot13Reader) Read(asciiBuffer []byte) (n int, err error) {
	n, err = rot13.r.Read(asciiBuffer) // Read contents of the underlying reader r into asciiBuffer
	fmt.Println("asciiBuffer:", string(asciiBuffer))
	rot13Transform(asciiBuffer)
	return
}

func rot13Transform(asciiBuffer []byte) {
	for i, asciiByte := range asciiBuffer {
		fmt.Println("original asciiByte:", string(asciiByte))
		if isAlphabetical(asciiByte) {
			asciiByte = getRot13(asciiByte)
		}
		fmt.Println("transformed asciiByte:", string(asciiByte))
		asciiBuffer[i] = asciiByte
	}
}

func isAlphabetical(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func getRot13(asciiByte byte) byte {
	var absoluteAsciiByteOffset byte
	if isUpperCase(asciiByte) {
		absoluteAsciiByteOffset = byte('A')
	} else {
		absoluteAsciiByteOffset = byte('a')
	}

	relativeRot13AsciiByte := (asciiByte + ROT13_OFFSET) % absoluteAsciiByteOffset % ALPHABET_LENGTH
	return absoluteAsciiByteOffset + relativeRot13AsciiByte
}

func isUpperCase(char byte) bool {
	return char >= 'A' && char <= 'Z'
}
