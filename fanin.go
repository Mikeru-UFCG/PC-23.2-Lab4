package main

import (
	"fmt"
	"math/rand"
	"time"
)

func request_stream() chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for {
			item := fmt.Sprintf("Item%d", rand.Intn(100))
			ch <- item
			time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
		}
	}()
	return ch
}

func ingest(in chan string, ch1, ch2 chan string) {
	for {
		select {
		case item := <-ch1:
			in <- item
		case item := <-ch2:
			in <- item
		}
	}
}

func main() {
	ch1 := request_stream()
	ch2 := request_stream()

	in := make(chan string)

	go ingest(in, ch1, ch2)

	for {
		item := <-in
		fmt.Println(item)
	}
}