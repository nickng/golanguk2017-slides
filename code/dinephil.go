// +build ignore

package main

import (
	"fmt"
)

type Chopstick struct{}

func chopstick(c chan Chopstick) {
	for {
		c <- Chopstick{}
		<-c
	}
}

func philosopher(name string, left, right chan Chopstick) {
	for {
		select {
		case l := <-left:
			select {
			case r := <-right:
				fmt.Printf("%s can eat!\n", name)
				left <- l
				right <- r
			default:
				// Try again later..
				left <- l
			}
		case r := <-right:
			select {
			case l := <-left:
				fmt.Printf("%s can eat!\n", name)
				right <- r
				left <- l
			default:
				// Try again later..
				right <- r
			}
		}
	}
}

func main() {
	c1, c2, c3 := make(chan Chopstick), make(chan Chopstick), make(chan Chopstick)
	go philosopher("Aristotle", c1, c2) // try to receive chopstick from both channels.
	go philosopher("Plato", c2, c3)
	go philosopher("Socrates", c3, c1)
	go chopstick(c1) // Make chopstick available for receiving.
	go chopstick(c2)
	chopstick(c3)
}
