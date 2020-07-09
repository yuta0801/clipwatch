package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	ch := make(chan Message)

	wg.Add(1)
	go WatchClipboard(ch, &wg)

	for msg := range ch {
		if msg.err != nil {
			fmt.Println("error:", msg.err)
		} else {
			fmt.Println(msg.text)
		}
	}

	wg.Wait()
}
