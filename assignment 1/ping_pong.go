    package main
     
    import (
    	"fmt"
    	"time"
    )
     
    func main() {
    	ping := make(chan struct{}, 1)
    	pong := make(chan struct{}, 1)
     
    	ping<- struct{}{}
     
    	go play(ping, pong)
     
    	time.Sleep(time.Millisecond)
    }
     
    func play(ping, pong chan struct{}) {
    	for {
    		select {
    		case <-ping:
    			fmt.Println("ping")
    			pong<- struct{}{}
    		case <-pong:
    			fmt.Println("    pong")
    			ping<- struct{}{}
    		}
    	}
    }



package main

import (
	"fmt"
	"time"
)

// STARTMAIN1 OMIT
type Ball struct{ hits int }

func main() {
	table := make(chan *Ball)
	go player("ping", table)
	go player("pong", table) // HL

	table <- new(Ball) // game on; toss the ball
	time.Sleep(1 * time.Second)
	<-table // game over; grab the ball
}

func player(name string, table chan *Ball) {
	for i := 0; ; i++ {
		ball := <-table
		ball.hits++
		fmt.Println(name, i, "hit", ball.hits)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}