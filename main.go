package main

import (
	"flag"
	"fmt"
)

func main() {
	once := flag.Bool("o", false, "only watch changes once")
	flag.Parse()

	ch := make(chan Message)

	go WatchClipboard(ch)

	for msg := range ch {
		if msg.err != nil {
			fmt.Println("error:", msg.err)
		} else {
			fmt.Println(msg.text)
		}
		if *once {
			break
		}
	}
}
