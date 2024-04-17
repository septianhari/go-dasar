package main

//simple ping pong part 1:
// a c (acceptece criteria)
// 2 player you can use any name for these players
// game stop after 1 s
// use concurency with 1 channel l
// display how many hits that ball got hit

import (
	"fmt"
	"time"
)

type Ball struct {
	hits int
}

func player(name string, ch chan *Ball) {
	for {
		ball := <-ch
		ball.hits++
		fmt.Printf("%s hits the ball, total hits: %d\n", name, ball.hits)
		time.Sleep(100 * time.Millisecond) // Simulate some processing time
		if ball.hits == 10 {
			break
		}
		ch <- ball
	}
}

func main() {
	ball := &Ball{}
	ch := make(chan *Ball)
	go player("Player 1", ch)
	go player("Player 2", ch)

	ch <- ball // Start the game by sending the ball to one of the players

	time.Sleep(1 * time.Second) // Let the game run for 1 second
	close(ch)                   // Close the channel to end the game
}
