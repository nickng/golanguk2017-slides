// +build ignore

package main

func main() {
	ch2 := make(chan bool, 1)
	ch2 <- true // This is OK.
	// ch2 is now full.
	ch2 <- false // No buffer space in ch2, ch2 blocks, main sleeps.
}
