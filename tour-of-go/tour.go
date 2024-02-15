package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
  ticTacToe()
  fmt.Println()

  fmt.Println(Pic(3, 3))
  fmt.Println()

  fmt.Println(WordCount("hello there world"))
  fmt.Println()

  hypot := func(x, y float64) float64 {
    return math.Sqrt(x*x + y*y)
  }
  fmt.Println(hypot(5, 12))

  fmt.Println(compute(hypot))
  fmt.Println(compute(math.Pow))
  fmt.Println()

  x := 144.0
  fmt.Printf("Guessing the square root of %v...\n", x)
  fmt.Println(Sqrt(x))

  // Print return value(s) (aka tuple of values)
  fmt.Println(Sqrt(2))
  fmt.Println(Sqrt(-2))
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
  // NOTE: If you don't convert e to float64 explicitly, then ->
  // Sprintf returns the format specifier of its arguments. So,
  // returning ErrNegativeSqrt will recall this method indefinitely
  // b/c this method is the format specifier for ErrNegativeSqrt
  return fmt.Sprintf("Cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
  if x < 0 {
    return 0, ErrNegativeSqrt(x)
  }

  z := 1.0
  for i := 0; i < 10; i++ {
    z -= (z*z - x) / (2*z)
  }

  return z, nil
}

func ticTacToe() {
  board := [][]string{
    {"_", "_", "_"},
    {"_", "_", "_"},
    {"_", "_", "_"},
  }

  board[0][0] = "X"
  board[1][1] = "O"
  board[2][2] = "X"

  for i := 0; i < len(board); i++ {
    fmt.Printf("%s\n", strings.Join(board[i], " "))
  }
}

func Pic(dx, dy int) [][]uint8 {
  pic := make([][]uint8, dy)
  for y := 0; y < dy; y++ {
    arr := make([]uint8, dx)
    for x := 0; x < dx; x++ {
      arr[x] = uint8(x * y)
    }
    pic[y] = arr
  }

  return pic
}

func WordCount(s string) map[string]int {
  wordCount := make(map[string]int)
  for _, word := range strings.Fields(s) {
    count := wordCount[word]
    wordCount[word] = count + 1
  }

  return wordCount
}

func compute(fn func(float64, float64) float64) float64 {
  return fn(3, 4)
}

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

// NOTE: This is how the Stringer interface is implemented
// type Stringer interface {
//   String() string
// }

// For instance, IPAddr{1, 2, 3, 4} should print as "1.2.3.4".
func (ip IPAddr) String() string {
  return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (reader MyReader) Read(b []byte) (int, error) {
  for i := range b {
    b[i] = 65
  }

  return len(b), nil
}
