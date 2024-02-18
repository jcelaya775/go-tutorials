package main

import "io"

type rot13Reader struct {
	r io.Reader
}

// Read reads from the underlying reader and modifies the buffer that was read using the rot13 algorithm.
func (rot13 rot13Reader) Read(b []byte) (n int, err error) {
	n, err = rot13.r.Read(b)
	rot13ModifyBuffer(b)
	return
}

// rot13ModifyBuffer takes a slice of bytes and transforms each alphabetical byte using the rot13 algorithm.
func rot13ModifyBuffer(b []byte) {
	for i, char := range b {
		if isAlphabetical(char) {
			char = getRot13Byte(char)
		}
		b[i] = char
	}
}

const ALPHABET_LENGTH byte = 26
const ROT13_OFFSET byte = 13

var absoluteLetterOffset byte

func getRot13Byte(originalLetter byte) byte {
	absoluteLetterOffset = getBaseLetter(originalLetter)
	relativeLetter := getRelativeLetter(originalLetter)
	return ((relativeLetter + ROT13_OFFSET) % ALPHABET_LENGTH) + absoluteLetterOffset
}

func getBaseLetter(char byte) byte {
	if isUpperCase(char) {
		return byte('A')
	} else {
		return byte('a')
	}
}

func getRelativeLetter(absoluteLetter byte) byte {
	return absoluteLetter - absoluteLetterOffset
}

func isAlphabetical(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func isUpperCase(char byte) bool {
	return char >= 'A' && char <= 'Z'
}

func isLowerCase(char byte) bool {
	return char >= 'a' && char <= 'z'
}
