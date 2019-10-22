package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Printf("%s\n", time.Now().String())
		time.Sleep(time.Second)
	}
}
