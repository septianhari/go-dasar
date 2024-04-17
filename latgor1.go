package main

import (
	"log"
	"math/rand"
	"time"
)

// simple ping pong part 1:
// a c (acceptance criteria) :
// 2 player you can use any name for these players
// game stop after 1 s
// use concurency with 1 channel l
// display how many hits that ball got hit

// next feature :
// goroutine for referee
// referee will get the ball and decide who the winner
// u can use random number and modulus to decide player when drop the ball

type Ball struct {
	hits       int
	lastPlayer string
}

func referee(table chan *Ball, done chan *Ball) {
	table <- new(Ball)

	for {
		select {
		case ball := <-done:
			log.Println(ball.lastPlayer, "wins the game")
			return
		}
	}
}

func main() {

	// declare type data that channel can accept
	table := make(chan *Ball)
	done := make(chan *Ball)

	go player("dino", table, done)
	go player("arya", table, done)

	referee(table, done)

	// <-table // game over

}

func player(name string, table chan *Ball, done chan *Ball) {
	for {
		ball := <-table //blocking

		randNumb := rand.Intn(10000)
		if randNumb%17 == 0 {
			log.Println(name, "failed hit the ball")
			done <- ball
			return
		}
		ball.hits++
		ball.lastPlayer = name
		log.Println(name, "hits the ball", ball.hits)
		time.Sleep(50 * time.Millisecond)
		table <- ball
	}
}
