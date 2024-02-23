package main

import "io"

type rot13Reader struct {
	r io.Reader
}

const (
	ALPHABET_LENGTH byte = 26
	ROT13_OFFSET    byte = 13
)

var asciiByteOffset byte

// Read reads from the underlying reader and modifies the buffer that was read using the rot13 algorithm.
func (rot13 rot13Reader) Read(asciiBuffer []byte) (n int, err error) {
	n, err = rot13.r.Read(asciiBuffer)
	rot13ModifyAsciiBuffer(asciiBuffer)
	return
}

func rot13ModifyAsciiBuffer(asciiBuffer []byte) {
	for i, asciiByte := range asciiBuffer {
		if isAlphabetical(asciiByte) {
			asciiByte = getRot13AsciiByte(asciiByte)
		}
		asciiBuffer[i] = asciiByte
	}
}

func isAlphabetical(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func getRot13AsciiByte(originalAsciiByte byte) byte {
	setAsciiByteOffset(originalAsciiByte)
	return getRot13AsciiByteCalculation(originalAsciiByte)
}

func setAsciiByteOffset(asciiByte byte) {
	if isUpperCase(asciiByte) {
		asciiByteOffset = byte('A')
	} else {
		asciiByteOffset = byte('a')
	}
}

func isUpperCase(char byte) bool {
	return char >= 'A' && char <= 'Z'
}

func getRot13AsciiByteCalculation(originalAsciiByte byte) byte {
	relativeRot13AsciiByte := (originalAsciiByte + ROT13_OFFSET) % ALPHABET_LENGTH
	actualRot13AsciiByte := relativeRot13AsciiByte + asciiByteOffset
	return actualRot13AsciiByte
}
