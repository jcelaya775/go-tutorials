package main

import "io"

type rot13Reader struct {
	r io.Reader
}

const (
	ALPHABET_LENGTH byte = 26
	ROT13_OFFSET    byte = 13
)

var absoluteAsciiByteOffset byte

// Read reads from the underlying reader and modifies the buffer that was read using the rot13 algorithm.
func (rot13 rot13Reader) Read(asciiBuffer []byte) (n int, err error) {
	n, err = rot13.r.Read(asciiBuffer)
	rot13Transform(asciiBuffer)
	return
}

func rot13Transform(asciiBuffer []byte) {
	for i, asciiByte := range asciiBuffer {
		if isAlphabetical(asciiByte) {
			asciiByte = getRot13(asciiByte)
		}
		asciiBuffer[i] = asciiByte
	}
}

func isAlphabetical(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func getRot13(asciiByte byte) byte {
	setOffset(asciiByte)
	return calculateRot13(asciiByte)
}

func setOffset(asciiByte byte) {
	if isUpperCase(asciiByte) {
		absoluteAsciiByteOffset = byte('A')
	} else {
		absoluteAsciiByteOffset = byte('a')
	}
}

func isUpperCase(char byte) bool {
	return char >= 'A' && char <= 'Z'
}

func calculateRot13(asciiByte byte) byte {
	relativeRot13AsciiByte := (asciiByte + ROT13_OFFSET) % ALPHABET_LENGTH
	actualRot13AsciiByte := relativeRot13AsciiByte + absoluteAsciiByteOffset
	return actualRot13AsciiByte
}
