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
}

func Sqrt(x float64) float64 {
  z := 1.0
  actualRoot := math.Sqrt(x)

  for math.Abs((z - actualRoot)) > 0.1 {
    fmt.Println("z: ", z)
    fmt.Println("diff =", math.Abs((z - actualRoot)))
    z -= (z*z - x) / (2*z)
  }

  fmt.Println("final value of z:", z)

  return z
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
    count += 1
    wordCount[word] = count
  }

  return wordCount
}

func compute(fn func(float64, float64) float64) float64 {
  return fn(3, 4)
}
