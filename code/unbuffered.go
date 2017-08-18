// +build ignore

package main

func main() {
	ch := make(chan int)
	<-ch // Nothing to receive, ch blocks, main sleeps.
}
