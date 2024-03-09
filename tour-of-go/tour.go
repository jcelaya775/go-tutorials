package main

import (
	"fmt"
	"image"
	"image/color"
	"time"

	"math"
	"strings"

	"golang.org/x/tour/tree"
)

func main() {
	ticTacToe()
	fmt.Println()

	fmt.Println(Pic(3, 3))
	fmt.Println()

	fmt.Println(WordCount("hello there world"))
	fmt.Println()

	fmt.Println("\nFunctions:")
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

	// A Reader populates a byte slice with data.
	myBuffer := make([]byte, 25)
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	rot13 := rot13Reader{s}

	fmt.Println("\nrot13Reader:")
	n, err := rot13.Read(myBuffer)
	fmt.Printf("n = %v err = %v\n", n, err)
	fmt.Printf("result: %v\n", string(myBuffer))

	fmt.Println()
	numbers := []int{10, 20, 15, -10}
	fmt.Println(Index(numbers, 10))

	words := []string{"foo", "bar", "obi-wan"}
	fmt.Println(Index(words, "obi-wan"))

	fmt.Println("\nLinked lists:")
	intLinkedListHead := createLinkedListFromSlice([]int{1, 2, 3, 4, 5})
	printLinkedList[int](intLinkedListHead)
	printLinkedList[int](intLinkedListHead)

	stringLinkedListHead := createLinkedListFromSlice([]string{"you're", "the", "chosen", "one", "obi-wan", "kenobi"})
	printLinkedList[string](stringLinkedListHead)
	printLinkedList[string](stringLinkedListHead)
	fmt.Println()

	fmt.Println("Goroutines:")
	go say("world")
	say("hello")
	fmt.Println()

	nums := []int{7, 2, 8, -9, 4, 0}

	fmt.Println("Distributing workload between two goroutines...")
	// i.e, myVar <- c (receive) waits for to something to send to c: c <- value, and vice-versa
	c := make(chan int)
	// Distribute two halves of the work between two goroutines & calculate final result
	go sum(nums[:len(nums)/2], c)
	go sum(nums[len(nums)/2:], c)
	// By default, sends and receives block until the other side is ready, allowing for synchronization
	secondHalf, firstHalf := <-c, <-c // receive from c
	fmt.Printf("firstHalf: %v, secondHalf: %v, firstHalf + secondHalf = %v",
		firstHalf, secondHalf, firstHalf+secondHalf)

	fmt.Println("\n\nbufferedChannelsExample:")
	bufferedChannelsExample()

	// Test whether a channel has been closed
	// v, ok := <-ch
	// Receieve values from channel repeatedly until channel is closed
	// for i := range ch

	fmt.Println("\nFibonacci sequence:")
	ch := make(chan int)
	// NOTE: Must initiate a goroutine for fibonacci function to be able to send to channel
	go fibonacci(10, ch)
	for i := range ch {
		fmt.Println(i)
	}

	// Select
	fmt.Println("\nFibonacci sequence with select block:")
	channel := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-channel)
		}
		// At this point, c will not receieve any more values from the fibonacciSelect function
		quit <- 0 // So, only the quit case will be executed
	}()
	fibonacciSelect(channel, quit)

	fmt.Println("\nDefault selection ticking time bomb:")
	defaultSelectionExample()

	fmt.Println("\nComparing trees...")
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println()

	// In this example, the counter makes sure that only one goroutine can access the counter
	// at any given time
	safeCounter := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go safeCounter.Inc("somekey")
	}

	// NOTE: Once the last goroutine is launched, the program will resume and exit, even if the
	// goroutines are stil running (which will be aborted). So, we will wait for the program to finish.
	time.Sleep(1 * time.Second)
	fmt.Println(safeCounter.Value("somekey"))

	wg.Add(1)
	Crawl("https://golang.org/", 4, fetcher)
	wg.Wait()
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
		z -= (z*z - x) / (2 * z)
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

// For instance, IPAddr{1, 2, 3, 4} should print as "1.2.3.4".
func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

type MyReader struct{}

func (reader MyReader) Read(b []byte) (int, error) {
	populateBytes(b)
	return len(b), nil
}

func populateBytes(b []byte) {
	for i := range b {
		b[i] = byte('A')
	}
}

type Image struct{}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 100, 100)
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x + y), uint8(x ^ y), 255, 255}
}

// Type parameters -> type params appear in brackets before the func arguments, but after its name.
// They allow for the multiple types to satisfy the parameter types, so long as said type satisfies
// the type constraint (the thing that comes after each type parameter is called a type constraint).

// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if x == v {
			return i
		}
	}
	return -1
}

// Generic types -> types can be parameterized with a type parameter, allowing for any type of value.

// List represents a singly-linked list that holds values of any type.
type List[T any] struct {
	val  T
	next *List[T]
}

// This function creates the following linked list and returns its head:
// 1 -> 2 -> 3 -> 4 -> 5 -> nil
func createLinkedListFromSlice[T any](values []T) *List[T] {
	dummyNode := &List[T]{}
	currentNode := dummyNode
	for _, value := range values {
		currentNode.next = &List[T]{val: value, next: nil}
		currentNode = currentNode.next
	}

	return dummyNode.next
}

func printLinkedList[T any](head *List[T]) {
	currentNode := head // copy the head ptr, so we don't lose the head
	for currentNode != nil {
		fmt.Printf("%v -> ", currentNode.val)
		currentNode = currentNode.next
	}
	fmt.Println("nil")
}
